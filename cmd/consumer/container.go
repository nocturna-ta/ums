package consumer

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/event/handler"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/utils/encryption"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/usecases"
	"github.com/nocturna-ta/ums/internal/usecases/consumer"
)

type container struct {
	Cfg          config.MainConfig
	ConsumerUc   usecases.Consumer
	EventHandler handler.EventHandler
}

type options struct {
	Cfg       *config.MainConfig
	Publisher event.MessagePublisher
	DB        *sql.Store
	Client    ethereum.Client
}

func newContainer(opts *options) *container {
	encryptor, err := encryption.NewEncryption(opts.Cfg.Encryption.Key)
	if err != nil {
		log.Fatal("failed to create encryption: %v", err)
	}

	voterRepo := dao.NewVoterRepository(&dao.OptsVoterRepository{
		Client:          opts.Client,
		ContractAddress: common.HexToAddress(opts.Cfg.Blockchain.VoterManagerAddress),
		DB:              opts.DB,
	})

	consumerUc := consumer.New(&consumer.Options{
		VoterRepo:  voterRepo,
		Publisher:  opts.Publisher,
		Topics:     opts.Cfg.Kafka.Topics,
		MaxRetries: opts.Cfg.Kafka.Consumer.MaxRetries,
		Encryptor:  encryptor,
	})

	eventHandler := handler.New(&handler.Options{
		RetryConfig: handler.RetryConfig{
			MaxRetry:          opts.Cfg.Kafka.Consumer.Retry.MaxRetry,
			RetryInitialDelay: opts.Cfg.Kafka.Consumer.Retry.RetryInitialDelay,
			MaxJitter:         opts.Cfg.Kafka.Consumer.Retry.MaxJitter,
			HandlerTimeout:    opts.Cfg.Kafka.Consumer.Retry.HandlerTimeout,
			BackOffConfig:     opts.Cfg.Kafka.Consumer.Retry.BackOffConfig,
		},
		Publisher:   opts.Publisher,
		DlqTopic:    opts.Cfg.Kafka.Topics.VoteDLQ.Value,
		ServiceName: "ums-service",
	})

	return &container{
		Cfg:          *opts.Cfg,
		ConsumerUc:   consumerUc,
		EventHandler: eventHandler,
	}
}
