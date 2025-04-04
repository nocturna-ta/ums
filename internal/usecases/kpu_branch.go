package usecases

import (
	"context"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type KPUBranchUseCases interface {
	RegisterKPUBranch(ctx context.Context, req *request.KPUBranchRegistrationRequest) (*response.KPUBranchRegistrationResponse, error)
	GetAllKPUBranch(ctx context.Context) (*[]response.KPUBranchResponse, error)
	GetKPUBranchByAddress(ctx context.Context) (*response.KPUBranchResponse, error)
}
