package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/ums/internal/domain/model"
)

type PendingRegistrationRepository interface {
	Insert(ctx context.Context, registration *model.PendingRegistration) error
	GetByUserID(ctx context.Context, userID uuid.UUID) (*model.PendingRegistration, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
