package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/ums/internal/domain/model"
)

type KPUKotaRepository interface {
	InsertKPUKota(ctx context.Context, kpu *model.KPUKota, signedTransaction string) error
	GetAllKPUKota(ctx context.Context) ([]model.KPUKota, error)
	GetKPUKotaByAddress(ctx context.Context, address string) (*model.KPUKota, error)
	UpdateKPUKotaPhoto(ctx context.Context, id uuid.UUID, photoPath string) error
	GetKPUKotaByID(ctx context.Context, id uuid.UUID) (*model.KPUKota, error)
}
