package controller

import (
	"context"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/tracing"
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
