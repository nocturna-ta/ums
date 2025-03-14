package usecases

import (
	"context"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type AuthUseCases interface {
	ValidateAuthorization(ctx context.Context, req *request.ValidateAuthorizationRequest) *response.ValidateAuthorizationResponse
}
