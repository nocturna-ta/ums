package model

import (
	"github.com/google/uuid"
	"github.com/nocturna-ta/common-model/models/event"
	"github.com/nocturna-ta/golib/utils/randomizer"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/pkg/common"
	"math/rand"
	"time"
)

type User struct {
	BaseModel
	ID                 uuid.UUID `db:"id"`
	Email              string    `db:"email"`
	Password           string    `db:"password"`
	PasswordSalt       string    `db:"password_salt"`
	Role               string    `db:"role"`
	RequestedRole      string    `db:"requested_role"`
	IsActive           bool      `db:"is_active"`
	VerificationStatus string    `db:"verification_status"`
}

const (
	VerificationStatusPending  = "pending"
	VerificationStatusApproved = "approved"
	VerificationStatusRejected = "rejected"
)

func (u *User) ToMessageModel() *event.UserMessage {
	msg := &event.UserMessage{
		BaseModelMessage: event.BaseModelMessage{
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			IsDeleted: u.IsDeleted,
		},
		ID:       u.ID,
		Email:    u.Email,
		Role:     u.Role,
		IsActive: u.IsActive,
	}

	return msg
}

func ConstructUserRegistration(req *request.UserRegistrationRequest) *User {
	now := time.Now()

	n := rand.Intn(10)
	if n < 6 {
		n += 4
	}

	salt := randomizer.RandomString(n)

	userId := uuid.New()

	user := &User{
		BaseModel: BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
		},
		ID:                 userId,
		Email:              req.Email,
		Password:           common.PasswordHash(req.Password, salt),
		PasswordSalt:       salt,
		Role:               "unverified",
		RequestedRole:      req.Role,
		IsActive:           false,
		VerificationStatus: VerificationStatusPending,
	}

	return user
}
