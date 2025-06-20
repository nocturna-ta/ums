package usecases

import (
	"context"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type UserLogUseCases interface {
	GetAllUserLog(ctx context.Context, limit, offset int) ([]response.UserLogResponse, error)
}
