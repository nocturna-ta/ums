package user

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/log"
	response2 "github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/ums/pkg/constants/errorcode"
	"github.com/nocturna-ta/ums/pkg/utils"
	"time"
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
		NIK:         utils.Encryption(user.NIK),
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
		NIK:         utils.Encryption(user.NIK),
		NoTelephone: user.NoTelephone.String,
		Email:       user.Email.String,
	}, nil

}

func (m *Module) Register(ctx context.Context, req *request.UserRegisterRequest) (*response.UserRegistrationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "[UserUseCases.Register]")
	defer span.End()

	var (
		user *model.User
		err  error
	)

	user = model.ConstructRegistration(req)

	if err := m.userRepo.Insert(ctx, user); err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"user":  user,
		}).ErrorWithCtx(ctx, "[UserUseCases.Register] Failed to insert user")

		if errors.Is(err, dao.ErrDuplicate) {
			return nil, &custerr.ErrChain{
				Message: "User already exists",
				Code:    400,
				Cause:   err,
				Type:    response2.ErrBadRequest,
			}
		}

		return nil, &custerr.ErrChain{
			Message: "Failed to insert user",
			Code:    500,
			Type:    response2.ErrInternalServerError,
			Cause:   err,
		}

	}
	return &response.UserRegistrationResponse{
		NIK: req.NIK,
	}, err
}

func (m *Module) Login(ctx context.Context, req *request.UserLoginRequest) (*response.UserLoginResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "[UserUseCases.Login]")
	defer span.End()

	existing, err := m.userRepo.GetByNIK(ctx, req.NIK)
	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"request": req,
		}).ErrorWithCtx(ctx, "[UserUseCases.Login] Failed to get user by nik")
		return nil, err
	}

	password := utils.PasswordHash(req.Password, existing.PasswordSalt)
	if password != existing.Password {
		return nil, &custerr.ErrChain{
			Message: errorcode.WrongPassword.Message,
			Code:    errorcode.WrongPassword.Code,
		}
	}

	expiresAt := jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour))
	token, err := m.jwtSvc.GenerateToken(ctx, &jwtsvc.AccessClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			ID:        uuid.New().String(),
		},
		JwtData: &jwtsvc.JwtData{
			UserID: existing.ID.String(),
		},
	})

	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"request": req,
		}).ErrorWithCtx(ctx, "[UserUseCases.Login] Failed to generate token")
		return nil, err
	}

	return &response.UserLoginResponse{
		Token:     token,
		ExpiresAt: &expiresAt.Time,
	}, nil
}
