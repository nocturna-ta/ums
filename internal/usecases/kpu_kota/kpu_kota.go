package kpu_kota

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

func (m *Module) RegisterKPUKota(ctx context.Context, req *request.KPUKotaRegistrationRequest) (*response.KPUKotaRegistrationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.RegisterKPUKota")
	defer span.End()

	var (
		kpuKota *model.KPUKota
	)

	transaction := func(txCtx context.Context) (any, error) {
		kpuKota = model.ConstructRegistrationKPUKota(req)

		errTx := m.kpuKotaRepo.InsertKPUKota(txCtx, kpuKota, req.SignedTransaction)
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

		errTx = m.publisher.Publish(txCtx, m.topics.MasterDataKPUKota.Value, kpuKota.ID.String(), kpuKota.ToMessageModel(), map[string]any{
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

	return &response.KPUKotaRegistrationResponse{
		Address:  kpuKota.Address,
		IsActive: true,
	}, err
}

func (m *Module) GetAllKPUKota(ctx context.Context) (*[]response.KPUKotaResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUKotaUseCases.GetAllKPUKota")
	defer span.End()

	kpuKotaList, err := m.kpuKotaRepo.GetAllKPUKota(ctx)
	if err != nil {
		return nil, err
	}

	var kpuKotaResponse []response.KPUKotaResponse
	for _, kpuKota := range kpuKotaList {
		kpuKotaResponse = append(kpuKotaResponse, response.KPUKotaResponse{
			ID:       kpuKota.ID.String(),
			Address:  kpuKota.Address,
			Region:   kpuKota.Region,
			IsActive: kpuKota.IsActive,
		})
	}

	return &kpuKotaResponse, nil
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
	}

	return &response.KPUKotaResponse{
		ID:       kpuKota.ID.String(),
		Address:  kpuKota.Address,
		Region:   kpuKota.Region,
		IsActive: kpuKota.IsActive,
	}, err

}
