package kpu_provinsi

import (
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	kpuProvinsiRepo repository.KPUProvinsiRepository
	jwtSvc          jwtsvc.JWT
	txMgr           txmanager.TxManager
	publisher       event.MessagePublisher
	topics          config.KafkaTopics
}

type Opts struct {
	KpuProvinsiRepo repository.KPUProvinsiRepository
	TxMgr           txmanager.TxManager
	JwtSvc          jwtsvc.JWT
	Publisher       event.MessagePublisher
	Topics          config.KafkaTopics
}

func New(opts *Opts) usecases.KPUProvinsiUseCases {
	return &Module{
		kpuProvinsiRepo: opts.KpuProvinsiRepo,
		jwtSvc:          opts.JwtSvc,
		txMgr:           opts.TxMgr,
		publisher:       opts.Publisher,
		topics:          opts.Topics,
	}
}
