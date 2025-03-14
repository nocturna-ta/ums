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

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body request.UserRegistrationRequest true "User registration request"
// @Success		200	{object}	jsonResponse{data=response.UserRegistrationResponse}
// @Router /v1/user/register [post]
func (api *API) RegisterUser(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.RegisterUser")
	defer span.End()

	var regisReq request.UserRegistrationRequest
	err := json.Unmarshal(req.RawBody(), &regisReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = regisReq.ValidateRegistrationUser()
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	res, err := api.userUc.RegisterUser(ctx, &regisReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetUserByEmail godoc
// @Summary Get user by email
// @Description Get user by email
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} jsonResponse{data=response.UserResponse} "User found"
// @Router /v1/user/{email} [get]
func (api *API) GetUserByEmail(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetUserByEmail")
	defer span.End()

	email := req.Params("email")

	res, err := api.userUc.GetUserByEmail(ctx, email)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetByID godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} jsonResponse{data=response.UserResponse} "User found"
// @Router /v1/user/id/{id} [get]
func (api *API) GetByID(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetByID")
	defer span.End()

	id := req.Params("id")

	res, err := api.userUc.GetByID(ctx, id)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body request.UserUpdateRequest true "User update request"
// @Success 200 {object} jsonResponse{data=response.UserResponse} "User updated"
// @Router /v1/user/update/{id} [put]
func (api *API) UpdateUser(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.UpdateUser")
	defer span.End()

	id := req.Params("id")

	var userReq request.UserUpdateRequest
	err := json.Unmarshal(req.RawBody(), &userReq)

	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = userReq.ValidateUpdateUser()
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	res, err := api.userUc.UpdateUser(ctx, id, &userReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
