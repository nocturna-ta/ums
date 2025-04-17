package model

import (
	"github.com/google/uuid"
	"github.com/nocturna-ta/common-model/models/event"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"time"
)

type KPUProvinsi struct {
	BaseModel
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Address  string    `db:"address"`
	Region   string    `db:"region"`
	IsActive bool      `db:"is_active"`
	UserID   uuid.UUID `db:"user_id"`
}

func (u *KPUProvinsi) ToMessageModel() *event.KPUProvinsiMessage {
	msg := &event.KPUProvinsiMessage{
		BaseModelMessage: event.BaseModelMessage{
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			IsDeleted: u.IsDeleted,
		},
		ID:      u.ID.String(),
		Name:    u.Name,
		Address: u.Address,
		Region:  u.Region,
	}
	return msg
}

func ConstructRegistrationKPUProvinsi(req *request.KPUProvinsiRegistrationRequest) *KPUProvinsi {
	now := time.Now()
	kpu := &KPUProvinsi{
		BaseModel: BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
			IsDeleted: false,
		},
		ID:       uuid.New(),
		Name:     req.Name,
		Address:  req.Address,
		Region:   req.Region,
		IsActive: req.IsActive,
	}
	return kpu
}
