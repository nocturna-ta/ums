package middleware

import (
	"context"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"net/http"
)

func RequireRole(requiredRole string) func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
		return func(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
			// Get request context
			reqCtx, err := libCtx.GetRequestContext(ctx)
			if err != nil {
				return nil, &custerr.ErrChain{
					Message: "Unauthorized: missing request context",
					Code:    http.StatusUnauthorized,
					Type:    response.ErrUnauthorized,
				}
			}

			// Get user role from metadata
			role, exists := reqCtx.GetMetadata("Role")
			if !exists {
				return nil, &custerr.ErrChain{
					Message: "Unauthorized: role not found",
					Code:    http.StatusUnauthorized,
					Type:    response.ErrUnauthorized,
				}
			}

			// Check if role matches required role
			if role != requiredRole {
				return nil, &custerr.ErrChain{
					Message: "Forbidden: insufficient permissions",
					Code:    http.StatusForbidden,
					Type:    response.ErrForbidden,
				}
			}

			// Continue to the handler
			return handler(ctx, req)
		}
	}
}
