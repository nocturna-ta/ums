package usecases

import (
	"context"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type KPUKotaUseCases interface {
	RegisterKPUKota(ctx context.Context, req *request.KPUKotaRegistrationRequest) (*response.KPUKotaRegistrationResponse, error)
	GetAllKPUKota(ctx context.Context) (*[]response.KPUKotaResponse, error)
	GetKPUKotaByAddress(ctx context.Context) (*response.KPUKotaResponse, error)
}
