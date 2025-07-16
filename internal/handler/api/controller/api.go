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
	prefix          string
	port            uint
	readTimeout     time.Duration
	writeTimeout    time.Duration
	requestTimeout  time.Duration
	enableSwagger   bool
	voterUc         usecases.VoterUseCases
	userUc          usecases.UserUseCases
	kpuProvinsiUc   usecases.KPUProvinsiUseCases
	kpuKotaUc       usecases.KPUKotaUseCases
	userLogUc       usecases.UserLogUseCases
	userStatisticUc usecases.UserStatisticUseCases
}

type Options struct {
	Prefix          string
	Port            uint
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	RequestTimeout  time.Duration
	EnableSwagger   bool
	VoterUc         usecases.VoterUseCases
	UserUc          usecases.UserUseCases
	KpuProvinsiUc   usecases.KPUProvinsiUseCases
	KpuKotaUc       usecases.KPUKotaUseCases
	UserLogUc       usecases.UserLogUseCases
	UserStatisticUc usecases.UserStatisticUseCases
}

func New(opts *Options) *API {
	return &API{
		prefix:          opts.Prefix,
		port:            opts.Port,
		readTimeout:     opts.ReadTimeout,
		writeTimeout:    opts.WriteTimeout,
		requestTimeout:  opts.RequestTimeout,
		enableSwagger:   opts.EnableSwagger,
		voterUc:         opts.VoterUc,
		userUc:          opts.UserUc,
		kpuProvinsiUc:   opts.KpuProvinsiUc,
		kpuKotaUc:       opts.KpuKotaUc,
		userLogUc:       opts.UserLogUc,
		userStatisticUc: opts.UserStatisticUc,
	}
}

func (api *API) RegisterRoute() *router.FastRouter {

	myRouter := router.New(&router.Options{
		Prefix:         api.prefix,
		Port:           api.port,
		ReadTimeout:    api.readTimeout,
		WriteTimeout:   api.writeTimeout,
		RequestTimeout: api.requestTimeout,
	})

	if api.enableSwagger {
		myRouter.CustomHandler("GET", "/user/docs/*", swagger.HandlerDefault, router.MustAuthorized(false))
	}

	myRouter.GET("/health", api.Ping, router.MustAuthorized(false))
	myRouter.Group("/v1", func(v1 *router.FastRouter) {
		v1.GET("/kpu-pusat/id", api.GetKPUPusat)
		v1.Group("/voter", func(voter *router.FastRouter) {
			voter.POST("/register", api.RegisterVoter, router.MustAuthorized(false))
			voter.GET("/nik/:nik", api.GetVoterByNIK)
			voter.GET("/address", api.GetVoterByAddress)
			voter.GET("/region/:region", api.GetVoterByRegion)
			voter.GET("/", api.GetAllVoter)
			voter.ATTACHMENT("/:id/photo", api.GetVoterKTPPhoto)
			voter.GET("/province", api.GetVoterByProvince)
			voter.GET("/kpu-kota", api.GetVoterByKPUKoota)
		})
		v1.Group("/kpu-provinsi", func(kpuProvinsi *router.FastRouter) {
			kpuProvinsi.GET("/", api.GetAllKPUProvinsi)
			kpuProvinsi.GET("/id", api.GetKPUProvinsiByUserID)
			kpuProvinsi.POST("/photo", api.UploadKPUProvinsiPhoto)
			kpuProvinsi.ATTACHMENT("/photo", api.GetKPUProvinsiPhoto)
			kpuProvinsi.ATTACHMENT("/photo/:id", api.GetKPUProvinsiPhotoUseID)
			kpuProvinsi.PUT("/update", api.UpdateKPUProvinsi)

		})
		v1.Group("/kpu-kota", func(kpuKota *router.FastRouter) {
			kpuKota.GET("/", api.GetAllKPUKota)
			kpuKota.GET("/id", api.GetKPUKotaByUserID)
			kpuKota.POST("/photo", api.UploadKPUKotaPhoto)
			kpuKota.ATTACHMENT("/photo", api.GetKPUKotaPhoto)
			kpuKota.ATTACHMENT("/photo/:id", api.GetKPUKotaPhotoUseID)
			kpuKota.PUT("/update", api.UpdateKPUKota)
		})
		v1.Group("/user", func(user *router.FastRouter) {
			user.POST("/register", api.RegisterUser, router.MustAuthorized(false))
			user.GET("/me", api.GetByID)
			user.PUT("/change-password", api.ChangePassword, router.MustAuthorized(false))
			user.POST("/login", api.LoginUser, router.MustAuthorized(false))
			user.GET("/:email", api.GetUserByEmail)
			user.GET("/verification-status/:email", api.CheckVerificationStatus)
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
		v1.Group("/user-logs", func(userLogs *router.FastRouter) {
			userLogs.GET("/", api.GetLogs, router.MustAuthorized(false))
		})
		v1.Group("/user-statistic", func(userStatistic *router.FastRouter) {
			userStatistic.GET("/approved-dpt", api.GetApprovedDPTStatistic)
			userStatistic.GET("/rejected-dpt", api.GetRejectedDPTStatistic)
			userStatistic.GET("/pending-dpt", api.GetPendingDPTStatistic)
			userStatistic.GET("/total-dpt", api.GetTotalDPTStatistic)
			userStatistic.GET("/kpu-provinsi-staff", api.GetKPUProvinceStaffStatistic)
			userStatistic.GET("/kpu-kota-staff", api.GetKPUKotaStaffStatistic)
			userStatistic.GET("/province-information-dpt", api.GetProvinceInformationDPTStatistic)
			userStatistic.GET("/kota-information-dpt", api.GetKotaInformationDPTStatistic)
			userStatistic.GET("/voted", api.GetVotedStatistic)
		})
	})

	return myRouter
}
