package voter

import (
	"context"
	"errors"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/log"
	response2 "github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/ums/pkg/constants"
)

func (m *Module) RegisterVoter(ctx context.Context, req *request.VoterRegistrationRequest) (*response.VoterRegistrationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterUseCases.RegisterVoter")
	defer span.End()

	var (
		voter *model.Voter
	)

	transaction := func(txCtx context.Context) (any, error) {
		voter = model.ConstructRegistration(req)

		errTx := m.voterRepo.InsertVoter(txCtx, voter, req.SignedTransaction)
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

		errTx = m.publisher.Publish(txCtx, m.topics.MasterDataVoter.Value, voter.ID.String(), voter.ToMessageModel(), map[string]any{
			constants.MetaDataOperation: constants.Create,
		})
		if errTx != nil {
			return nil, errTx
		}

		return nil, nil
	}

	_, err := m.txMgr.Execute(ctx, transaction, nil)
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
	}

	return &response.VoterResponse{
		ID:           voter.ID.String(),
		NIK:          voter.NIK,
		VoterAddress: voter.VoterAddress,
		Region:       voter.Region,
		IsRegistered: voter.IsRegistered,
		HasVoted:     voter.HasVoted,
	}, err
}

func (m *Module) GetVoterByAddress(ctx context.Context, address string) (*response.VoterResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "VoterUseCases.GetVoterByAddress")
	defer span.End()

	voter, err := m.voterRepo.GetVoterByAddress(ctx, address)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": address,
		}).ErrorWithCtx(ctx, "[VoterUseCases.GetVoterByAddress] Failed to get voter by address")
	}

	return &response.VoterResponse{
		ID:           voter.ID.String(),
		NIK:          voter.NIK,
		VoterAddress: voter.VoterAddress,
		Region:       voter.Region,
		IsRegistered: voter.IsRegistered,
		HasVoted:     voter.HasVoted,
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
	}

	var res []response.VoterResponse
	for _, voter := range voters {
		res = append(res, response.VoterResponse{
			ID:           voter.ID.String(),
			NIK:          voter.NIK,
			VoterAddress: voter.VoterAddress,
			Region:       voter.Region,
			IsRegistered: voter.IsRegistered,
			HasVoted:     voter.HasVoted,
		})
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

	var res []response.VoterResponse
	for _, voter := range voters {
		res = append(res, response.VoterResponse{
			ID:           voter.ID.String(),
			NIK:          voter.NIK,
			VoterAddress: voter.VoterAddress,
			Region:       voter.Region,
			IsRegistered: voter.IsRegistered,
			HasVoted:     voter.HasVoted,
		})
	}

	return &res, err
}
