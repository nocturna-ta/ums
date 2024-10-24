package user

import (
	"context"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/log"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

func (m *Module) GetUserByID(ctx context.Context) (*response.UserResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetUserByID")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := m.userRepo.GetById(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id":    reqCtx.UserId,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetUserByID] Failed to get user by id")
		return nil, err
	}

	return &response.UserResponse{
		ID:          user.ID.String(),
		Name:        user.Name,
		NIK:         user.NIK,
		NoTelephone: user.NoTelephone.String,
		Email:       user.Email.String,
	}, nil

}

func (m *Module) ChangePassword(ctx context.Context, req request.ChangeUserPasswordRequest) (*response.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Module) GetUserByNIK(ctx context.Context, nik string) (*response.UserResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "[UserUseCases.GetUserByNIK]")
	defer span.End()

	user, err := m.userRepo.GetByNIK(ctx, nik)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"nik":   nik,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetUserByNIK] Failed to get user by nik")
		return nil, err
	}

	return &response.UserResponse{
		ID:          user.ID.String(),
		Name:        user.Name,
		NIK:         user.NIK,
		NoTelephone: user.NoTelephone.String,
		Email:       user.Email.String,
	}, nil

}
