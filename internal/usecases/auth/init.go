package auth

import (
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	jwtSvc jwtsvc.JWT
}

type Opts struct {
	JwtSvc jwtsvc.JWT
}

func New(opts *Opts) usecases.AuthUseCases {
	return &Module{
		jwtSvc: opts.JwtSvc,
	}
}
