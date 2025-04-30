package controller

import (
	"github.com/gofiber/swagger"
	"github.com/nocturna-ta/golib/router"
	_ "github.com/nocturna-ta/ums/docs"
	"github.com/nocturna-ta/ums/internal/handler/api/middleware"
	"github.com/nocturna-ta/ums/internal/usecases"
	"github.com/nocturna-ta/ums/pkg/roles"
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
	userUc         usecases.UserUseCases
	kpuProvinsiUc  usecases.KPUProvinsiUseCases
	kpuKotaUc      usecases.KPUKotaUseCases
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
	UserUc         usecases.UserUseCases
	KpuProvinsiUc  usecases.KPUProvinsiUseCases
	KpuKotaUc      usecases.KPUKotaUseCases
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
		userUc:         opts.UserUc,
		kpuProvinsiUc:  opts.KpuProvinsiUc,
		kpuKotaUc:      opts.KpuKotaUc,
	}
}

func (api *API) RegisterRoute() *router.FastRouter {
	corsConfig := &router.CorsConfig{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
		AllowHeaders:     "Content-Type, Authorization, x-user-id",
		AllowCredentials: false,
		MaxAge:           300,
	}
	myRouter := router.New(&router.Options{
		Prefix:         api.prefix,
		Port:           api.port,
		ReadTimeout:    api.readTimeout,
		WriteTimeout:   api.writeTimeout,
		RequestTimeout: api.requestTimeout,
		CorsConfig:     corsConfig,
	})

	if api.enableSwagger {
		myRouter.CustomHandler("GET", "/docs/*", swagger.HandlerDefault, router.MustAuthorized(false))
	}

	myRouter.GET("/health", api.Ping, router.MustAuthorized(false))
	myRouter.Group("/v1", func(v1 *router.FastRouter) {
		v1.Group("/voter", func(voter *router.FastRouter) {
			voter.POST("/register", api.RegisterVoter, router.MustAuthorized(false))
			voter.GET("/nik/:nik", api.GetVoterByNIK, router.MustAuthorized(false))
			voter.GET("/address", api.GetVoterByAddress, router.MustAuthorized(false))
			voter.GET("/region/:region", api.GetVoterByRegion, router.MustAuthorized(false))
			voter.GET("/", api.GetAllVoter, router.MustAuthorized(false))
		})
		v1.Group("/kpu-provinsi", func(kpuProvinsi *router.FastRouter) {
			kpuProvinsi.GET("/", middleware.KPUProvinsiOnly()(api.GetAllKPUProvinsi), router.MustAuthorized(false))
			kpuProvinsi.POST("/register", api.RegisterKPUProvinsi, router.MustAuthorized(false))
			kpuProvinsi.GET("/address", api.GetKPUProvinsiByAddress, router.MustAuthorized(false))
			kpuProvinsi.GET("/:id", api.GetKPUProvinsiByID, router.MustAuthorized(false))
			kpuProvinsi.POST("/:id/photo", api.UploadKPUProvinsiPhoto, router.MustAuthorized(false))
			kpuProvinsi.ATTACHMENT("/:id/photo", api.GetKPUProvinsiPhoto, router.MustAuthorized(false))
			kpuProvinsi.PUT("/update", api.UpdateKPUProvinsi, router.MustAuthorized(false))

		})
		v1.Group("/kpu-kota", func(kpuKota *router.FastRouter) {
			kpuKota.GET("/", api.GetAllKPUKota, router.MustAuthorized(false))
			kpuKota.POST("/register", api.RegisterKPUKota, router.MustAuthorized(false))
			kpuKota.GET("/address", api.GetKPUKotaByAddress, router.MustAuthorized(false))
			kpuKota.GET("/:id", api.GetKPUKotaByID, router.MustAuthorized(false))
			kpuKota.POST("/:id/photo", api.UploadKPUKotaPhoto, router.MustAuthorized(false))
			kpuKota.ATTACHMENT("/:id/photo", api.GetKPUKotaPhoto, router.MustAuthorized(false))
			kpuKota.PUT("/update", api.UpdateKPUKota, router.MustAuthorized(false))
		})
		v1.Group("/user", func(user *router.FastRouter) {
			user.POST("/register", api.RegisterUser, router.MustAuthorized(false))
			user.GET("/me", api.GetByID)
			user.PUT("/update", api.UpdateUser)
			user.POST("/login", api.LoginUser, router.MustAuthorized(false))

			user.GET("/verification-status/:email", api.CheckVerificationStatus, router.MustAuthorized(false))
			user.GET("/my-verification-status", api.GetMyVerificationStatus)
		})

		v1.Group("/verifications", func(verifications *router.FastRouter) {
			verifications.GET("/pending", middleware.RequireAnyRole(
				roles.RoleKPUPusat,
				roles.RoleKPUProvinsi,
				roles.RoleKPUKota)(api.GetPendingVerificationsForRole))
			verifications.GET("/details/:user_id", middleware.RequireAnyRole(
				roles.RoleKPUPusat,
				roles.RoleKPUProvinsi,
				roles.RoleKPUKota)(api.GetPendingVerificationDetails))
			verifications.POST("/approve", middleware.RequireAnyRole(
				roles.RoleKPUPusat,
				roles.RoleKPUProvinsi,
				roles.RoleKPUKota)(api.ApproveVerification))
			verifications.POST("/reject", middleware.RequireAnyRole(
				roles.RoleKPUPusat,
				roles.RoleKPUProvinsi,
				roles.RoleKPUKota)(api.RejectVerification))
		})

		v1.Group("/admin", func(admin *router.FastRouter) {
			admin.Group("/verifications", func(verifications *router.FastRouter) {
				verifications.GET("/pending", middleware.KPUPusat()(api.GetPendingVerifications))
				verifications.GET("/details/:user_id", middleware.KPUPusat()(api.GetPendingVerificationDetails))
				verifications.POST("/approve", middleware.KPUPusat()(api.ApproveVerification))
				verifications.POST("/reject", middleware.KPUPusat()(api.RejectVerification))
			})
		})
	})

	return myRouter
}
