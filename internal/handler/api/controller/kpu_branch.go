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

// RegisterKPUBranch godoc
// @Summary Register KPU Branch
// @Description Register KPU Branch
// @Tags kpu_branch
// @Accept json
// @Param users body request.KPUBranchRegistrationRequest true "Register Request"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUBranchRegistrationResponse}
// @Router /v1/kpu-branch/register [post]
func (api *API) RegisterKPUBranch(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.RegisterKPUBranch")
	defer span.End()

	var regisReq request.KPUBranchRegistrationRequest
	err := json.Unmarshal(req.RawBody(), &regisReq)

	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = regisReq.ValidateRegistrationRequest()

	res, err := api.kpuBranchUc.RegisterKPUBranch(ctx, &regisReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetAllKPUBranch godoc
// @Summary Get All KPU Branch
// @Description Get All KPU Branch
// @Tags kpu_branch
// @Accept json
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUBranchResponse}
// @Router /v1/kpu-branch [get]
func (api *API) GetAllKPUBranch(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAllKPUBranch")
	defer span.End()

	res, err := api.kpuBranchUc.GetAllKPUBranch(ctx)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetKPUBranchByAddress godoc
// @Summary Get KPU Branch By Address
// @Description Get KPU Branch By Address
// @Tags kpu_branch
// @Accept json
// @Param address path string true "Address"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUBranchResponse}
// @Router /v1/kpu-branch/address/{address} [get]
func (api *API) GetKPUBranchByAddress(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetKPUBranchByAddress")
	defer span.End()

	address := req.Params("address")

	res, err := api.kpuBranchUc.GetKPUBranchByAddress(ctx, address)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
