package controller

import (
	"context"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/infrastructures/custresp"
	"strconv"
)

// GetLogs godoc
// @Summary      Get user logs
// @Description  Get user logs with pagination
// @Tags         User Log
// @Accept       json
// @Produce      json
// @Param        limit  query int  false "Limit of logs to return (default: 10, max: 1000)"
// @Param        offset query int  false "Offset for pagination (default: 0)"
// @Success      200  {object}  jsonResponse{data=[]response.UserLogResponse}
// @Router       /v1/user-logs [get]
func (api *API) GetLogs(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "LogsController.GetLogs")
	defer span.End()

	limit, err := strconv.Atoi(req.Query("limit"))
	offset, err := strconv.Atoi(req.Query("offset"))
	if err != nil {
		return custresp.CustomErrorResponse(&custerr.ErrChain{
			Message: "invalid limit or offset",
			Code:    400,
		})
	}

	userLogs, err := api.userLogUc.GetAllUserLog(ctx, limit, offset)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(userLogs), nil
}
