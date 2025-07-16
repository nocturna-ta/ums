package voter

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/infrastructures/wilayah"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
	"github.com/nocturna-ta/votechain-contract/binding/voterManager"
	"github.com/nocturna-ta/votechain-contract/interfaces"
)

type Module struct {
	userRepo         repository.UserRepository
	voterRepo        repository.VoterRepository
	kpuKotaRepo      repository.KPUKotaRepository
	kpuProvinsiRepo  repository.KPUProvinsiRepository
	jwtSvc           jwtsvc.JWT
	txMgr            txmanager.TxManager
	publisher        event.MessagePublisher
	topics           config.KafkaTopics
	voterContract    interfaces.VoterManagerInterface
	client           ethereum.Client
	wilayahAPIClient *wilayah.WilayahAPIClient
}

type Opts struct {
	UserRepo         repository.UserRepository
	VoterRepo        repository.VoterRepository
	KPUKotaRepo      repository.KPUKotaRepository
	KPUProvinsiRepo  repository.KPUProvinsiRepository
	TxMgr            txmanager.TxManager
	JwtSvc           jwtsvc.JWT
	Publisher        event.MessagePublisher
	Topics           config.KafkaTopics
	VoterContract    interfaces.VoterManagerInterface
	Client           ethereum.Client
	ContractAddress  common.Address
	WilayahAPIClient *wilayah.WilayahAPIClient
}

func New(opts *Opts) usecases.VoterUseCases {
	var contractInterface interfaces.VoterManagerInterface
	contract, err := voterManager.NewVoterManager(opts.ContractAddress, opts.Client.GetEthClient())
	if err != nil {
		return nil
	}
	contractInterface = contract

	wilayahClient := opts.WilayahAPIClient
	if wilayahClient == nil {
		wilayahClient = wilayah.NewWilayahAPIClient()
	}

	return &Module{
		userRepo:         opts.UserRepo,
		voterRepo:        opts.VoterRepo,
		kpuKotaRepo:      opts.KPUKotaRepo,
		kpuProvinsiRepo:  opts.KPUProvinsiRepo,
		jwtSvc:           opts.JwtSvc,
		txMgr:            opts.TxMgr,
		publisher:        opts.Publisher,
		topics:           opts.Topics,
		voterContract:    contractInterface,
		client:           opts.Client,
		wilayahAPIClient: wilayahClient,
	}
}
