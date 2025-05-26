package kpu_provinsi

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/fileutils"
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
	"io"
)

func (m *Module) RegisterKPUProvinsi(ctx context.Context, req *request.KPUProvinsiRegistrationRequest) (*response.KPUProvinsiRegistrationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.RegisterKPUProvinsi")
	defer span.End()

	var (
		kpuProvinsi *model.KPUProvinsi
	)

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	transaction := func(txCtx context.Context) (any, error) {
		kpuProvinsi = model.ConstructRegistrationKPUProvinsi(req)
		kpuProvinsi.UserID = reqCtx.GetUserId()

		errTx := m.kpuProvinsiRepo.InsertKPUProvinsi(txCtx, kpuProvinsi)
		if errTx != nil {
			if errors.Is(errTx, dao.ErrDuplicate) {
				return nil, &custerr.ErrChain{
					Message: "KPU Provinsi already exists",
					Code:    400,
					Type:    response2.ErrBadRequest,
					Cause:   errTx,
				}
			}

			return nil, errTx
		}

		txHash, errTx := m.kpuProvinsiRepo.SendTxKPUProvinsiBlockchain(txCtx, req.SignedTransaction)
		if errTx != nil {
			log.WithFields(log.Fields{
				"error": errTx,
				"id":    kpuProvinsi.ID,
			}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.RegisterKPUProvinsi] Failed to send transaction to blockchain")
			return nil, errTx
		}

		if txHash == "" {
			return nil, err
		}

		errTx = m.publisher.Publish(txCtx, m.topics.MasterDataKPUProvinsi.Value, kpuProvinsi.ID.String(), kpuProvinsi.ToMessageModel(), map[string]any{
			constants.MetaDataOperation: constants.Create,
		})

		if errTx != nil {
			return nil, errTx
		}

		return nil, nil
	}

	_, err = m.txMgr.Execute(ctx, transaction, nil)
	if err != nil {
		return nil, err
	}

	return &response.KPUProvinsiRegistrationResponse{
		ID:       kpuProvinsi.ID.String(),
		Address:  kpuProvinsi.Address,
		IsActive: true,
	}, err
}

