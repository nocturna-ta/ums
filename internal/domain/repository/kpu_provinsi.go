package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/ums/internal/domain/model"
)

type KPUProvinsiRepository interface {
	InsertKPUProvinsi(ctx context.Context, kpu *model.KPUProvinsi, signedTransaction string) error
	GetAllKPUProvinsi(ctx context.Context) ([]model.KPUProvinsi, error)
	GetKPUProvinsiByAddress(ctx context.Context, address string) (*model.KPUProvinsi, error)
	UpdateKPUProvinsiPhoto(ctx context.Context, id uuid.UUID, photoPath string) error
	GetKPUProvinsiByID(ctx context.Context, id uuid.UUID) (*model.KPUProvinsi, error)
}
