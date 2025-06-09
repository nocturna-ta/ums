package user

import (
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/infrastructures/wilayah"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
	"time"
)

type Module struct {
	userRepo            repository.UserRepository
	pendingRegRepo      repository.PendingRegistrationRepository
	kpuProvinsiRepo     repository.KPUProvinsiRepository
	kpuKotaRepo         repository.KPUKotaRepository
	voterRepo           repository.VoterRepository
	kpuProvinsiUseCases usecases.KPUProvinsiUseCases
	kpuKotaUseCases     usecases.KPUKotaUseCases
	voterUseCases       usecases.VoterUseCases
	txMgr               txmanager.TxManager
	jwtSvc              jwtsvc.JWT
	publisher           event.MessagePublisher
	topics              config.KafkaTopics
	wilayahAPIClient    *wilayah.WilayahAPIClient
	regionCache         *wilayah.RegionCache
}

type Opts struct {
	UserRepo            repository.UserRepository
	PendingRegRepo      repository.PendingRegistrationRepository
	KPUProvinsiRepo     repository.KPUProvinsiRepository
	KPUKotaRepo         repository.KPUKotaRepository
	VoterRepo           repository.VoterRepository
	KPUProvinsiUseCases usecases.KPUProvinsiUseCases
	KPUKotaUseCases     usecases.KPUKotaUseCases
	VoterUseCases       usecases.VoterUseCases
	TxMgr               txmanager.TxManager
	JWTSvc              jwtsvc.JWT
	Publisher           event.MessagePublisher
	Topics              config.KafkaTopics
	WilayahAPIClient    *wilayah.WilayahAPIClient
	RegionCache         *wilayah.RegionCache
}

func New(opts *Opts) usecases.UserUseCases {
	var wilayahClient *wilayah.WilayahAPIClient
	var regionCache *wilayah.RegionCache

	wilayahClient = wilayah.NewWilayahAPIClient()
	regionCache = wilayah.NewRegionCache(24 * time.Hour)

	return &Module{
		userRepo:            opts.UserRepo,
		pendingRegRepo:      opts.PendingRegRepo,
		kpuProvinsiRepo:     opts.KPUProvinsiRepo,
		kpuKotaRepo:         opts.KPUKotaRepo,
		voterRepo:           opts.VoterRepo,
		kpuProvinsiUseCases: opts.KPUProvinsiUseCases,
		kpuKotaUseCases:     opts.KPUKotaUseCases,
		voterUseCases:       opts.VoterUseCases,
		txMgr:               opts.TxMgr,
		jwtSvc:              opts.JWTSvc,
		publisher:           opts.Publisher,
		topics:              opts.Topics,
		wilayahAPIClient:    wilayahClient,
		regionCache:         regionCache,
	}
}
