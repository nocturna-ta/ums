package kpu_branch

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
)

func (m *Module) RegisterKPUBranch(ctx context.Context, req *request.KPUBranchRegistrationRequest) (*response.KPUBranchRegistrationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUBranchUseCases.RegisterKPUBranch")
	defer span.End()

	var (
		kpuBranch *model.KPUBranch
		err       error
	)

	kpuBranch = model.ConstructRegistrationKPUBranch(req)

	if err := m.kpuBranchRepo.InsertKPUBranch(ctx, kpuBranch); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUBranchUseCases.RegisterKPUBranch] Failed to register kpu branch")
	}

	if errors.Is(err, dao.ErrDuplicate) {
		return nil, &custerr.ErrChain{
			Message: "KPU Branch already exists",
			Code:    400,
			Cause:   err,
			Type:    response2.ErrBadRequest,
		}
	}

	return &response.KPUBranchRegistrationResponse{
		BranchAddress: kpuBranch.BranchAddress,
		IsActive:      true,
	}, err
}

func (m *Module) GetAllKPUBranch(ctx context.Context) (*[]response.KPUBranchResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUBranchUseCases.GetAllKPUBranch")
	defer span.End()

	kpuBranches, err := m.kpuBranchRepo.GetAllKPUBranch(ctx)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[KPUBranchUseCases.GetAllKPUBranch] Failed to get all kpu branch")
	}

	var res []response.KPUBranchResponse
	for _, kpuBranch := range kpuBranches {
		res = append(res, response.KPUBranchResponse{
			ID:            kpuBranch.ID.String(),
			BranchAddress: kpuBranch.BranchAddress,
			Region:        kpuBranch.Region,
			IsActive:      kpuBranch.IsActive,
		})
	}

	return &res, err
}

func (m *Module) GetKPUBranchByAddress(ctx context.Context, address string) (*response.KPUBranchResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "KPUBranchUseCases.GetKPUBranchByAddress")
	defer span.End()

	kpuBranch, err := m.kpuBranchRepo.GetKPUBranchByAddress(ctx, address)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"address": address,
		}).ErrorWithCtx(ctx, "[KPUBranchUseCases.GetKPUBranchByAddress] Failed to get kpu branch by address")
	}

	return &response.KPUBranchResponse{
		ID:            kpuBranch.ID.String(),
		BranchAddress: kpuBranch.BranchAddress,
		Region:        kpuBranch.Region,
		IsActive:      kpuBranch.IsActive,
	}, err
}
