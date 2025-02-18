package repository

import (
	"context"
	"github.com/nocturna-ta/ums/internal/domain/model"
)

type KPUBranchRepository interface {
	InsertKPUBranch(ctx context.Context, kpuBranch *model.KPUBranch) error
	GetAllKPUBranch(ctx context.Context) ([]model.KPUBranch, error)
	GetKPUBranchByAddress(ctx context.Context, address string) (*model.KPUBranch, error)
}
