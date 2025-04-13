package auth

import (
	"context"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/ums/pkg/constants"
	"strings"
)

func (m *Module) ValidateAuthorization(ctx context.Context, req *request.ValidateAuthorizationRequest) *response.ValidateAuthorizationResponse {
	span, ctx := tracing.StartSpanFromContext(ctx, "AuthUseCases.ValidateAuthorization")
	defer span.End()

	token := strings.Replace(req.Headers[constants.Authorization], "Bearer ", "", 1)
	if token == constants.EmptyString {
		return &response.ValidateAuthorizationResponse{}
	}

	parsedToken, err := m.jwtSvc.Validate(ctx, token, jwtsvc.AccessType)
	if err != nil {
		return &response.ValidateAuthorizationResponse{}
	}

	claims, ok := parsedToken.(*jwtsvc.AccessClaims)
	if !ok {
		return &response.ValidateAuthorizationResponse{}
	}

	return &response.ValidateAuthorizationResponse{
		IsValid: true,
		ExplodeHeader: map[string]string{
			"UserID": claims.UserID,
			"Role":   claims.Role,
		},
	}
}
