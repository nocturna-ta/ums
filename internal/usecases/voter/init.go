package voter

import (
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Module struct {
	voterRepo repository.VoterRepository
	jwtSvc    jwtsvc.JWT
	txMgr     txmanager.TxManager
}

type Opts struct {
	VoterRepo repository.VoterRepository
	TxMgr     txmanager.TxManager
	JwtSvc    jwtsvc.JWT
}

func New(opts *Opts) usecases.VoterUseCases {
	return &Module{
		voterRepo: opts.VoterRepo,
		jwtSvc:    opts.JwtSvc,
		txMgr:     opts.TxMgr,
	}
}
