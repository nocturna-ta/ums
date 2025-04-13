package usecases

import (
	"context"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type KPUProvinsiUseCases interface {
	RegisterKPUProvinsi(ctx context.Context, req *request.KPUProvinsiRegistrationRequest) (*response.KPUProvinsiRegistrationResponse, error)
	GetAllKPUProvinsi(ctx context.Context) (*[]response.KPUProvinsiResponse, error)
	GetKPUProvinsiByAddress(ctx context.Context) (*response.KPUProvinsiResponse, error)
}
