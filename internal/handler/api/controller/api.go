package controller

import (
	"github.com/gofiber/swagger"
	"github.com/nocturna-ta/golib/router"
	_ "github.com/nocturna-ta/ums/docs"
	"github.com/nocturna-ta/ums/internal/usecases"
	"time"
)

type API struct {
	prefix         string
	port           uint
	readTimeout    time.Duration
	writeTimeout   time.Duration
	requestTimeout time.Duration
	enableSwagger  bool
	corsConfig     *router.CorsConfig
	voterUc        usecases.VoterUseCases
	kpuBranchUc    usecases.KPUBranchUseCases
}

type Options struct {
	Prefix         string
	Port           uint
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	RequestTimeout time.Duration
	EnableSwagger  bool
	CorsConfig     *router.CorsConfig
	VoterUc        usecases.VoterUseCases
	KpuBranchUc    usecases.KPUBranchUseCases
}

func New(opts *Options) *API {
	return &API{
		prefix:         opts.Prefix,
		port:           opts.Port,
		readTimeout:    opts.ReadTimeout,
		writeTimeout:   opts.WriteTimeout,
		requestTimeout: opts.RequestTimeout,
		enableSwagger:  opts.EnableSwagger,
		corsConfig:     opts.CorsConfig,
		voterUc:        opts.VoterUc,
		kpuBranchUc:    opts.KpuBranchUc,
	}
}

func (api *API) RegisterRoute() *router.FastRouter {
	myRouter := router.New(&router.Options{
		Prefix:         api.prefix,
		Port:           api.port,
		ReadTimeout:    api.readTimeout,
		WriteTimeout:   api.writeTimeout,
		RequestTimeout: api.requestTimeout,
		CorsConfig:     api.corsConfig,
	})

	if api.enableSwagger {
		myRouter.CustomHandler("GET", "/docs/*", swagger.HandlerDefault, router.MustAuthorized(false))
	}

	myRouter.GET("/health", api.Ping, router.MustAuthorized(false))
	myRouter.Group("/v1", func(v1 *router.FastRouter) {
		v1.Group("/voter", func(voter *router.FastRouter) {
			voter.POST("/register", api.RegisterVoter, router.MustAuthorized(false))
			voter.GET("/:nik", api.GetVoterByNIK, router.MustAuthorized(false))
			voter.GET("/address/:address", api.GetVoterByAddress, router.MustAuthorized(false))
			voter.GET("/region/:region", api.GetVoterByRegion, router.MustAuthorized(false))
			voter.GET("/", api.GetAllVoter, router.MustAuthorized(false))
		})
		v1.Group("/kpu-branch", func(kpuBranch *router.FastRouter) {
			kpuBranch.GET("/", api.GetAllKPUBranch, router.MustAuthorized(false))
			kpuBranch.POST("/register", api.RegisterKPUBranch, router.MustAuthorized(false))
			kpuBranch.GET("/address/:address", api.GetKPUBranchByAddress, router.MustAuthorized(false))
		})
	})

	return myRouter
}
