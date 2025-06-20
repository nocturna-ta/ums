package consumer

import (
	"context"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/infrastructures/ethreum"
	"github.com/nocturna-ta/ums/internal/infrastructures/kafka"
	"github.com/spf13/cobra"
)

var (
	serveConsumerCmd = &cobra.Command{
		Use:   "run-consumer",
		Short: "Vote Consumer Service",
		RunE:  run,
	}
)

func ServeConsumerCmd() *cobra.Command {
	serveConsumerCmd.Flags().StringP("config", "c", "", "Config Path, both relative or absolute. i.e: /usr/local/bin/config/files")
	return serveConsumerCmd
}

func run(cmd *cobra.Command, args []string) error {
	configLocation, _ := cmd.Flags().GetString("config")
	cfg := &config.MainConfig{}
	config.ReadConfig(cfg, configLocation)

	publisher, err := kafka.NewPublisher(context.Background(), cfg.Kafka.Producer)
	if err != nil {
		log.Fatalf("Failed to instantiate kafka publisher: %w", err)
	}

	database := sql.New(sql.DBConfig{
		SlaveDSN:        cfg.Database.SlaveDSN,
		MasterDSN:       cfg.Database.MasterDSN,
		RetryInterval:   cfg.Database.RetryInterval,
		MaxIdleConn:     cfg.Database.MaxIdleConn,
		MaxConn:         cfg.Database.MaxConn,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	}, sql.DriverPostgres)

	databaseClickhouse := sql.New(sql.DBConfig{
		SlaveDSN:        cfg.DatabaseClickhouse.SlaveDSN,
		MasterDSN:       cfg.DatabaseClickhouse.MasterDSN,
		RetryInterval:   cfg.DatabaseClickhouse.RetryInterval,
		MaxIdleConn:     cfg.DatabaseClickhouse.MaxIdleConn,
		MaxConn:         cfg.DatabaseClickhouse.MaxConn,
		ConnMaxLifetime: cfg.DatabaseClickhouse.ConnMaxLifetime,
	}, sql.DriverClickHouse)

	client, err := ethreum.GetEthereumClient(&cfg.Blockchain)
	if err != nil {
		return err
	}

	defer client.Close()

	appContainer := newContainer(&options{
		Cfg:          cfg,
		Publisher:    publisher,
		DB:           database,
		DBClickhouse: databaseClickhouse,
		Client:       client,
	})

	consumer, err := kafka.NewConsumer(context.Background(), cfg.Kafka.Consumer, &appContainer.EventHandler)
	if err != nil {
		log.Fatalf("Failed to instantiate kafka consumer: %w", err)
	}

	topicHandler := map[event.TopicName]event.ConsumerHandlerConfig{
		event.TopicName(cfg.Kafka.Topics.VoteProcessed.Value): {
			Handler:           appContainer.ConsumerUc.UpdateVoterVoteStatus,
			ConsumerGroup:     cfg.Kafka.Consumer.ConsumerGroup,
			ErrorHandlerLevel: cfg.Kafka.Topics.VoteProcessed.ErrorHandler,
			WithBackOff:       cfg.Kafka.Topics.VoteProcessed.WithBackOff,
		},
		event.TopicName(cfg.Kafka.Topics.UserLogs.Value): {
			ConsumerGroup:     cfg.Kafka.Consumer.ConsumerGroup,
			ErrorHandlerLevel: cfg.Kafka.Topics.UserLogs.ErrorHandler,
			Handler:           appContainer.ConsumerUc.InsertUserLog,
			WithBackOff:       cfg.Kafka.Topics.UserLogs.WithBackOff,
		},
	}

	consumer.RunWithHandlerConfig(topicHandler)

	return nil
}
