package server

import (
	"context"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/txmanager"
	txSql "github.com/nocturna-ta/golib/txmanager/sql"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/usecases"
	"github.com/nocturna-ta/ums/internal/usecases/user"
)

type container struct {
	Cfg    config.MainConfig
	UserUC usecases.UserUseCases
}

type options struct {
	Cfg *config.MainConfig
	DB  *sql.Store
}

func newContainer(opts *options) *container {
	userRepo := dao.NewUserRepository(&dao.OptsUserRepository{DB: opts.DB})

	_, err := txmanager.New(context.Background(), &txmanager.DriverConfig{
		Type: "sql",
		Config: txSql.Config{
			DB: opts.DB,
		},
	})
	if err != nil {
		log.Fatal("Failed to instantiate transaction manager ")
	}

	userUc := user.New(&user.Opts{
		UserRepo: userRepo,
	})

	//jwtSvc := jwtsvc

	return &container{
		Cfg:    *opts.Cfg,
		UserUC: userUc,
	}

}
