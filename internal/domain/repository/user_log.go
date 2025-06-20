package repository

import (
	"context"
	"github.com/nocturna-ta/ums/internal/domain/model"
)

type UserLogRepository interface {
	InsertLog(ctx context.Context, log *model.UserLogs) error
	GetUserLogs(ctx context.Context, limit, offset int) ([]model.UserLogs, error)
}
