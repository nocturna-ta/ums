package api

import (
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/handler/api/controller"
	"github.com/nocturna-ta/ums/internal/usecases"
)

type Options struct {
	Cfg             config.MainConfig
	VoterUc         usecases.VoterUseCases
	UserUc          usecases.UserUseCases
	KpuProvinsiUc   usecases.KPUProvinsiUseCases
	KpuKotaUc       usecases.KPUKotaUseCases
	UserLogUc       usecases.UserLogUseCases
	UserStatisticUc usecases.UserStatisticUseCases
}

type Handler struct {
	opts        *Options
	listenErrCh chan error
	myRouter    *router.FastRouter
}

func New(opts *Options) *Handler {
	handler := &Handler{
		opts: opts,
	}

	handler.myRouter = controller.New(&controller.Options{
		Prefix:          opts.Cfg.API.BasePath,
		Port:            opts.Cfg.Server.Port,
		ReadTimeout:     opts.Cfg.Server.ReadTimeout,
		WriteTimeout:    opts.Cfg.Server.WriteTimeout,
		RequestTimeout:  opts.Cfg.API.APITimeout,
		EnableSwagger:   opts.Cfg.API.EnableSwagger,
		VoterUc:         opts.VoterUc,
		UserUc:          opts.UserUc,
		KpuProvinsiUc:   opts.KpuProvinsiUc,
		KpuKotaUc:       opts.KpuKotaUc,
		UserLogUc:       opts.UserLogUc,
		UserStatisticUc: opts.UserStatisticUc,
	}).RegisterRoute()

	return handler
}
func (h *Handler) Run() {
	log.Infof("API Listening on %d", h.opts.Cfg.Server.Port)
	h.listenErrCh <- h.myRouter.StartServe()
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
