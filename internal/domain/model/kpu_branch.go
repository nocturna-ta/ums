package model

import (
	"github.com/google/uuid"
	"github.com/nocturna-ta/common-model/models/event"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"time"
)

type KPUBranch struct {
	BaseModel
	ID            uuid.UUID `db:"id"`
	Name          string    `db:"name"`
	BranchAddress string    `db:"branch_address"`
	Region        string    `db:"region"`
	IsActive      bool      `db:"is_active"`
}

func (u *KPUBranch) ToMessageModel() *event.KPUBranchMessage {
	msg := &event.KPUBranchMessage{
		BaseModelMessage: event.BaseModelMessage{
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			IsDeleted: u.IsDeleted,
		},
		ID:            u.ID.String(),
		Name:          u.Name,
		BranchAddress: u.BranchAddress,
		Region:        u.Region,
	}
	return msg
}

func ConstructRegistrationKPUBranch(req *request.KPUBranchRegistrationRequest) *KPUBranch {
	now := time.Now()
	user := &KPUBranch{
		BaseModel: BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
			IsDeleted: false,
		},
		ID:            uuid.New(),
		Name:          req.Name,
		BranchAddress: req.BranchAddress,
		Region:        req.Region,
		IsActive:      req.IsActive,
	}
	return user
}
