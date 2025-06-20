package voter

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/nocturna-ta/common-model/models/event"
	"github.com/nocturna-ta/common-model/models/logging"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/http"
	"github.com/nocturna-ta/golib/http/filehandler"
	"github.com/nocturna-ta/golib/log"
	response2 "github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/ums/pkg/constants"
	"time"
)

func (m *Module) RegisterVoter(ctx context.Context, req *request.VoterRegistrationRequest) (*response.VoterRegistrationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterUseCases.RegisterVoter")
	defer span.End()

	var (
		voter *model.Voter
	)

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	transaction := func(txCtx context.Context) (any, error) {
		voter = model.ConstructRegistration(req)
		voter.UserID = reqCtx.GetUserId()

		errTx := m.voterRepo.InsertVoter(txCtx, voter)
		if errTx != nil {
			if errors.Is(errTx, dao.ErrDuplicate) {
				return nil, &custerr.ErrChain{
					Message: "User already exists",
					Code:    400,
					Type:    response2.ErrBadRequest,
					Cause:   errTx,
				}
			}

			return nil, errTx
		}

		txHash, err := m.voterRepo.SendTxVoterBlockchain(txCtx, req.SignedTransaction)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"voter": voter,
			}).ErrorWithCtx(txCtx, "[VoterUseCases.RegisterVoter] failed to send transaction to blockchain")
			return nil, err
		}

		if txHash == "" {
			return nil, err
		}

		errTx = m.publisher.Publish(txCtx, m.topics.MasterDataVoter.Value, voter.ID.String(), voter.ToMessageModel(), map[string]any{
			constants.MetaDataOperation: constants.Create,
		})
		if errTx != nil {
			return nil, errTx
		}

		errTx = m.publisher.Publish(txCtx, m.topics.UserLogs.Value, reqCtx.GetUserId().String(), logging.KPULogs{
			BaseModelMessage: event.BaseModelMessage{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsDeleted: false,
			},
			UserID:   reqCtx.GetUserId().String(),
			Address:  reqCtx.GetAddress(),
			Role:     reqCtx.Role,
			Activity: "Voter Registered With ID : " + voter.ID.String() + " by " + reqCtx.GetUserId().String(),
		}, map[string]any{
			constants.MetaDataOperation: constants.Create,
		})

		return nil, nil
	}

	_, err = m.txMgr.Execute(ctx, transaction, nil)
	if err != nil {
		return nil, err
	}

	return &response.VoterRegistrationResponse{
		IsRegistered: true,
	}, err
}

func (m *Module) GetVoterByNIK(ctx context.Context, nik string) (*response.VoterResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterUseCases.GetVoterByNIK")
	defer span.End()

	voter, err := m.voterRepo.GetVoterByNIK(ctx, nik)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"nik":   nik,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByNIK] Failed to get voter by NIK")
		return nil, &custerr.ErrChain{
			Message: "Failed to get voter by NIK",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	voterContract, err := m.voterContract.GetVoterByNIK(nil, nik)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"nik":   nik,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByNIK] Failed to get voter by NIK from contract")
		return nil, &custerr.ErrChain{
			Message: "Failed to get voter by NIK from contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	if voter.NIK != voterContract.Nik {
		log.WithFields(log.Fields{
			"error": err,
			"nik":   nik,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByNIK] NIK mismatch between database and contract")
		return nil, &custerr.ErrChain{
			Message: "NIK mismatch between database and contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
		}
	}

	return &response.VoterResponse{
		ID:                 voter.ID.String(),
		UserID:             voter.UserID.String(),
		NIK:                voter.NIK,
		FullName:           voter.FullName,
		Gender:             voter.Gender,
		BirthPlace:         voter.BirthPlace,
		BirthDate:          voter.BirthDate.Format("2006-01-02"),
		ResidentialAddress: voter.ResidentialAddress,
		VoterAddress:       voter.VoterAddress,
		Region:             voter.Region,
		IsRegistered:       voter.IsRegistered,
		HasVoted:           voter.HasVoted,
	}, err
}

func (m *Module) GetVoterByAddress(ctx context.Context) (*response.VoterResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterUseCases.GetVoterByAddress")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	voter, err := m.voterRepo.GetVoterByAddress(ctx, reqCtx.GetAddress())
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByAddress] Failed to get voter by address")
		return nil, &custerr.ErrChain{
			Message: "Failed to get voter by address",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	voterContract, err := m.voterContract.GetVoterByAddress(nil, common.HexToAddress(reqCtx.GetAddress()))
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByAddress] Failed to get voter by address from contract")
		return nil, &custerr.ErrChain{
			Message: "Failed to get voter by address from contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	if voter.VoterAddress != voterContract.VoterAddress.String() {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByAddress] Address mismatch between database and contract")
		return nil, &custerr.ErrChain{
			Message: "Address mismatch between database and contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
		}
	}

	return &response.VoterResponse{
		ID:                 voter.ID.String(),
		UserID:             voter.UserID.String(),
		NIK:                voter.NIK,
		FullName:           voter.FullName,
		Gender:             voter.Gender,
		BirthPlace:         voter.BirthPlace,
		BirthDate:          voter.BirthDate.Format("2006-01-02"),
		ResidentialAddress: voter.ResidentialAddress,
		VoterAddress:       voter.VoterAddress,
		Region:             voter.Region,
		IsRegistered:       voter.IsRegistered,
		HasVoted:           voter.HasVoted,
	}, err
}

