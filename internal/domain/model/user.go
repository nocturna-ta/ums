package model

import (
	"github.com/google/uuid"
	"github.com/nocturna-ta/common-model/models/event"
	"github.com/nocturna-ta/golib/utils/randomizer"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/pkg/constants"
	"github.com/nocturna-ta/ums/pkg/utils"
	"math/rand"
	"strings"
	"time"
)

type User struct {
	BaseModel
	ID           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	PasswordSalt string    `db:"password_salt"`
	Role         string    `db:"role"`
	IsActive     bool      `db:"is_active"`
}

type UserUpdate struct {
	Username string `db:"username"`
}

func (u *User) ToMessageModel() *event.UserMessage {
	msg := &event.UserMessage{
		BaseModelMessage: event.BaseModelMessage{
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			IsDeleted: u.IsDeleted,
		},
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
		IsActive: u.IsActive,
	}

	return msg
}

func ConstructUserRegistration(req *request.UserRegistrationRequest) *User {
	now := time.Now()
	username := req.Username
	if username == constants.EmptyString {
		username = "user-" + now.Format("20060102150405")
	}
	username = strings.TrimSpace(username)

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
		ID:           userId,
		Username:     username,
		Email:        req.Email,
		Password:     utils.PasswordHash(req.Password, salt),
		PasswordSalt: salt,
		Role:         req.Role,
		IsActive:     true,
	}

	return user
}
