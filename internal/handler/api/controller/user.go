package controller

import (
	"context"
	"encoding/json"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
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
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
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
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Success 200 {object} jsonResponse{data=response.UserResponse} "User found"
// @Router /v1/user/me [get]
func (api *API) GetByID(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetByID")
	defer span.End()

	res, err := api.userUc.GetByID(ctx)
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
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Param user body request.UserUpdateRequest true "User update request"
// @Success 200 {object} jsonResponse{data=response.UserResponse} "User updated"
// @Router /v1/user/update [put]
func (api *API) UpdateUser(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.UpdateUser")
	defer span.End()

	var userReq request.UserUpdateRequest
	err := json.Unmarshal(req.RawBody(), &userReq)

	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = userReq.ValidateUpdateUser()
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	res, err := api.userUc.UpdateUser(ctx, &userReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// LoginUser godoc
// @Summary Login user
// @Description Login user
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Param user body request.UserLoginRequest true "User login request"
// @Success 200 {object} jsonResponse{data=response.UserLoginResponse} "User logged in"
// @Router /v1/user/login [post]
func (api *API) LoginUser(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.LoginUser")
	defer span.End()

	var loginReq request.UserLoginRequest
	err := json.Unmarshal(req.RawBody(), &loginReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = loginReq.ValidateLogin()
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	res, err := api.userUc.LoginUser(ctx, &loginReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change user password
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Param user body request.UserChangePasswordRequest true "User change password request"
// @Success 200 {object} jsonResponse{data=response.UserResponse} "Password changed"
// @Router /v1/user/change-password [put]
func (api *API) ChangePassword(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.ChangePassword")
	defer span.End()

	var changePassReq request.UserChangePasswordRequest
	err := json.Unmarshal(req.RawBody(), &changePassReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = changePassReq.ValidateChangePassword()
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = api.userUc.ChangePassword(ctx, &changePassReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData("Password changed"), nil

}

// CheckVerificationStatus godoc
// @Summary Check verification status
// @Description Check the verification status of a user account
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} jsonResponse{data=response.UserVerificationResponse} "User verification status"
// @Router /v1/user/verification-status/{email} [get]
func (api *API) CheckVerificationStatus(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.CheckVerificationStatus")
	defer span.End()

	email := req.Params("email")
	if email == "" {
		return cutresp.CustomErrorResponse(&custerr.ErrChain{
			Message: "Email is required",
			Code:    400,
			Type:    response.ErrBadRequest,
		})
	}

	res, err := api.userUc.CheckVerificationStatus(ctx, email)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetMyVerificationStatus godoc
// @Summary Get current user's verification status
// @Description Get the verification status and details for the currently logged in user
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Id header string true "User ID"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string true "Role"
// @Success 200 {object} jsonResponse{data=response.UserVerificationStatusResponse} "User verification status"
// @Router /v1/user/my-verification-status [get]
func (api *API) GetMyVerificationStatus(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetMyVerificationStatus")
	defer span.End()

	res, err := api.userUc.GetMyVerificationStatus(ctx)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
