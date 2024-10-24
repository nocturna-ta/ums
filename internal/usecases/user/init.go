package user

import (
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	userRepo repository.UserRepository
}

type Opts struct {
	UserRepo repository.UserRepository
}

func New(opts *Opts) usecases.UserUseCases {
	return &Module{
		userRepo: opts.UserRepo,
	}
}
