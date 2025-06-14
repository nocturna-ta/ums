package controller

import (
	"context"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/infrastructures/custresp"
)

// GetApprovedDPTStatistic godoc
// @Summary Get Approved DPT Statistic
// @Description Get the percentage of approved DPTs and total DPTs
// @Tags UserStatistic
// @Accept json
// @Produce json
// @Param region query string false "Region to filter the statistic"
// @Success 200 {object} jsonResponse{data=response.ApprovedDPTResponse}
// @Router /v1/user-statistic/approved-dpt [get]
func (api *API) GetApprovedDPTStatistic(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticController.GetApprovedDPTStatistic")
	defer span.End()

	region := req.Query("region")

	res, err := api.userStatisticUc.GetApprovedDPTStatistic(ctx, region)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetRejectedDPTStatistic godoc
// @Summary Get Rejected DPT Statistic
// @Description Get the percentage of rejected DPTs and total DPTs
// @Tags UserStatistic
// @Accept json
// @Produce json
// @Param region query string false "Region to filter the statistic"
// @Success 200 {object} jsonResponse{data=response.RejectedDPTResponse}
// @Router /v1/user-statistic/rejected-dpt [get]
func (api *API) GetRejectedDPTStatistic(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticController.GetRejectedDPTStatistic")
	defer span.End()

	region := req.Query("region")

	res, err := api.userStatisticUc.GetRejectedDPTStatistic(ctx, region)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetPendingDPTStatistic godoc
// @Summary Get Pending DPT Statistic
// @Description Get the percentage of pending DPTs and total DPTs
// @Tags UserStatistic
// @Accept json
// @Produce json
// @Param region query string false "Region to filter the statistic"
// @Success 200 {object} jsonResponse{data=response.PendingDPTResponse}
// @Router /v1/user-statistic/pending-dpt [get]
func (api *API) GetPendingDPTStatistic(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticController.GetPendingDPTStatistic")
	defer span.End()

	region := req.Query("region")

	res, err := api.userStatisticUc.GetPendingDPTStatistic(ctx, region)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetKPUProvinceStaffStatistic godoc
// @Summary Get KPU Staff Statistic
// @Description Get the percentage of KPU staff and total DPTs
// @Tags UserStatistic
// @Accept json
// @Produce json
// @Param region query string false "Region to filter the statistic"
// @Success 200 {object} jsonResponse{data=response.StaffKPUResponse}
// @Router /v1/user-statistic/kpu-provinsi-staff [get]
func (api *API) GetKPUProvinceStaffStatistic(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticController.GetKPuStaffStatistic")
	defer span.End()

	region := req.Query("region")

	res, err := api.userStatisticUc.GetStaffKPUProvinceStatistic(ctx, region)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetKPUKotaStaffStatistic godoc
// @Summary Get KPU Kota Staff Statistic
// @Description Get the percentage of KPU Kota staff and total DPTs
// @Tags UserStatistic
// @Accept json
// @Produce json
// @Param region query string false "Region to filter the statistic"
// @Success 200 {object} jsonResponse{data=response.StaffKPUResponse}
// @Router /v1/user-statistic/kpu-kota-staff [get]
func (api *API) GetKPUKotaStaffStatistic(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticController.GetKPUKotaStaffStatistic")
	defer span.End()

	region := req.Query("region")

	res, err := api.userStatisticUc.GetStaffKPUKotaStatistic(ctx, region)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetTotalDPTStatistic godoc
// @Summary Get Total DPT Statistic
// @Description Get the total number of DPTs
// @Tags UserStatistic
// @Accept json
// @Produce json
// @Param region query string false "Region to filter the statistic"
// @Success 200 {object} jsonResponse{data=response.TotalDPTResponse}
// @Router /v1/user-statistic/total-dpt [get]
func (api *API) GetTotalDPTStatistic(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticController.GetTotalDPTStatistic")
	defer span.End()

	region := req.Query("region")

	res, err := api.userStatisticUc.GetTotalDPTStatistic(ctx, region)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetProvinceInformationDPTStatistic godoc
// @Summary Get Province Information DPT Statistic
// @Description Get the percentage of DPTs with province information
// @Tags UserStatistic
// @Accept json
// @Produce json
// @Success 200 {object} jsonResponse{data=[]response.DPTInformationResponse}
// @Router /v1/user-statistic/province-information-dpt [get]
func (api *API) GetProvinceInformationDPTStatistic(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticController.GetProvinceInformationDPTStatistic")
	defer span.End()

	res, err := api.userStatisticUc.GetProvinceInformationDPTStatistic(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetKotaInformationDPTStatistic godoc
// @Summary Get Kota Information DPT Statistic
// @Description Get the percentage of DPTs with kota information
// @Tags UserStatistic
// @Accept json
// @Produce json
// @Success 200 {object} jsonResponse{data=[]response.DPTInformationResponse}
// @Router /v1/user-statistic/kota-information-dpt [get]
func (api *API) GetKotaInformationDPTStatistic(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticController.GetKotaInformationDPTStatistic")
	defer span.End()

	res, err := api.userStatisticUc.GetKotaInformationDPTStatistic(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetVotedStatistic godoc
// @Summary Get Voted Statistic
// @Description Get the percentage of users who have voted
// @Tags UserStatistic
// @Accept json
// @Produce json
// @Param region query string false "Region to filter the statistic"
// @Success 200 {object} jsonResponse{data=response.VotedStatisticResponse}
// @Router /v1/user-statistic/voted [get]
func (api *API) GetVotedStatistic(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserStatisticController.GetVotedStatistic")
	defer span.End()

	region := req.Query("region")

	res, err := api.userStatisticUc.GetVotedStatistic(ctx, region)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
