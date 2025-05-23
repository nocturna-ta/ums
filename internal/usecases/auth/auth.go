package auth

import (
	"context"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/ums/pkg/constants"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (m *Module) ValidateAuthorization(ctx context.Context, req *request.ValidateAuthorizationRequest) (*response.ValidateAuthorizationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "AuthUseCases.ValidateAuthorization")
	defer span.End()

	authHeader, ok := req.Headers[constants.Authorization]
	if !ok || authHeader == "" {
		return nil, status.Error(codes.Unauthenticated, "Missing or empty authorization header")
	}

	token := strings.Replace(authHeader, "Bearer ", "", 1)
	if token == "" {
		return nil, status.Error(codes.Unauthenticated, "Missing token")
	}

	parsedToken, err := m.jwtSvc.Validate(ctx, token, jwtsvc.AccessType)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Token validation failed: %v", err)
	}

	claims, ok := parsedToken.(*jwtsvc.AccessClaims)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "Invalid token claims")
	}

	return &response.ValidateAuthorizationResponse{
		IsValid: true,
		ExplodeHeader: map[string]string{
			"X-User-Id": claims.UserID,
			"Role":      claims.Role,
		},
	}, nil
}
