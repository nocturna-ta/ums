package controller

import (
	"context"
	"encoding/json"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/infrastructures/cutresp"
	"github.com/nocturna-ta/ums/internal/usecases/request"
)

// RegisterKPUKota godoc
// @Summary Register KPU Kota
// @Description Register KPU Kota
// @Tags kpu_kota
// @Accept json
// @Param users body request.KPUKotaRegistrationRequest true "Register Request"
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUKotaRegistrationResponse}
// @Router /v1/kpu-kota/register [post]
func (api *API) RegisterKPUKota(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.RegisterKPUKota")
	defer span.End()

	var regisReq request.KPUKotaRegistrationRequest
	err := json.Unmarshal(req.RawBody(), &regisReq)

	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = regisReq.ValidateRegistrationRequest()

	res, err := api.kpuKotaUc.RegisterKPUKota(ctx, &regisReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetAllKPUKota godoc
// @Summary Get All KPU Kota
// @Description Get All KPU Kota
// @Tags kpu_kota
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Accept json
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUKotaResponse}
// @Router /v1/kpu-kota [get]
func (api *API) GetAllKPUKota(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAllKPUKota")
	defer span.End()

	res, err := api.kpuKotaUc.GetAllKPUKota(ctx)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetKPUKotaByAddress godoc
// @Summary Get KPU Kota By Address
// @Description Get KPU Kota By Address
// @Tags kpu_kota
// @Accept json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUKotaResponse}
// @Router /v1/kpu-kota/address [get]
func (api *API) GetKPUKotaByAddress(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetKPUKotaByAddress")
	defer span.End()

	res, err := api.kpuKotaUc.GetKPUKotaByAddress(ctx)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
