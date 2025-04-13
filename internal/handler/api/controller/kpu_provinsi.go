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

// RegisterKPUProvinsi godoc
// @Summary Register KPU Provinsi
// @Description Register KPU Provinsi
// @Tags kpu_provinsi
// @Accept json
// @Param users body request.KPUProvinsiRegistrationRequest true "Register Request"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUProvinsiRegistrationResponse}
// @Router /v1/kpu-provinsi/register [post]
func (api *API) RegisterKPUProvinsi(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.RegisterKPUProvinsi")
	defer span.End()

	var regisReq request.KPUProvinsiRegistrationRequest
	err := json.Unmarshal(req.RawBody(), &regisReq)

	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = regisReq.ValidateRegistrationRequest()

	res, err := api.kpuProvinsiUc.RegisterKPUProvinsi(ctx, &regisReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetAllKPUProvinsi godoc
// @Summary Get All KPU Provinsi
// @Description Get All KPU Provinsi
// @Tags kpu_provinsi
// @Accept json
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUProvinsiResponse}
// @Router /v1/kpu-provinsi [get]
func (api *API) GetAllKPUProvinsi(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAllKPUProvinsi")
	defer span.End()

	res, err := api.kpuProvinsiUc.GetAllKPUProvinsi(ctx)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetKPUProvinsiByAddress godoc
// @Summary Get KPU Provinsi By Address
// @Description Get KPU Provinsi By Address
// @Tags kpu_provinsi
// @Accept json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUProvinsiResponse}
// @Router /v1/kpu-provinsi/address [get]
func (api *API) GetKPUProvinsiByAddress(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetKPUProvinsiByAddress")
	defer span.End()

	res, err := api.kpuProvinsiUc.GetKPUProvinsiByAddress(ctx)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
