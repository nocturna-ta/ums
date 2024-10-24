package model

import (
	"database/sql"
	"github.com/google/uuid"
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
