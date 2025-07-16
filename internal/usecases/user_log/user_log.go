package user_log

import (
	"context"
	"github.com/nocturna-ta/golib/custerr"
	log2 "github.com/nocturna-ta/golib/log"
	response2 "github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

func (m *Module) GetAllUserLog(ctx context.Context, limit, offset int) (*[]response.UserLogResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserLogUseCases.GetAllUserLog")
	defer span.End()

	if limit <= 0 {
		limit = 10
	}
	if limit > 1000 {
		limit = 1000
	}
	if offset < 0 {
		offset = 0
	}

	logs, err := m.userLogRepo.GetUserLogs(ctx, limit, offset)
	if err != nil {
		log2.WithFields(log2.Fields{
			"error":  err,
			"limit":  limit,
			"offset": offset,
			"log":    logs,
		}).ErrorWithCtx(ctx, "[UserLogUseCases.GetAllUserLog] Failed to get user logs")
		return nil, err
	}

	var userLogResponses []response.UserLogResponse
	for _, log := range logs {
		userLogResponses = append(userLogResponses, response.UserLogResponse{
			ID:           log.ID,
			UserID:       log.UserID,
			Username:     log.Username,
			Name:         log.Name,
			Role:         log.Role,
			Time:         log.Time,
			Activity:     log.Activity,
			ActivityType: log.ActivityType,
		})
	}

	if len(userLogResponses) == 0 {
		return nil, custerr.ErrChain{
			Message: "No user logs found",
			Code:    404,
			Type:    response2.ErrNotFound,
		}
	}

	return &userLogResponses, nil
}
