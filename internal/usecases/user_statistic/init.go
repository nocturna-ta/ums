package user_statistic

import (
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	userStatisticRepo repository.UserStatisticRepository
	kpuProvinsiRepo   repository.KPUProvinsiRepository
	kpuKotaRepo       repository.KPUKotaRepository
}

type Opts struct {
	UserStatisticRepo repository.UserStatisticRepository
	KPUProvinsiRepo   repository.KPUProvinsiRepository
	KPUKotaRepo       repository.KPUKotaRepository
}

func New(opts *Opts) usecases.UserStatisticUseCases {
	return &Module{
		userStatisticRepo: opts.UserStatisticRepo,
		kpuProvinsiRepo:   opts.KPUProvinsiRepo,
		kpuKotaRepo:       opts.KPUKotaRepo,
	}
}
