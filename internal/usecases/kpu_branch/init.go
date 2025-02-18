package kpu_branch

import (
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	kpuBranchRepo repository.KPUBranchRepository
	jwtSvc        jwtsvc.JWT
	TxMgr         txmanager.TxManager
}

type Opts struct {
	KpuBranchRepo repository.KPUBranchRepository
	TxMgr         txmanager.TxManager
	JwtSvc        jwtsvc.JWT
}

func New(opts *Opts) usecases.KPUBranchUseCases {
	return &Module{
		kpuBranchRepo: opts.KpuBranchRepo,
		jwtSvc:        opts.JwtSvc,
		TxMgr:         opts.TxMgr,
	}
}
