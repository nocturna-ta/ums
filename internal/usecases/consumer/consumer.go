package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	event2 "github.com/nocturna-ta/common-model/models/event"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
)

func (m *Module) UpdateVoterVoteStatus(ctx context.Context, message *event.EventConsumeMessage) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "Consumer.UpdateVoterVoteStatus")
	defer span.End()

	requestId := libCtx.ReadRequestId(ctx)
	log.WithFields(log.Fields{
		"request_id": requestId,
		"topic":      message.Topic,
		"key":        message.Key,
	}).InfoWithCtx(ctx, "[Consumer.UpdateVoterVoteStatus] consume message")

	var voteProcessedMessage event2.VoteProcessedMessage
	err := json.Unmarshal(message.Data, &voteProcessedMessage)
	if err != nil {
		log.WithFields(log.Fields{
			"error":      err,
			"topic":      message.Topic,
			"data":       string(message.Data),
			"request_id": requestId,
		}).ErrorWithCtx(ctx, "[Consumer.UpdateVoterVoteStatus] failed to unmarshal message")
		return nil
	}

	voterId, err := m.encryptor.Decrypt(voteProcessedMessage.VoterID)
	if err != nil {
		log.WithFields(log.Fields{
			"error":      err,
			"topic":      message.Topic,
			"data":       string(message.Data),
			"request_id": requestId,
		}).ErrorWithCtx(ctx, "[Consumer.UpdateVoterVoteStatus] failed to decrypt voter id")
		return nil
	}

	voterIdUUID, err := uuid.Parse(voterId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":      err,
			"topic":      message.Topic,
			"data":       string(message.Data),
			"request_id": requestId,
		}).ErrorWithCtx(ctx, "[Consumer.UpdateVoterVoteStatus] failed to parse voter id")
		return nil
	}

	voter, err := m.voterRepo.GetVoterByID(ctx, voterIdUUID)
	if err != nil {
		if errors.Is(err, dao.ErrNoResult) {
			log.WithFields(log.Fields{
				"error":    err,
				"voter_id": voterId,
			}).ErrorWithCtx(ctx, "[Consumer.UpdateVoterVoteStatus] voter not found")
			return nil
		}

		log.WithFields(log.Fields{
			"error":    err,
			"voter_id": voterId,
		}).ErrorWithCtx(ctx, "[Consumer.UpdateVoterVoteStatus] failed to get voter by id")
		return err
	}

	voter.HasVoted = true
	voter.VotedAt = voteProcessedMessage.ProcessedAt

	err = m.voterRepo.UpdateVoter(ctx, voter)
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"voter_id": voterId,
		}).ErrorWithCtx(ctx, "[Consumer.UpdateVoterVoteStatus] failed to update voter")
		return err
	}

	log.WithFields(log.Fields{
		"voter_id":   voterId,
		"request_id": requestId,
	}).InfoWithCtx(ctx, "[Consumer.UpdateVoterVoteStatus] voter updated successfully")
	return nil
}
