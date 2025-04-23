package user

import (
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	userRepo        repository.UserRepository
	pendingRegRepo  repository.PendingRegistrationRepository
	kpuProvinsiRepo repository.KPUProvinsiRepository
	kpuKotaRepo     repository.KPUKotaRepository
	voterRepo       repository.VoterRepository

	txMgr     txmanager.TxManager
	jwtSvc    jwtsvc.JWT
	publisher event.MessagePublisher
	topics    config.KafkaTopics
}

type Opts struct {
	UserRepo        repository.UserRepository
	PendingRegRepo  repository.PendingRegistrationRepository
	KPUProvinsiRepo repository.KPUProvinsiRepository
	KPUKotaRepo     repository.KPUKotaRepository
	VoterRepo       repository.VoterRepository
	TxMgr           txmanager.TxManager
	JWTSvc          jwtsvc.JWT
	Publisher       event.MessagePublisher
	Topics          config.KafkaTopics
}

func New(opts *Opts) usecases.UserUseCases {
	return &Module{
		userRepo:        opts.UserRepo,
		pendingRegRepo:  opts.PendingRegRepo,
		kpuProvinsiRepo: opts.KPUProvinsiRepo,
		kpuKotaRepo:     opts.KPUKotaRepo,
		voterRepo:       opts.VoterRepo,
		txMgr:           opts.TxMgr,
		jwtSvc:          opts.JWTSvc,
		publisher:       opts.Publisher,
		topics:          opts.Topics,
	}
}
