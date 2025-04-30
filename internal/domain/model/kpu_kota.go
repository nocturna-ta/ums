package model

import (
	"github.com/google/uuid"
	"github.com/nocturna-ta/common-model/models/event"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"time"
)

type KPUKota struct {
	BaseModel
	ID           uuid.UUID `db:"id"`
	UserID       uuid.UUID `db:"user_id"`
	Name         string    `db:"name"`
	Address      string    `db:"address"`
	Region       string    `db:"region"`
	IsActive     bool      `db:"is_active"`
	PhotoPath    string    `db:"photo_path"`
	Telephone    string    `db:"telephone"`
	RegisteredAt time.Time `db:"registered_at"`
}

func (u *KPUKota) ToMessageModel() *event.KPUKotaMessage {
	msg := &event.KPUKotaMessage{
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

func ConstructRegistrationKPUKota(req *request.KPUKotaRegistrationRequest) *KPUKota {
	now := time.Now()
	kpu := &KPUKota{
		BaseModel: BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
			IsDeleted: false,
		},
		ID:           uuid.New(),
		Name:         req.Name,
		Address:      req.Address,
		Region:       req.Region,
		IsActive:     req.IsActive,
		PhotoPath:    "",
		Telephone:    "",
		RegisteredAt: time.Now(),
	}
	return kpu
}
