package user_statistic

import (
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/infrastructures/wilayah"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	userStatisticRepo repository.UserStatisticRepository
	kpuProvinsiRepo   repository.KPUProvinsiRepository
	kpuKotaRepo       repository.KPUKotaRepository
	wilayahAPIClient  *wilayah.WilayahAPIClient
}

type Opts struct {
	UserStatisticRepo repository.UserStatisticRepository
	KPUProvinsiRepo   repository.KPUProvinsiRepository
	KPUKotaRepo       repository.KPUKotaRepository
	WilayahAPIClient  *wilayah.WilayahAPIClient
}

func New(opts *Opts) usecases.UserStatisticUseCases {
	wilayahClient := opts.WilayahAPIClient
	if wilayahClient == nil {
		wilayahClient = wilayah.NewWilayahAPIClient()
	}
	return &Module{
		userStatisticRepo: opts.UserStatisticRepo,
		kpuProvinsiRepo:   opts.KPUProvinsiRepo,
		kpuKotaRepo:       opts.KPUKotaRepo,
		wilayahAPIClient:  wilayahClient,
	}
}
