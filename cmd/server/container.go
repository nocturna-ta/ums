package server

import (
	"context"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/txmanager"
	txSql "github.com/nocturna-ta/golib/txmanager/sql"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
	"github.com/nocturna-ta/ums/internal/usecases/user"
)

type container struct {
	Cfg    config.MainConfig
	UserUC usecases.UserUseCases
}

type options struct {
	Cfg       *config.MainConfig
	DB        *sql.Store
	Publisher event.MessagePublisher
}

func newContainer(opts *options) *container {
	userRepo := dao.NewUserRepository(&dao.OptsUserRepository{DB: opts.DB})

	txMgr, err := txmanager.New(context.Background(), &txmanager.DriverConfig{
		Type: "sql",
		Config: txSql.Config{
			DB: opts.DB,
		},
	})
	if err != nil {
		log.Fatal("Failed to instantiate transaction manager ")
	}

	jwtSvc := jwtsvc.New(&jwtsvc.Options{
		Config: opts.Cfg.JWT,
	})

	userUc := user.New(&user.Opts{
		UserRepo:  userRepo,
		TxMgr:     txMgr,
		JwtSvc:    jwtSvc,
		Publisher: opts.Publisher,
		Topics:    opts.Cfg.Kafka.Topics,
	})

	return &container{
		Cfg:    *opts.Cfg,
		UserUC: userUc,
	}

}
