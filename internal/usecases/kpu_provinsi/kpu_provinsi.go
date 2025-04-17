package kpu_provinsi

import (
	"context"
	"errors"
	libCtx "github.com/nocturna-ta/golib/context"
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
		res = append(res, response.KPUProvinsiResponse{
			ID:       kpuProvinsis.ID.String(),
			Address:  kpuProvinsis.Address,
			Region:   kpuProvinsis.Region,
			IsActive: kpuProvinsis.IsActive,
		})
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

	return &response.KPUProvinsiResponse{
		ID:       kpuProvinsi.ID.String(),
		Address:  kpuProvinsi.Address,
		Region:   kpuProvinsi.Region,
		IsActive: kpuProvinsi.IsActive,
	}, err
}
