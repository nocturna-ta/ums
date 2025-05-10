package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/ums/internal/domain/model"
)

type UserRepository interface {
	Insert(ctx context.Context, user *model.User) error
	GetById(ctx context.Context, id uuid.UUID) (*model.User, error)
	ChangePassword(ctx context.Context, id uuid.UUID, newPass string) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateVerificationStatus(ctx context.Context, id uuid.UUID, status string, role string) error
	GetPendingVerificationUsers(ctx context.Context) ([]model.User, error)
	GetPendingVerificationUsersByRequestedRole(ctx context.Context, requestedRole string) ([]model.User, error)
}
