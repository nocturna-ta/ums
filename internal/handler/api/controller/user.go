package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/infrastructures/cutresp"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/pkg/utils"
)

// GetUserByID godoc
// @Summary 	Get User By ID
// @Description	Get User BY ID
// @Tags		users
// @Accept		json
// @Param		X-Channel-Id			header		string	false 	"channel where request comes from"	default(web)
// @Param		X-User-Id				header		string 	false	"user id"
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.UserResponse}
// @Router		/v1/users/me	[get]
func (api *API) GetUserByID(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetUserByID")
	defer span.End()

	res, err := api.userUc.GetUserByID(ctx)
	if err != nil {
		return nil, err

	}

	return rest.NewJSONResponse().SetData(res), nil

}

// GetUserByNIK godoc
// @Summary 	Get User By NIK
// @Description	Get User BY NIK
// @Tags		users
// @Accept		json
// @Param		X-Channel-Id			header		string	false 	"channel where request comes from"	default(web)
// @Param		NIK						path		string 	true	"user nik"
// @Produce		json
// @Success		200	{object}	jsonResponse{data=response.UserResponse}
// @Router		/v1/users/{nik}	[get]
func (api *API) GetUserByNIK(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetUserByNIK")
	defer span.End()

	nik := utils.Encryption(req.Params("nik"))
	fmt.Println(nik)

	res, err := api.userUc.GetUserByNIK(ctx, nik)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}
	return rest.NewJSONResponse().SetData(res), nil
}

// Register godoc
// @Summary 	Register User
// @Description Register new user
// @Tags		users
// @Accept		json
// @Param		X-Channel-Id	header		string	false 	"channel where request comes from"	default(web)
// @Param		request		body 		request.UserRegisterRequest	true	"Register Request"
// @Produce	json
// @Success	200	{object}	jsonResponse{data=response.UserRegistrationResponse}
// @Router		/v1/register [post]
func (api *API) Register(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.Register")
	defer span.End()

	var regisReq request.UserRegisterRequest
	err := json.Unmarshal(req.RawBody(), &regisReq)

	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = regisReq.ValidateRegisterRequest()

	res, err := api.userUc.Register(ctx, &regisReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
