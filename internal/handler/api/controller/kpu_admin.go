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

// GetPendingVerifications godoc
// @Summary Get pending verification requests
// @Description Get all users with pending verification status
// @Tags Admin
// @Accept json
// @Produce json
// @Param X-User-Id header string true "Admin User ID"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string true "Admin Role"
// @Success 200 {object} jsonResponse{data=[]response.UserVerificationResponse} "Pending verification users"
// @Router /v1/admin/verifications/pending [get]
func (api *API) GetPendingVerifications(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetPendingVerifications")
	defer span.End()

	res, err := api.userUc.GetPendingVerificationUsers(ctx)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetPendingVerificationDetails godoc
// @Summary Get pending verification details for a specific user
// @Description Get detailed information about a user's pending verification including entity-specific data
// @Tags Admin
// @Accept json
// @Produce json
// @Param X-User-Id header string true "Admin User ID"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string true "Admin Role"
// @Param user_id path string true "User ID"
// @Success 200 {object} jsonResponse{data=response.UserVerificationDetailsResponse} "User verification details"
// @Router /v1/admin/verifications/details/{user_id} [get]
func (api *API) GetPendingVerificationDetails(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetPendingVerificationDetails")
	defer span.End()

	userID := req.Params("user_id")
	if userID == "" {
		return cutresp.CustomErrorResponse(&custerr.ErrChain{
			Message: "User ID is required",
			Code:    400,
			Type:    response.ErrBadRequest,
		})
	}

	res, err := api.userUc.GetVerificationDetails(ctx, userID)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// ApproveVerification godoc
// @Summary Approve a user verification request
// @Description Approve a pending user verification, activating their account with the requested role
// @Tags Admin
// @Accept json
// @Produce json
// @Param X-User-Id header string true "Admin User ID"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string true "Admin Role"
// @Param verification body request.UserVerificationRequest true "Verification approval request"
// @Success 200 {object} jsonResponse{data=string} "Success message"
// @Router /v1/admin/verifications/approve [post]
func (api *API) ApproveVerification(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.ApproveVerification")
	defer span.End()

	var verificationReq request.UserVerificationRequest
	err := json.Unmarshal(req.RawBody(), &verificationReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = verificationReq.ValidateVerificationRequest()
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = api.userUc.ApproveUserVerification(ctx, &verificationReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData("User verification approved successfully"), nil
}

// RejectVerification godoc
// @Summary Reject a user verification request
// @Description Reject a pending user verification, keeping their account as unverified
// @Tags Admin
// @Accept json
// @Produce json
// @Param X-User-Id header string true "Admin User ID"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string true "Admin Role"
// @Param verification body request.UserVerificationRequest true "Verification rejection request"
// @Success 200 {object} jsonResponse{data=string} "Success message"
// @Router /v1/admin/verifications/reject [post]
func (api *API) RejectVerification(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.RejectVerification")
	defer span.End()

	var verificationReq request.UserVerificationRequest
	err := json.Unmarshal(req.RawBody(), &verificationReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = verificationReq.ValidateVerificationRequest()
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	err = api.userUc.RejectUserVerification(ctx, &verificationReq)
	if err != nil {
		return cutresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData("User verification rejected"), nil
}
