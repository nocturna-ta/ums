package usecases

import (
	"context"
	"github.com/nocturna-ta/golib/event"
)

type Consumer interface {
	UpdateVoterVoteStatus(ctx context.Context, message *event.EventConsumeMessage) error
	InsertUserLog(ctx context.Context, message *event.EventConsumeMessage) error
}
