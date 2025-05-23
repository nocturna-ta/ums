package controller

import (
	"context"
	"github.com/nocturna-ta/api-gateway-grpc-lib/proto"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/usecases/request"
)

func (s *server) ValidateAuthorization(ctx context.Context, req *proto.AuthValidateRequest) (*proto.AuthValidateResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "GrpcServer.ValidateAuthorization")
	defer span.End()

	res, err := s.authUc.ValidateAuthorization(ctx, &request.ValidateAuthorizationRequest{
		Headers:       req.Headers,
		Path:          req.Path,
		TargetService: req.TargetService,
	})

	if err != nil {
		return nil, err
	}

	return &proto.AuthValidateResponse{
		IsValid:       res.IsValid,
		ExplodeHeader: res.ExplodeHeader,
	}, nil
}
