package user_log

import (
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	userLogRepo repository.UserLogRepository
}

type Opts struct {
	UserLogRepo repository.UserLogRepository
}

func New(opts *Opts) usecases.UserLogUseCases {
	return &Module{
		userLogRepo: opts.UserLogRepo,
	}
}
