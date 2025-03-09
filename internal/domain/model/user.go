package model

import "github.com/google/uuid"

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
