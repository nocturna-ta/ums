package model

import (
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/utils/randomizer"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/pkg/utils"
	"math/rand"
	"time"
)

type KPUBranch struct {
	BaseModel
	ID            uuid.UUID `db:"id"`
	Name          string    `db:"name"`
	BranchAddress string    `db:"branch_address"`
	Region        string    `db:"region"`
	IsActive      bool      `db:"is_active"`
	Password      string    `db:"password"`
	PasswordSalt  string    `db:"password_salt"`
}

func ConstructRegistrationKPUBranch(req *request.KPUBranchRegistrationRequest) *KPUBranch {

	now := time.Now()

	n := rand.Intn(10)
	if n < 6 {
		n += 4
	}

	salt := randomizer.RandomString(n)

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
		Password:      utils.PasswordHash(req.BranchAddress, salt),
		PasswordSalt:  salt,
	}
	return user
}
