package kafka

import (
	"context"
	"github.com/nocturna-ta/golib/event"
	_ "github.com/nocturna-ta/golib/event/kafka"
	"github.com/nocturna-ta/ums/config"
)

func NewPublisher(ctx context.Context, config config.KafkaProducerConfig) (event.MessagePublisher, error) {
	conf := &event.PublisherConfig{
		DriverConfig: &event.DriverConfig{
			Type: "kafka",
			Config: map[string]any{
				"brokers":      config.Brokers,
				"max_attempts": config.MaxAttempt,
				"idempotent":   config.Idempotent,
			},
		},
	}

	return event.NewPublisher(ctx, conf)
}
