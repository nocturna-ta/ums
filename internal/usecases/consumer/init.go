package consumer

import (
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/utils/encryption"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	voterRepo       repository.VoterRepository
	kpuProvinsiRepo repository.KPUProvinsiRepository
	kpuKotaRepo     repository.KPUKotaRepository
	userRepo        repository.UserRepository
	userLogRepo     repository.UserLogRepository
	publisher       event.MessagePublisher
	topics          config.KafkaTopics
	maxRetries      int
	encryptor       *encryption.Encryption
}

type Options struct {
	VoterRepo       repository.VoterRepository
	KPUProvinsiRepo repository.KPUProvinsiRepository
	KPUKotaRepo     repository.KPUKotaRepository
	UserRepo        repository.UserRepository
	UserLogRepo     repository.UserLogRepository
	Publisher       event.MessagePublisher
	Topics          config.KafkaTopics
	MaxRetries      int
	Encryptor       *encryption.Encryption
}

func New(opts *Options) usecases.Consumer {
	maxRetries := 3
	if opts.MaxRetries > 0 {
		maxRetries = opts.MaxRetries
	}

	return &Module{
		voterRepo:       opts.VoterRepo,
		kpuProvinsiRepo: opts.KPUProvinsiRepo,
		kpuKotaRepo:     opts.KPUKotaRepo,
		userRepo:        opts.UserRepo,
		userLogRepo:     opts.UserLogRepo,
		publisher:       opts.Publisher,
		topics:          opts.Topics,
		maxRetries:      maxRetries,
		encryptor:       opts.Encryptor,
	}
}
