package consumer

import (
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/utils/encryption"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	voterRepo  repository.VoterRepository
	publisher  event.MessagePublisher
	topics     config.KafkaTopics
	maxRetries int
	encryptor  *encryption.Encryption
}

type Options struct {
	VoterRepo  repository.VoterRepository
	Publisher  event.MessagePublisher
	Topics     config.KafkaTopics
	MaxRetries int
	Encryptor  *encryption.Encryption
}

func New(opts *Options) usecases.Consumer {
	maxRetries := 3
	if opts.MaxRetries > 0 {
		maxRetries = opts.MaxRetries
	}

	return &Module{
		voterRepo:  opts.VoterRepo,
		publisher:  opts.Publisher,
		topics:     opts.Topics,
		maxRetries: maxRetries,
		encryptor:  opts.Encryptor,
	}
}
