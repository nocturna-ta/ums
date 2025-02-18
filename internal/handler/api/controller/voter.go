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

// RegisterVoter godoc
// @Summary 	Register Voter
// @Description Register Voter
// @Tags		voters
// @Accept		json
// @Param		users		body 		request.VoterRegistrationRequest	true	"Register Request"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterRegistrationResponse}
// @Router		/v1/voter/register [post]
func (api *API) RegisterVoter(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.Register")
	defer span.End()

	var regisReq request.VoterRegistrationRequest
	err := json.Unmarshal(req.RawBody(), &regisReq)

	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = regisReq.ValidateRegisterRequest()

	res, err := api.voterUc.RegisterVoter(ctx, &regisReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetVoterByNIK godoc
// @Summary 	Get Voter By NIK
// @Description Get Voter By NIK
// @Tags		voters
// @Accept		json
// @Param		nik		path 		string	true	"NIK"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterResponse}
// @Router		/v1/voter/{nik} [get]
func (api *API) GetVoterByNIK(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetVoterByNIK")
	defer span.End()

	nik := req.Params("nik")

	res, err := api.voterUc.GetVoterByNIK(ctx, nik)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetVoterByAddress godoc
// @Summary 	Get Voter By Address
// @Description Get Voter By Address
// @Tags		voters
// @Accept		json
// @Param		address		path 		string	true	"Address"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterResponse}
// @Router		/v1/voter/address/{address} [get]
func (api *API) GetVoterByAddress(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetVoterByAddress")
	defer span.End()

	address := req.Params("address")

	res, err := api.voterUc.GetVoterByAddress(ctx, address)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetVoterByRegion godoc
// @Summary 	Get Voter By Region
// @Description Get Voter By Region
// @Tags		voters
// @Accept		json
// @Param		region		path 		string	true	"Region"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterResponse}
// @Router		/v1/voter/region/{region} [get]
func (api *API) GetVoterByRegion(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetVoterByRegion")
	defer span.End()

	region := req.Params("region")

	res, err := api.voterUc.GetVoterByRegion(ctx, region)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetAllVoter godoc
// @Summary 	Get All Voter
// @Description Get All Voter
// @Tags		voters
// @Accept		json
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterResponse}
// @Router		/v1/voter [get]
func (api *API) GetAllVoter(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAllVoter")
	defer span.End()

	res, err := api.voterUc.GetAllVoter(ctx)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
