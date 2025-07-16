package controller

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/infrastructures/custresp"
	"github.com/nocturna-ta/ums/internal/usecases/request"
)

// RegisterVoter godoc
// @Summary 	Register Voter
// @Description Register Voter
// @Tags		voters
// @Accept		json
// @Param		users		body 		request.VoterRegistrationRequest	true	"Register Request"
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterRegistrationResponse}
// @Router		/v1/voter/register [post]
func (api *API) RegisterVoter(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.Register")
	defer span.End()

	var regisReq request.VoterRegistrationRequest
	err := json.Unmarshal(req.RawBody(), &regisReq)

	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	err = regisReq.ValidateRegisterRequest()

	res, err := api.voterUc.RegisterVoter(ctx, &regisReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetVoterByNIK godoc
// @Summary 	Get Voter By NIK
// @Description Get Voter By NIK
// @Tags		voters
// @Accept		json
// @Param		nik		path 		string	true	"NIK"
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterResponse}
// @Router		/v1/voter/nik/{nik} [get]
func (api *API) GetVoterByNIK(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetVoterByNIK")
	defer span.End()

	nik := req.Params("nik")

	res, err := api.voterUc.GetVoterByNIK(ctx, nik)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetVoterByAddress godoc
// @Summary 	Get Voter By Address
// @Description Get Voter By Address
// @Tags		voters
// @Accept		json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterResponse}
// @Router		/v1/voter/address [get]
func (api *API) GetVoterByAddress(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetVoterByAddress")
	defer span.End()

	res, err := api.voterUc.GetVoterByAddress(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetVoterByRegion godoc
// @Summary 	Get Voter By Region
// @Description Get Voter By Region
// @Tags		voters
// @Accept		json
// @Param		region		path 		string	true	"Region"
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterResponse}
// @Router		/v1/voter/region/{region} [get]
func (api *API) GetVoterByRegion(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetVoterByRegion")
	defer span.End()

	region := req.Params("region")

	res, err := api.voterUc.GetVoterByRegion(ctx, region)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetAllVoter godoc
// @Summary 	Get All Voter
// @Description Get All Voter
// @Tags		voters
// @Accept		json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.VoterResponse}
// @Router		/v1/voter [get]
func (api *API) GetAllVoter(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAllVoter")
	defer span.End()

	res, err := api.voterUc.GetAllVoter(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetVoterKTPPhoto godoc
// @Summary Get KTP photo for Voter
// @Description Get the photo for a KTP Photo
// @Tags voters
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Param id path string true "Voter ID"
// @Produce octet-stream
// @Success 200
// @Router /v1/voter/{id}/photo [get]
func (api *API) GetVoterKTPPhoto(ctx context.Context, req *router.Request) (*rest.AttachmentResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetVoterKTPPhoto")
	defer span.End()

	voterID, err := uuid.Parse(req.Params("id"))
	if err != nil {
		return nil, &custerr.ErrChain{
			Message: "Invalid Voter ID",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	file, contentType, err := api.voterUc.GetVoterKTPPhoto(ctx, voterID)
	if err != nil {
		return nil, err
	}

	return rest.NewAttachmentResponse().
		SetFile(file).
		SetFileName(file.FileName).
		SetContentType(contentType), nil
}

// GetVoterByProvince godoc
// @Summary Get Voter by Province
// @Description Get Voter by Province
// @Tags voters
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce octet-stream
// @Success 200
// @Router /v1/voter/province [get]
func (api *API) GetVoterByProvince(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetVoteByProvince")
	defer span.End()

	res, err := api.voterUc.GetVoterByProvince(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetVoterByKPUKoota godoc
// @Summary Get Voter by Province
// @Description Get Voter by Province
// @Tags voters
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce json
// @Success 200
// @Router /v1/voter/kpu-kota [get]
func (api *API) GetVoterByKPUKoota(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetVoterByKPUKoota")
	defer span.End()

	res, err := api.voterUc.GetVoterByKPUKota(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
