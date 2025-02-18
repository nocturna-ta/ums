package model

import (
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/utils/randomizer"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/pkg/utils"
	"math/rand"
	"time"
)

type Voter struct {
	BaseModel
	ID           uuid.UUID `db:"id"`
	NIK          string    `db:"nik"`
	VoterAddress string    `db:"voter_address"`
	Region       string    `db:"region"`
	IsRegistered bool      `db:"is_registered"`
	HasVoted     bool      `db:"has_voted"`
	Password     string    `db:"password"`
	PasswordSalt string    `db:"password_salt"`
	VotedAt      time.Time `db:"voted_at"`
	LastLogin    time.Time `db:"last_login"`
}

func ConstructRegistration(req *request.VoterRegistrationRequest) *Voter {
	now := time.Now()

	n := rand.Intn(10)
	if n < 6 {
		n += 4
	}

	salt := randomizer.RandomString(n)
	voterId := uuid.New()

	user := &Voter{
		BaseModel: BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
			IsDeleted: false,
		},
		ID:           voterId,
		NIK:          req.NIK,
		VoterAddress: req.VoterAddress,
		Password:     utils.PasswordHash(req.NIK, salt),
		PasswordSalt: salt,
		IsRegistered: req.IsRegistered,
		HasVoted:     false,
		VotedAt:      time.Time{},
		Region:       req.Region,
		LastLogin:    now,
	}
	return user
}
