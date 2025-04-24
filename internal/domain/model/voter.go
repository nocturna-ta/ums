package model

import (
	"github.com/google/uuid"
	"github.com/nocturna-ta/common-model/models/event"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"time"
)

type Voter struct {
	BaseModel
	ID                 uuid.UUID `db:"id"`
	NIK                string    `db:"nik"`
	FullName           string    `db:"full_name"`
	BirthPlace         string    `db:"birth_place"`
	BirthDate          time.Time `db:"birth_date"`
	ResidentialAddress string    `db:"residential_address"`
	VoterAddress       string    `db:"voter_address"`
	Region             string    `db:"region"`
	IsRegistered       bool      `db:"is_registered"`
	HasVoted           bool      `db:"has_voted"`
	VotedAt            time.Time `db:"voted_at"`
	LastLogin          time.Time `db:"last_login"`
}

func (v *Voter) ToMessageModel() *event.VoterMessage {
	msg := &event.VoterMessage{
		BaseModelMessage: event.BaseModelMessage{
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			IsDeleted: v.IsDeleted,
		},
		ID:           v.ID.String(),
		NIK:          v.NIK,
		VoterAddress: v.VoterAddress,
		Region:       v.Region,
		IsRegistered: v.IsRegistered,
	}
	return msg
}

func ConstructRegistration(req *request.VoterRegistrationRequest) *Voter {
	now := time.Now()

	voterId := uuid.New()
	layout := "2006-01-02"
	birthDate, err := time.Parse(layout, req.BirthDate)
	if err != nil {
		return nil
	}

	user := &Voter{
		BaseModel: BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
			IsDeleted: false,
		},
		ID:                 voterId,
		NIK:                req.NIK,
		FullName:           req.FullName,
		BirthPlace:         req.BirthPlace,
		BirthDate:          birthDate,
		ResidentialAddress: req.ResidentialAddress,
		VoterAddress:       req.VoterAddress,
		Region:             req.Region,
		HasVoted:           false,
		VotedAt:            time.Time{},
		LastLogin:          now,
	}
	return user
}
