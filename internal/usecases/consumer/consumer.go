package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	event2 "github.com/nocturna-ta/common-model/models/event"
	"github.com/nocturna-ta/common-model/models/logging"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/pkg/constants"
	"github.com/nocturna-ta/ums/pkg/roles"
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
	voter.VotedAt = &voteProcessedMessage.ProcessedAt

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

func (m *Module) InsertUserLog(ctx context.Context, message *event.EventConsumeMessage) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "Consumer.InsertUserLog")
	defer span.End()

	requestId := libCtx.ReadRequestId(ctx)
	log.WithFields(log.Fields{
		"request_id": requestId,
		"topic":      message.Topic,
		"key":        message.Key,
	}).InfoWithCtx(ctx, "[Consumer.InsertUserLog] consume message")

	var userLogMessage logging.KPULogs
	err := json.Unmarshal(message.Data, &userLogMessage)
	if err != nil {
		log.WithFields(log.Fields{
			"error":      err,
			"topic":      message.Topic,
			"data":       string(message.Data),
			"request_id": requestId,
		}).ErrorWithCtx(ctx, "[Consumer.InsertUserLog] failed to unmarshal message")
		return err
	}

	operation, ok := message.Metadata[constants.MetaDataOperation].(string)
	if !ok {
		log.WithFields(log.Fields{
			"request_id": requestId,
			"error":      "missing operation metadata",
			"topic":      message.Topic,
			"metadata":   message.Metadata,
		}).ErrorWithCtx(ctx, "[Consumer.InsertUserLog] missing operation metadata")
		operation = constants.Create
	}

	userId, err := uuid.Parse(userLogMessage.UserID)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).ErrorWithCtx(ctx, "[Consumer.InsertUserLog] failed to parse user id")
		return err
	}

	user, err := m.userRepo.GetById(ctx, userId)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"user_id": userId,
		}).ErrorWithCtx(ctx, "[Consumer.InsertUserLog] failed to get user by id")
		return err
	}

	var (
		username string
		name     string
	)

	if user.Role == roles.RoleKPUProvinsi {
		kpu, err := m.kpuProvinsiRepo.GetKPUProvinsiByUserID(ctx, user.ID)
		if err != nil {
			log.WithFields(log.Fields{
				"error":   err,
				"user_id": userId,
			}).ErrorWithCtx(ctx, "[Consumer.InsertUserLog] failed to get KPU Provinsi by user id")
			return err
		}

		username = kpu.Username
		name = kpu.Name
	} else if user.Role == roles.RoleKPUKota {
		kpu, err := m.kpuKotaRepo.GetKPUKotaByUserID(ctx, user.ID)
		if err != nil {
			log.WithFields(log.Fields{
				"error":   err,
				"user_id": userId,
			}).ErrorWithCtx(ctx, "[Consumer.InsertUserLog] failed to get KPU Kota by user id")
			return err
		}

		username = kpu.Username
		name = kpu.Name
	} else if user.Role == roles.RoleKPUPusat {
		username = "KPU Pusat"
		name = "KPU Pusat"
	} else {
		log.WithFields(log.Fields{
			"error":   "unsupported user role",
			"user_id": userId,
			"role":    user.Role,
		}).ErrorWithCtx(ctx, "[Consumer.InsertUserLog] unsupported user role")
		return errors.New("unsupported user role")
	}

	logEntry := &model.UserLogs{
		ID:           uuid.New().String(),
		UserID:       user.ID.String(),
		Username:     username,
		Name:         name,
		Role:         user.Role,
		Time:         userLogMessage.CreatedAt,
		Activity:     userLogMessage.Activity,
		ActivityType: operation,
	}

	err = m.userLogRepo.InsertLog(ctx, logEntry)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"user_id": userId,
			"log":     logEntry,
		}).ErrorWithCtx(ctx, "[Consumer.InsertUserLog] failed to insert user log")
		return err
	}

	log.WithFields(log.Fields{
		"user_id":    userId,
		"activity":   userLogMessage.Activity,
		"operation":  operation,
		"request_id": requestId,
	}).InfoWithCtx(ctx, "[Consumer.InsertUserLog] user log inserted successfully")

	return nil
}