func (m *Module) GetAllKPUProvinsi(ctx context.Context) (*[]response.KPUProvinsiResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.GetAllKPUProvinsi")
	defer span.End()

	kpuProvinsi, err := m.kpuProvinsiRepo.GetAllKPUProvinsi(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetAllKPUProvinsi] Failed to get all kpu provinsi")
		return nil, &custerr.ErrChain{
			Message: "Failed to get all KPU Provinsi",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	kpuProvinsiContract, err := m.kpuContract.GetAllKPUProvinsi(nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetAllKPUProvinsi] Failed to get all kpu provinsi from contract")
		return nil, &custerr.ErrChain{
			Message: "Failed to get all KPU Provinsi from contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	contractAddress := make(map[string]bool)
	for _, contract := range kpuProvinsiContract {
		contractAddress[contract.Address.String()] = true
	}

	var res []response.KPUProvinsiResponse
	for _, kpuProvinsis := range kpuProvinsi {
		if contractAddress[kpuProvinsis.Address] {
			resp := response.KPUProvinsiResponse{
				ID:           kpuProvinsis.ID.String(),
				UserID:       kpuProvinsis.UserID.String(),
				Name:         kpuProvinsis.Name,
				Address:      kpuProvinsis.Address,
				Region:       kpuProvinsis.Region,
				IsActive:     kpuProvinsis.IsActive,
				PhotoURL:     kpuProvinsis.PhotoPath,
				Telephone:    kpuProvinsis.Telephone,
				RegisteredAt: kpuProvinsis.RegisteredAt.String(),
			}

			if kpuProvinsis.PhotoPath != "" {
				resp.PhotoURL = fmt.Sprintf("/v1/kpu-provinsi/%s/photo", kpuProvinsis.ID.String())
			}

			res = append(res, resp)
		}
	}

	return &res, err
}

func (m *Module) GetKPUProvinsiByAddress(ctx context.Context) (*response.KPUProvinsiResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.GetKPUProvinsiByAddress")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	kpuProvinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByAddress(ctx, reqCtx.GetAddress())
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetKPUProvinsiByAddress] Failed to get kpu provinsi by address")
		return nil, &custerr.ErrChain{
			Message: "Failed to get KPU Provinsi by address",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	kpuProvinsiContract, err := m.kpuContract.GetKpuProvinsiByAddress(nil, common.HexToAddress(reqCtx.GetAddress()))
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetKPUProvinsiByAddress] Failed to get kpu provinsi from contract")
		return nil, &custerr.ErrChain{
			Message: "Failed to get KPU Provinsi from contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	if kpuProvinsiContract.Address.String() != kpuProvinsi.Address {
		log.WithFields(log.Fields{
			"address":          kpuProvinsi.Address,
			"contract_address": kpuProvinsiContract.Address.String(),
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetKPUProvinsiByAddress] Address mismatch between contract and database")
		return nil, &custerr.ErrChain{
			Message: "Address mismatch between contract and database",
			Code:    500,
			Type:    response2.ErrInternalServerError,
		}
	}

	res := &response.KPUProvinsiResponse{
		ID:           kpuProvinsi.ID.String(),
		UserID:       kpuProvinsi.UserID.String(),
		Name:         kpuProvinsi.Name,
		Address:      kpuProvinsi.Address,
		Region:       kpuProvinsi.Region,
		IsActive:     kpuProvinsi.IsActive,
		PhotoURL:     kpuProvinsi.PhotoPath,
		Telephone:    kpuProvinsi.Telephone,
		RegisteredAt: kpuProvinsi.RegisteredAt.String(),
	}

	if kpuProvinsi.PhotoPath != "" {
		res.PhotoURL = fmt.Sprintf("/v1/kpu-provinsi/%s/photo", kpuProvinsi.ID.String())
	}

	return res, err
}

func (m *Module) GetKPUProvinsiByID(ctx context.Context, id uuid.UUID) (*response.KPUProvinsiResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.GetKPUProvinsiByID")
	defer span.End()

	kpuProvinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByID(ctx, id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetKPUProvinsiByID] Failed to get kpu provinsi by ID")
		return nil, &custerr.ErrChain{
			Message: "KPU Provinsi not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	kpuProvinsiContract, err := m.kpuContract.GetKpuProvinsiByAddress(nil, common.HexToAddress(kpuProvinsi.Address))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetKPUProvinsiByID] Failed to get kpu provinsi from contract")
		return nil, &custerr.ErrChain{
			Message: "Failed to get KPU Provinsi from contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	if kpuProvinsiContract.Address.String() != kpuProvinsi.Address {
		log.WithFields(log.Fields{
			"address":          kpuProvinsi.Address,
			"contract_address": kpuProvinsiContract.Address.String(),
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetKPUProvinsiByID] Address mismatch between contract and database")
		return nil, &custerr.ErrChain{
			Message: "Address mismatch between contract and database",
			Code:    500,
			Type:    response2.ErrInternalServerError,
		}
	}

	res := &response.KPUProvinsiResponse{
		ID:           kpuProvinsi.ID.String(),
		UserID:       kpuProvinsi.UserID.String(),
		Name:         kpuProvinsi.Name,
		Address:      kpuProvinsi.Address,
		Region:       kpuProvinsi.Region,
		IsActive:     kpuProvinsi.IsActive,
		PhotoURL:     kpuProvinsi.PhotoPath,
		Telephone:    kpuProvinsi.Telephone,
		RegisteredAt: kpuProvinsi.RegisteredAt.String(),
	}

	if kpuProvinsi.PhotoPath != "" {
		res.PhotoURL = fmt.Sprintf("/v1/kpu-provinsi/%s/photo", kpuProvinsi.ID.String())
	}

	return res, nil
}

func (m *Module) UploadKPUProvinsiPhoto(ctx context.Context, fileData io.Reader, fileName string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.UploadKPUProvinsiPhoto")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return err
	}

	kpuProvinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByUserID(ctx, reqCtx.GetUserId())
	if err != nil {
		return &custerr.ErrChain{
			Message: "KPU Provinsi not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	if kpuProvinsi.PhotoPath != "" {
		_ = fileutils.DeleteFile(ctx, kpuProvinsi.PhotoPath)
	}

	fileConfig := fileutils.DefaultConfig()
	fileConfig.SetAllowedImageExtension()
	fileConfig.EntityType = "kpu_provinsi"

	photoPath, err := fileutils.StoreFile(ctx, fileData, fileName, fileConfig)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"id":      kpuProvinsi.ID,
			"user_id": reqCtx.GetUserId(),
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.UploadKPUProvinsiPhoto] Failed to store file")
		return &custerr.ErrChain{
			Message: "Failed to upload photo",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	err = m.kpuProvinsiRepo.UpdateKPUProvinsiPhoto(ctx, kpuProvinsi.ID, photoPath)
	if err != nil {
		_ = fileutils.DeleteFile(ctx, photoPath)
		return &custerr.ErrChain{
			Message: "Failed to update KPU Provinsi with photo information",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}
	return nil
}

func (m *Module) GetKPUProvinsiPhoto(ctx context.Context) (*http.File, string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.GetKPUProvinsiPhoto")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, "", err
	}

	provinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByUserID(ctx, reqCtx.GetUserId())
	if err != nil {
		return nil, "", &custerr.ErrChain{
			Message: "KPU Provinsi not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	if provinsi.PhotoPath == "" {
		return nil, "", &custerr.ErrChain{
			Message: "No photo available for this KPU Provinsi",
			Code:    404,
			Type:    response2.ErrNotFound,
		}
	}

	file, contentType, err := filehandler.GetFileFromPath(ctx, provinsi.PhotoPath, filehandler.DisplayModeInline)
	if err != nil {
		return nil, "", &custerr.ErrChain{
			Message: "Failed to retrieve photo",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	return file, contentType, nil
}

func (m *Module) UpdateKPUProvinsi(ctx context.Context, updateRequest *request.KPUProvinsiUpdateRequest) (*response.KPUProvinsiResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.UpdateKPUProvinsi")
	defer span.End()

	var (
		existing *model.KPUProvinsi
	)

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	transaction := func(txCtx context.Context) (any, error) {
		existing, errTx := m.kpuProvinsiRepo.GetKPUProvinsiByAddress(ctx, reqCtx.GetAddress())
		if errTx != nil {
			return nil, &custerr.ErrChain{
				Message: "KPU Provinsi not found",
				Code:    404,
				Type:    response2.ErrNotFound,
				Cause:   errTx,
			}
		}
		existing.Name = updateRequest.Name
		existing.Telephone = updateRequest.Telephone
		existing.Region = updateRequest.Region

		errTx = m.kpuProvinsiRepo.UpdateKPUProvinsi(ctx, existing)
		if err != nil {
			log.WithFields(log.Fields{
				"error": errTx,
				"id":    existing.ID,
			}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.UpdateKPUProvinsi] Failed to update kpu provinsi")
			return nil, errTx
		}

		txHash, errTx := m.kpuProvinsiRepo.SendTxKPUProvinsiBlockchain(ctx, updateRequest.SignedTransaction)
		if errTx != nil {
			log.WithFields(log.Fields{
				"error": errTx,
				"id":    existing.ID,
			}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.UpdateKPUProvinsi] Failed to send transaction to blockchain")
			return nil, errTx
		}

		if txHash == "" {
			return nil, errTx
		}

		errTx = m.publisher.Publish(ctx, m.topics.MasterDataKPUProvinsi.Value, existing.ID.String(), existing.ToMessageModel(), map[string]any{
			constants.MetaDataOperation: constants.Update,
		})

		if errTx != nil {
			return nil, errTx
		}

		return nil, nil

	}

	_, err = m.txMgr.Execute(ctx, transaction, nil)
	if err != nil {
		return nil, err
	}

	res := &response.KPUProvinsiResponse{
		ID:           existing.ID.String(),
		UserID:       existing.UserID.String(),
		Name:         existing.Name,
		Address:      reqCtx.Address,
		Region:       existing.Region,
		IsActive:     existing.IsActive,
		Telephone:    existing.Telephone,
		RegisteredAt: existing.RegisteredAt.String(),
		PhotoURL:     existing.PhotoPath,
	}

	return res, nil
}

func (m *Module) GetKPUProvinsiByUserID(ctx context.Context) (*response.KPUProvinsiResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.GetKPUProvinsiByUserID")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	kpuProvinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByUserID(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"id":      kpuProvinsi.ID,
			"user_id": reqCtx.GetUserId(),
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetKPUProvinsiByUserID] Failed to get kpu provinsi by User ID")
		return nil, &custerr.ErrChain{
			Message: "KPU Provinsi not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	kpuProvinsiContract, err := m.kpuContract.GetKpuProvinsiByAddress(nil, common.HexToAddress(kpuProvinsi.Address))
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"id":      kpuProvinsi.ID,
			"user_id": reqCtx.GetUserId(),
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetKPUProvinsiByID] Failed to get kpu provinsi from contract")
		return nil, &custerr.ErrChain{
			Message: "Failed to get KPU Provinsi from contract",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	if kpuProvinsiContract.Address.String() != kpuProvinsi.Address {
		log.WithFields(log.Fields{
			"address":          kpuProvinsi.Address,
			"contract_address": kpuProvinsiContract.Address.String(),
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.GetKPUProvinsiByID] Address mismatch between contract and database")
		return nil, &custerr.ErrChain{
			Message: "Address mismatch between contract and database",
			Code:    500,
			Type:    response2.ErrInternalServerError,
		}
	}

	res := &response.KPUProvinsiResponse{
		ID:           kpuProvinsi.ID.String(),
		UserID:       kpuProvinsi.UserID.String(),
		Name:         kpuProvinsi.Name,
		Address:      kpuProvinsi.Address,
		Region:       kpuProvinsi.Region,
		IsActive:     kpuProvinsi.IsActive,
		PhotoURL:     kpuProvinsi.PhotoPath,
		Telephone:    kpuProvinsi.Telephone,
		RegisteredAt: kpuProvinsi.RegisteredAt.String(),
	}

	if kpuProvinsi.PhotoPath != "" {
		res.PhotoURL = fmt.Sprintf("/v1/kpu-provinsi/%s/photo", kpuProvinsi.ID.String())
	}

	return res, nil
}
