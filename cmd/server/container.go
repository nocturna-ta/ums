package server

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/txmanager"
	txSql "github.com/nocturna-ta/golib/txmanager/sql"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
	"github.com/nocturna-ta/ums/internal/usecases/auth"
	"github.com/nocturna-ta/ums/internal/usecases/kpu_kota"
	"github.com/nocturna-ta/ums/internal/usecases/kpu_provinsi"
	"github.com/nocturna-ta/ums/internal/usecases/user"
	"github.com/nocturna-ta/ums/internal/usecases/voter"
)

type container struct {
	Cfg         config.MainConfig
	VoterUc     usecases.VoterUseCases
	AuthUc      usecases.AuthUseCases
	UserUc      usecases.UserUseCases
	KpuKota     usecases.KPUKotaUseCases
	KpuProvinsi usecases.KPUProvinsiUseCases
}

type options struct {
	Cfg       *config.MainConfig
	DB        *sql.Store
	Client    ethereum.Client
	Publisher event.MessagePublisher
}

func newContainer(opts *options) *container {

	voterRepo := dao.NewVoterRepository(&dao.OptsVoterRepository{
		DB:     opts.DB,
		Client: opts.Client,
	})

	kpuKotaRepo := dao.NewKPUKotaRepository(&dao.OptsKPUKotaRepository{
		DB:     opts.DB,
		Client: opts.Client,
	})

	kpuProvinsiRepo := dao.NewKPUProvinsiRepository(&dao.OptsKPUProvinsiRepository{
		DB:     opts.DB,
		Client: opts.Client,
	})

	usersRepo := dao.NewUserRepository(&dao.OptsUserRepository{
		DB: opts.DB,
	})

	pendingRepo := dao.NewPendingRegistrationRepository(&dao.OptsPendingRegistrationRepository{
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
		VoterRepo:       voterRepo,
		TxMgr:           txMgr,
		JwtSvc:          jwtSvc,
		Publisher:       opts.Publisher,
		Topics:          opts.Cfg.Kafka.Topics,
		ContractAddress: common.HexToAddress(opts.Cfg.Blockchain.VoterManagerAddress),
		Client:          opts.Client,
	})

	kpuKotaUc := kpu_kota.New(&kpu_kota.Opts{
		KpuKotaRepo:     kpuKotaRepo,
		TxMgr:           txMgr,
		JwtSvc:          jwtSvc,
		Publisher:       opts.Publisher,
		Topics:          opts.Cfg.Kafka.Topics,
		ContractAddress: common.HexToAddress(opts.Cfg.Blockchain.KPUManagerAddress),
		Client:          opts.Client,
	})

	kpuProvinsiUc := kpu_provinsi.New(&kpu_provinsi.Opts{
		KpuProvinsiRepo: kpuProvinsiRepo,
		TxMgr:           txMgr,
		JwtSvc:          jwtSvc,
		Publisher:       opts.Publisher,
		Topics:          opts.Cfg.Kafka.Topics,
		Client:          opts.Client,
		ContractAddress: common.HexToAddress(opts.Cfg.Blockchain.KPUManagerAddress),
	})

	userUc := user.New(&user.Opts{
		UserRepo:        usersRepo,
		PendingRegRepo:  pendingRepo,
		KPUProvinsiRepo: kpuProvinsiRepo,
		KPUKotaRepo:     kpuKotaRepo,
		VoterRepo:       voterRepo,
		TxMgr:           txMgr,
		JWTSvc:          jwtSvc,
		Publisher:       opts.Publisher,
		Topics:          opts.Cfg.Kafka.Topics,
	})

	authUc := auth.New(&auth.Opts{
		JwtSvc: jwtSvc,
	})

	return &container{
		Cfg:         *opts.Cfg,
		VoterUc:     voterUc,
		AuthUc:      authUc,
		UserUc:      userUc,
		KpuKota:     kpuKotaUc,
		KpuProvinsi: kpuProvinsiUc,
	}

}
