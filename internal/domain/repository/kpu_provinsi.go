package repository

import (
	"context"
	"github.com/nocturna-ta/ums/internal/domain/model"
)

type KPUProvinsiRepository interface {
	InsertKPUProvinsi(ctx context.Context, kpu *model.KPUProvinsi, signedTransaction string) error
	GetAllKPUProvinsi(ctx context.Context) ([]model.KPUProvinsi, error)
	GetKPUProvinsiByAddress(ctx context.Context, address string) (*model.KPUProvinsi, error)
}
