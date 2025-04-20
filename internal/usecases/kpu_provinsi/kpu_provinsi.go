package kpu_provinsi

import (
	"context"
	"errors"
	"fmt"
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

		errTx := m.kpuProvinsiRepo.InsertKPUProvinsi(txCtx, kpuProvinsi, req.SignedTransaction)
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
	}

	var res []response.KPUProvinsiResponse
	for _, kpuProvinsis := range kpuProvinsi {
		resp := response.KPUProvinsiResponse{
			ID:       kpuProvinsis.ID.String(),
			Address:  kpuProvinsis.Address,
			Region:   kpuProvinsis.Region,
			IsActive: kpuProvinsis.IsActive,
			Name:     kpuProvinsis.Name,
		}

		if kpuProvinsis.PhotoPath != "" {
			resp.PhotoURL = fmt.Sprintf("/v1/kpu-provinsi/%s/photo", kpuProvinsis.ID.String())
		}

		res = append(res, resp)
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
	}

	res := &response.KPUProvinsiResponse{
		ID:       kpuProvinsi.ID.String(),
		Address:  kpuProvinsi.Address,
		Region:   kpuProvinsi.Region,
		IsActive: kpuProvinsi.IsActive,
		Name:     kpuProvinsi.Name,
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
		return nil, err
	}

	res := &response.KPUProvinsiResponse{
		ID:       kpuProvinsi.ID.String(),
		Address:  kpuProvinsi.Address,
		Region:   kpuProvinsi.Region,
		IsActive: kpuProvinsi.IsActive,
		Name:     kpuProvinsi.Name,
	}

	if kpuProvinsi.PhotoPath != "" {
		res.PhotoURL = fmt.Sprintf("/v1/kpu-provinsi/%s/photo", kpuProvinsi.ID.String())
	}

	return res, nil
}

func (m *Module) UploadKPUProvinsiPhoto(ctx context.Context, kpuProvinsiID uuid.UUID, fileData io.Reader, fileName string) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.UploadKPUProvinsiPhoto")
	defer span.End()

	kpuProvinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByID(ctx, kpuProvinsiID)
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
			"error": err,
			"id":    kpuProvinsiID,
		}).ErrorWithCtx(ctx, "[KPUProvinsiUseCases.UploadKPUProvinsiPhoto] Failed to store file")
		return &custerr.ErrChain{
			Message: "Failed to upload photo",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}
	}

	err = m.kpuProvinsiRepo.UpdateKPUProvinsiPhoto(ctx, kpuProvinsiID, photoPath)
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

func (m *Module) GetKPUProvinsiPhoto(ctx context.Context, kpuProvinsiID uuid.UUID) (*http.File, string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUProvinsiUseCases.GetKPUProvinsiPhoto")
	defer span.End()

	provinsi, err := m.kpuProvinsiRepo.GetKPUProvinsiByID(ctx, kpuProvinsiID)
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

	file, contentType, err := filehandler.GetFileFromPath(ctx, provinsi.PhotoPath)
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