func (m *Module) GetVoterByRegion(ctx context.Context, region string) (*[]response.VoterResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterUseCases.GetVoterByRegion")
	defer span.End()

	voters, err := m.voterRepo.GetVoterByRegion(ctx, region)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"region": region,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByRegion] Failed to get voter by region")
		return nil, &custerr.ErrChain{
			Message: "Failed to get voter by region",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	voterContract, err := m.voterContract.GetVoterByRegion(nil, region)
	if err != nil {
		log.WithFields(log.Fields{
			"error":  err,
			"region": region,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByRegion] Failed to get voter by region from contract")
		return nil, &custerr.ErrChain{
			Message: "Failed to get voter by region from contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	regions := make(map[string]bool)
	for _, voter := range voterContract {
		regions[voter.Region] = true
	}

	var res []response.VoterResponse
	for _, voter := range voters {
		if regions[voter.Region] {
			res = append(res, response.VoterResponse{
				ID:                 voter.ID.String(),
				UserID:             voter.UserID.String(),
				NIK:                voter.NIK,
				FullName:           voter.FullName,
				Gender:             voter.Gender,
				BirthPlace:         voter.BirthPlace,
				BirthDate:          voter.BirthDate.Format("2006-01-02"),
				ResidentialAddress: voter.ResidentialAddress,
				VoterAddress:       voter.VoterAddress,
				Region:             voter.Region,
				IsRegistered:       voter.IsRegistered,
				HasVoted:           voter.HasVoted,
			})
		}
	}
	return &res, err
}

func (m *Module) GetAllVoter(ctx context.Context) (*[]response.VoterResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterUseCases.GetAllVoter")
	defer span.End()

	voters, err := m.voterRepo.GetAllVoter(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetAllVoter] Failed to get all voter")
	}

	voterContract, err := m.voterContract.GetAllVoter(nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetAllVoter] Failed to get all voter from contract")
		return nil, &custerr.ErrChain{
			Message: "Failed to get all voter from contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	contractAddress := make(map[string]bool)
	for _, voter := range voterContract {
		contractAddress[voter.VoterAddress.String()] = true
	}

	var res []response.VoterResponse
	for _, voter := range voters {
		if contractAddress[voter.VoterAddress] {
			res = append(res, response.VoterResponse{
				ID:                 voter.ID.String(),
				UserID:             voter.UserID.String(),
				NIK:                voter.NIK,
				FullName:           voter.FullName,
				Gender:             voter.Gender,
				BirthPlace:         voter.BirthPlace,
				BirthDate:          voter.BirthDate.Format("2006-01-02"),
				ResidentialAddress: voter.ResidentialAddress,
				VoterAddress:       voter.VoterAddress,
				Region:             voter.Region,
				IsRegistered:       voter.IsRegistered,
				HasVoted:           voter.HasVoted,
			})
		}
	}

	return &res, err
}

func (m *Module) GetVoterKTPPhoto(ctx context.Context, id uuid.UUID) (*http.File, string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterUseCases.GetVoterKTPPhoto")
	defer span.End()

	voter, err := m.voterRepo.GetVoterByID(ctx, id)
	if err != nil {
		return nil, "", &custerr.ErrChain{
			Message: "Failed to get voter by ID",
			Code:    400,
			Type:    response2.ErrBadRequest,
			Cause:   err,
		}
	}

	if voter.KTPPhotoPath == "" {
		return nil, "", &custerr.ErrChain{
			Message: "KTP photo not found",
			Code:    404,
			Type:    response2.ErrNotFound,
		}
	}

	file, contentType, err := filehandler.GetFileFromPath(ctx, voter.KTPPhotoPath, filehandler.DisplayModeInline)
	if err != nil {
		return nil, "", &custerr.ErrChain{
			Message: "Failed to get KTP photo",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	return file, contentType, nil
}

func (m *Module) GetVoterByUserID(ctx context.Context) (*response.VoterResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterUseCases.GetVoterByUserID")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	voter, err := m.voterRepo.GetVoterByUserID(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByAddress] Failed to get voter by user id")
		return nil, &custerr.ErrChain{
			Message: "Failed to get voter by address",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	voterContract, err := m.voterContract.GetVoterByAddress(nil, common.HexToAddress(voter.VoterAddress))
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByUserID] Failed to get voter by address from contract")
		return nil, &custerr.ErrChain{
			Message: "Failed to get voter by address from contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	if voter.VoterAddress != voterContract.VoterAddress.String() {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByUserID] Address mismatch between database and contract")
		return nil, &custerr.ErrChain{
			Message: "Address mismatch between database and contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
		}
	}

	return &response.VoterResponse{
		ID:                 voter.ID.String(),
		UserID:             voter.UserID.String(),
		NIK:                voter.NIK,
		FullName:           voter.FullName,
		Gender:             voter.Gender,
		BirthPlace:         voter.BirthPlace,
		BirthDate:          voter.BirthDate.Format("2006-01-02"),
		ResidentialAddress: voter.ResidentialAddress,
		VoterAddress:       voter.VoterAddress,
		Region:             voter.Region,
		IsRegistered:       voter.IsRegistered,
		HasVoted:           voter.HasVoted,
	}, err
}
