package consumer

import (
	"context"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/utils/encryption"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"testing"
)

func TestModule_UpdateVoterVoteStatus(t *testing.T) {
	type fields struct {
		voterRepo  repository.VoterRepository
		publisher  event.MessagePublisher
		topics     config.KafkaTopics
		maxRetries int
		encryptor  *encryption.Encryption
	}
	type args struct {
		ctx     context.Context
		message *event.EventConsumeMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				voterRepo:  tt.fields.voterRepo,
				publisher:  tt.fields.publisher,
				topics:     tt.fields.topics,
				maxRetries: tt.fields.maxRetries,
				encryptor:  tt.fields.encryptor,
			}
			if err := m.UpdateVoterVoteStatus(tt.args.ctx, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("UpdateVoterVoteStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
