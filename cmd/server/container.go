package server

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/txmanager"
	txSql "github.com/nocturna-ta/golib/txmanager/sql"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
	"github.com/nocturna-ta/ums/internal/usecases/kpu_branch"
	"github.com/nocturna-ta/ums/internal/usecases/voter"
)

type container struct {
	Cfg         config.MainConfig
	VoterUc     usecases.VoterUseCases
	KpuBranchUc usecases.KPUBranchUseCases
}

type options struct {
	Cfg    *config.MainConfig
	DB     *sql.Store
	Client *ethclient.Client
}

func newContainer(opts *options) *container {
	voterRepo := dao.NewVoterRepository(&dao.OptsVoterRepository{
		DB:              opts.DB,
		Client:          opts.Client,
		ContractAddress: common.HexToAddress(opts.Cfg.Blockchain.ContractAddress),
	})

	kpuBranchRepo := dao.NewKPUBranchRepository(&dao.OptsKPUBranchRepository{
		DB:              opts.DB,
		Client:          opts.Client,
		ContractAddress: common.HexToAddress(opts.Cfg.Blockchain.ContractAddress),
	})

	usersRepo := dao.NewUserRepository(&dao.OptsUserRepository{
		DB: opts.DB,
	})

	txMgr, err := txmanager.New(context.Background(), &txmanager.DriverConfig{
		Type: "sql",
		Config: txSql.Config{
			DB: opts.DB,
		},
	})
	if err != nil {
		log.Fatal("Failed to instantiate transaction manager ")
	}

	jwtSvc := jwtsvc.New(&jwtsvc.Options{
		Config: opts.Cfg.JWT,
	})

	voterUc := voter.New(&voter.Opts{
		VoterRepo: voterRepo,
		TxMgr:     txMgr,
		JwtSvc:    jwtSvc,
	})

	kpuBranchUc := kpu_branch.New(&kpu_branch.Opts{
		KpuBranchRepo: kpuBranchRepo,
		TxMgr:         txMgr,
		JwtSvc:        jwtSvc,
	})

	return &container{
		Cfg:         *opts.Cfg,
		VoterUc:     voterUc,
		KpuBranchUc: kpuBranchUc,
	}

}
