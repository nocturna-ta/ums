package kpu_provinsi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
	kpuManager2 "github.com/nocturna-ta/votechain-contract/binding/kpuManager"
	"github.com/nocturna-ta/votechain-contract/interfaces"
)

type Module struct {
	kpuProvinsiRepo repository.KPUProvinsiRepository
	jwtSvc          jwtsvc.JWT
	txMgr           txmanager.TxManager
	publisher       event.MessagePublisher
	topics          config.KafkaTopics
	kpuContract     interfaces.KpuManagerInterface
	client          ethereum.Client
}

type Opts struct {
	KpuProvinsiRepo repository.KPUProvinsiRepository
	TxMgr           txmanager.TxManager
	JwtSvc          jwtsvc.JWT
	Publisher       event.MessagePublisher
	Topics          config.KafkaTopics
	KpuContract     interfaces.KpuManagerInterface
	Client          ethereum.Client
	ContractAddress common.Address
}

func New(opts *Opts) usecases.KPUProvinsiUseCases {
	var contractInterface interfaces.KpuManagerInterface
	contract, err := kpuManager2.NewKpuManager(opts.ContractAddress, opts.Client.GetEthClient())
	if err != nil {
		return nil
	}
	contractInterface = contract
	return &Module{
		kpuProvinsiRepo: opts.KpuProvinsiRepo,
		jwtSvc:          opts.JwtSvc,
		txMgr:           opts.TxMgr,
		publisher:       opts.Publisher,
		topics:          opts.Topics,
		kpuContract:     contractInterface,
		client:          opts.Client,
	}
}
