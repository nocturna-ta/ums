package kpu_kota

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

func (m *Module) RegisterKPUKota(ctx context.Context, req *request.KPUKotaRegistrationRequest) (*response.KPUKotaRegistrationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.RegisterKPUKota")
	defer span.End()

	var (
		kpuKota *model.KPUKota
	)

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	transaction := func(txCtx context.Context) (any, error) {
		kpuKota = model.ConstructRegistrationKPUKota(req)
		kpuKota.UserID = reqCtx.GetUserId()

		errTx := m.kpuKotaRepo.InsertKPUKota(txCtx, kpuKota)
		if errTx != nil {
			if errors.Is(errTx, dao.ErrDuplicate) {
				return nil, &custerr.ErrChain{
					Message: "KPU Kota already exists",
					Code:    400,
					Type:    response2.ErrBadRequest,
					Cause:   errTx,
				}
			}

			return nil, errTx
		}

		txHash, err := m.kpuKotaRepo.SendTxKPUKotaBlockchain(txCtx, req.SignedTransaction)

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"id":    kpuKota.ID,
				"req":   req,
			}).ErrorWithCtx(ctx, "[KPUKotaUseCases.RegisterKPUKota] Failed to send transaction to blockchain")
			return nil, err
		}

		if txHash == "" {
			return nil, err
		}

		errTx = m.publisher.Publish(txCtx, m.topics.MasterDataKPUKota.Value, kpuKota.ID.String(), kpuKota.ToMessageModel(), map[string]any{
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

	return &response.KPUKotaRegistrationResponse{
		Address:  kpuKota.Address,
		IsActive: true,
	}, err
}

func (m *Module) GetAllKPUKota(ctx context.Context) (*[]response.KPUKotaResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.GetAllKPUKota")
	defer span.End()

	_, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	kpuKotaList, err := m.kpuKotaRepo.GetAllKPUKota(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetAllKPUKota] Failed to get all KPU Kota")
		return nil, &custerr.ErrChain{
			Message: "Failed to get all KPU Kota",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	kpuKotaContract, err := m.kpuContract.GetAllKPUKota(nil)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetAllKPUKota] Failed to get all KPU Kota contracts")
		return nil, &custerr.ErrChain{
			Message: "Failed to get KPU Kota contracts",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	contractAddresses := make(map[string]bool)
	for _, contract := range kpuKotaContract {
		contractAddresses[contract.Address.String()] = true
	}

	var res []response.KPUKotaResponse
	for _, kpuKota := range kpuKotaList {
		if contractAddresses[kpuKota.Address] {
			resp := response.KPUKotaResponse{
				ID:           kpuKota.ID.String(),
				UserID:       kpuKota.UserID.String(),
				Address:      kpuKota.Address,
				Username:     kpuKota.Username,
				Name:         kpuKota.Name,
				Region:       kpuKota.Region,
				IsActive:     kpuKota.IsActive,
				PhotoURL:     kpuKota.PhotoPath,
				Telephone:    kpuKota.Telephone,
				RegisteredAt: kpuKota.RegisteredAt.String(),
			}

			if kpuKota.PhotoPath != "" {
				resp.PhotoURL = fmt.Sprintf("/v1/kpu-kota/%s/photo", kpuKota.ID.String())
			}

			res = append(res, resp)
		}
	}

	return &res, nil
}

func (m *Module) GetKPUKotaByAddress(ctx context.Context) (*response.KPUKotaResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.GetKPUKotaByAddress")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	kpuKota, err := m.kpuKotaRepo.GetKPUKotaByAddress(ctx, reqCtx.GetAddress())
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetKPUKotaByAddress] Failed to get kpu kota by address")
		return nil, &custerr.ErrChain{
			Message: "KPU Kota not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	kpuKotaContract, err := m.kpuContract.GetKpuKotaByAddress(nil, common.HexToAddress(reqCtx.GetAddress()))
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": reqCtx.Address,
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetKPUKotaByAddress] Failed to get kpu kota contact by address")
		return nil, &custerr.ErrChain{
			Message: "KPU Kota in contract not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	if kpuKota.Address != kpuKotaContract.Address.String() {
		log.WithFields(log.Fields{
			"repo_address":     kpuKota.Address,
			"contract_address": kpuKotaContract.Address.String(),
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetKPUKotaByAddress] Address mismatch between repo and contract")
		return nil, &custerr.ErrChain{
			Message: "Address mismatch between repository and contract data",
			Code:    404,
			Type:    response2.ErrNotFound,
		}
	}

	res := &response.KPUKotaResponse{
		ID:           kpuKota.ID.String(),
		UserID:       kpuKota.UserID.String(),
		Address:      kpuKota.Address,
		Username:     kpuKota.Username,
		Name:         kpuKota.Name,
		Region:       kpuKota.Region,
		IsActive:     kpuKota.IsActive,
		PhotoURL:     kpuKota.PhotoPath,
		Telephone:    kpuKota.Telephone,
		RegisteredAt: kpuKota.RegisteredAt.String(),
	}

	if kpuKota.PhotoPath != "" {
		res.PhotoURL = fmt.Sprintf("/v1/kpu-kota/%s/photo", kpuKota.ID.String())
	}

	return res, err
}

func (m *Module) GetKPUKotaByID(ctx context.Context, id uuid.UUID) (*response.KPUKotaResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.GetKPUKotaByID")
	defer span.End()

	kpuKota, err := m.kpuKotaRepo.GetKPUKotaByID(ctx, id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetKPUKotaByID] Failed to get kpu kota by ID")
		return nil, err
	}

	kpuKotaContract, err := m.kpuContract.GetKpuKotaByAddress(nil, common.HexToAddress(kpuKota.Address))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    id,
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetKPUKotaByID] Failed to get kpu kota contract by ID")
		return nil, &custerr.ErrChain{
			Message: "KPU Kota in contract not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	if kpuKota.Address != kpuKotaContract.Address.String() {
		log.WithFields(log.Fields{
			"repo_address":     kpuKota.Address,
			"contract_address": kpuKotaContract.Address.String(),
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetKPUKotaByAddress] Address mismatch between repo and contract")
		return nil, &custerr.ErrChain{
			Message: "Address mismatch between repository and contract data",
			Code:    404,
			Type:    response2.ErrNotFound,
		}
	}

	res := &response.KPUKotaResponse{
		ID:           kpuKota.ID.String(),
		UserID:       kpuKota.UserID.String(),
		Address:      kpuKota.Address,
		Username:     kpuKota.Username,
		Name:         kpuKota.Name,
		Region:       kpuKota.Region,
		IsActive:     kpuKota.IsActive,
		PhotoURL:     kpuKota.PhotoPath,
		Telephone:    kpuKota.Telephone,
		RegisteredAt: kpuKota.RegisteredAt.String(),
	}

	if kpuKota.PhotoPath != "" {
		res.PhotoURL = fmt.Sprintf("/v1/kpu-kota/%s/photo", kpuKota.ID.String())
	}

	return res, nil
}

func (m *Module) UploadKPUKotaPhoto(ctx context.Context, fileData io.Reader, fileName string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.UploadKPUKotaPhoto")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return err
	}

	kpuKota, err := m.kpuKotaRepo.GetKPUKotaByUserID(ctx, reqCtx.GetUserId())
	if err != nil {
		return &custerr.ErrChain{
			Message: "KPU Provinsi not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	if kpuKota.PhotoPath != "" {
		_ = fileutils.DeleteFile(ctx, kpuKota.PhotoPath)
	}

	fileConfig := fileutils.DefaultConfig()
	fileConfig.SetAllowedImageExtension()
	fileConfig.EntityType = "kpu_kota"

	photoPath, err := fileutils.StoreFile(ctx, fileData, fileName, fileConfig)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"id":      kpuKota.ID,
			"user_id": reqCtx.GetUserId(),
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.UploadKPUKotaPhoto] Failed to store file")
		return &custerr.ErrChain{
			Message: "Failed to upload photo",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	err = m.kpuKotaRepo.UpdateKPUKotaPhoto(ctx, kpuKota.ID, photoPath)
	if err != nil {
		_ = fileutils.DeleteFile(ctx, photoPath)
		return &custerr.ErrChain{
			Message: "Failed to update KPU Kota with photo information",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}
	return nil
}

func (m *Module) GetKPUKotaPhoto(ctx context.Context) (*http.File, string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.GetKPUKotaPhoto")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, "", err
	}

	kota, err := m.kpuKotaRepo.GetKPUKotaByUserID(ctx, reqCtx.GetUserId())
	if err != nil {
		return nil, "", &custerr.ErrChain{
			Message: "KPU Kota not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	if kota.PhotoPath == "" {
		return nil, "", &custerr.ErrChain{
			Message: "No photo available for this KPU Kota",
			Code:    404,
			Type:    response2.ErrNotFound,
		}
	}

	file, contentType, err := filehandler.GetFileFromPath(ctx, kota.PhotoPath, filehandler.DisplayModeInline)
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

func (m *Module) GetKPUKotaPhotoUseID(ctx context.Context, id uuid.UUID) (*http.File, string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.GetKPUProvinsiPhoto")
	defer span.End()

	_, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, "", err
	}

	provinsi, err := m.kpuKotaRepo.GetKPUKotaByUserID(ctx, id)
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

func (m *Module) UpdateKPUKota(ctx context.Context, updateRequest *request.KPUKotaUpdateRequest) (*response.KPUKotaResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.UpdateKPUKota")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	var (
		existing model.KPUKota
	)

	transaction := func(txCtx context.Context) (any, error) {
		existing, errTx := m.kpuKotaRepo.GetKPUKotaByUserID(ctx, reqCtx.GetUserId())
		if errTx != nil {
			if errors.Is(errTx, dao.ErrNoUpdateHappened) {
				return nil, &custerr.ErrChain{
					Message: "No update happened, KPU Provinsi may not exist",
					Code:    404,
					Type:    response2.ErrNotFound,
				}
			}
			return nil, errTx
		}
		existing.Name = updateRequest.Name
		existing.Region = updateRequest.Region
		existing.Telephone = updateRequest.Telephone

		errTx = m.kpuKotaRepo.UpdateKPUKota(ctx, existing)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"id":    existing.ID,
				"req":   updateRequest,
			}).ErrorWithCtx(ctx, "[KPUKotaUseCases.UpdateKPUKota] Failed to update kpu kota")
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

	return &response.KPUKotaResponse{
		ID:           existing.ID.String(),
		UserID:       existing.UserID.String(),
		Address:      reqCtx.Address,
		Username:     existing.Username,
		Name:         existing.Name,
		Region:       existing.Region,
		IsActive:     existing.IsActive,
		PhotoURL:     existing.PhotoPath,
		Telephone:    existing.Telephone,
		RegisteredAt: existing.RegisteredAt.String(),
	}, nil
}

func (m *Module) GetKPUKotaByUserID(ctx context.Context) (*response.KPUKotaResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.GetKPUKotaByUserID")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	kpuKota, err := m.kpuKotaRepo.GetKPUKotaByUserID(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"id":      kpuKota.ID,
			"user_id": reqCtx.GetUserId(),
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetKPUKotaByUserID] Failed to get kpu kota by ID")
		return nil, &custerr.ErrChain{
			Message: "KPU Kota not found in database",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	kpuKotaContract, err := m.kpuContract.GetKpuKotaByAddress(nil, common.HexToAddress(kpuKota.Address))
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"id":      kpuKota.ID,
			"user_id": reqCtx.GetUserId(),
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetKPUKotaByUserID] Failed to get kpu kota contract by User ID")
		return nil, &custerr.ErrChain{
			Message: "KPU Kota in contract not found",
			Code:    404,
			Type:    response2.ErrNotFound,
			Cause:   err,
		}
	}

	if kpuKota.Address != kpuKotaContract.Address.String() {
		log.WithFields(log.Fields{
			"repo_address":     kpuKota.Address,
			"contract_address": kpuKotaContract.Address.String(),
		}).ErrorWithCtx(ctx, "[KPUKotaUseCases.GetKPUKotaByUserID] Address mismatch between repo and contract")
		return nil, &custerr.ErrChain{
			Message: "Address mismatch between repository and contract data",
			Code:    404,
			Type:    response2.ErrNotFound,
		}
	}

	res := &response.KPUKotaResponse{
		ID:           kpuKota.ID.String(),
		UserID:       kpuKota.UserID.String(),
		Username:     kpuKota.Username,
		Address:      kpuKota.Address,
		Name:         kpuKota.Name,
		Region:       kpuKota.Region,
		IsActive:     kpuKota.IsActive,
		PhotoURL:     kpuKota.PhotoPath,
		Telephone:    kpuKota.Telephone,
		RegisteredAt: kpuKota.RegisteredAt.String(),
	}

	if kpuKota.PhotoPath != "" {
		res.PhotoURL = fmt.Sprintf("/v1/kpu-kota/%s/photo", kpuKota.ID.String())
	}

	return res, nil
}
