package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/ums/internal/domain/model"
)

type UserRepository interface {
	Insert(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByNIK(ctx context.Context, nik string) (*model.User, error)
	ChangePassword(ctx context.Context, id uuid.UUID, newPass string) error
}
