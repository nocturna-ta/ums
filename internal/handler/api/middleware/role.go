package middleware

import (
	"context"
	"fmt"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/pkg/roles"
	"net/http"
)

func RequireRole(requiredRole string) func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
		return func(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
			span, ctx := tracing.StartSpanFromContext(ctx, "Middleware.RequireRole")
			defer span.End()

			reqCtx, err := libCtx.GetRequestContext(ctx)
			if err != nil {
				return nil, &custerr.ErrChain{
					Message: "Unauthorized: missing request context",
					Code:    http.StatusUnauthorized,
					Type:    response.ErrUnauthorized,
				}
			}

			if !reqCtx.HasRole(requiredRole) {
				return nil, &custerr.ErrChain{
					Message: "Forbidden: insufficient permissions",
					Code:    http.StatusForbidden,
					Type:    response.ErrForbiddenResource,
				}
			}

			return handler(ctx, req)
		}
	}
}

func RequireAnyRole(requiredRoles ...string) func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
		return func(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
			span, ctx := tracing.StartSpanFromContext(ctx, "Middleware.RequireAnyRole")
			defer span.End()

			reqCtx, err := libCtx.GetRequestContext(ctx)
			if err != nil {
				return nil, &custerr.ErrChain{
					Message: "Unauthorized: missing request context",
					Code:    http.StatusUnauthorized,
					Type:    response.ErrUnauthorized,
				}
			}

			if !reqCtx.HasAnyRole(requiredRoles...) {
				return nil, &custerr.ErrChain{
					Message: "Forbidden: insufficient permissions",
					Code:    http.StatusForbidden,
					Type:    response.ErrForbiddenResource,
				}
			}

			return handler(ctx, req)
		}
	}
}

func KPUPusat() func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return RequireRole(roles.RoleKPUPusat)
}

func VoterOnly() func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return RequireRole(roles.RoleVoter)
}

func KPUKotaOnly() func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return RequireRole(roles.RoleKPUKota)
}

func KPUProvinsiOnly() func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return RequireRole(roles.RoleKPUProvinsi)
}

func KPUOnly() func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return RequireAnyRole(roles.RoleKPUKota, roles.RoleKPUProvinsi)
}

func CanVerifyRole(targetRole string) func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
		return func(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
			span, ctx := tracing.StartSpanFromContext(ctx, "Middleware.CanVerifyRole")
			defer span.End()

			reqCtx, err := libCtx.GetRequestContext(ctx)
			if err != nil {
				return nil, &custerr.ErrChain{
					Message: "Unauthorized: missing request context",
					Code:    http.StatusUnauthorized,
					Type:    response.ErrUnauthorized,
				}
			}

			if !roles.CanVerify(reqCtx.GetRole(), targetRole) {
				return nil, &custerr.ErrChain{
					Message: fmt.Sprintf("Forbidden: %s cannot verify %s users",
						roles.GetRoleDisplayName(reqCtx.GetRole()),
						roles.GetRoleDisplayName(targetRole)),
					Code: http.StatusForbidden,
					Type: response.ErrForbiddenResource,
				}
			}

			return handler(ctx, req)
		}
	}
}

func CanVerifyKPUProvinsi() func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return CanVerifyRole(roles.RoleKPUProvinsi)
}

func CanVerifyKPUKota() func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return CanVerifyRole(roles.RoleKPUKota)
}

func CanVerifyVoter() func(handler router.Handler[rest.JSONResponse]) router.Handler[rest.JSONResponse] {
	return CanVerifyRole(roles.RoleVoter)
}
