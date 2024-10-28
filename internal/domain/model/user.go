package model

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/utils/randomizer"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/pkg/sqlutils"
	"github.com/nocturna-ta/ums/pkg/utils"
	"math/rand"
	"time"
)

type User struct {
	BaseModel
	ID           uuid.UUID      `db:"id"`
	NIK          string         `db:"nik"`
	NoTelephone  sql.NullString `db:"no_telephone"`
	Email        sql.NullString `db:"email"`
	Name         string         `db:"name"`
	Password     string         `db:"password"`
	PasswordSalt string         `db:"password_salt"`
}

func ConstructRegistration(req *request.UserRegisterRequest) *User {
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
			IsDeleted: false,
		},
		ID:           userId,
		NIK:          req.NIK,
		NoTelephone:  sqlutils.NewNullString(nil),
		Email:        sqlutils.NewNullString(nil),
		Name:         req.Name,
		Password:     utils.PasswordHash(req.NIK, salt),
		PasswordSalt: salt,
	}
	return user
}
