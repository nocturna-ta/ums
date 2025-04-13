package user

import (
	"context"
	"errors"
	"fmt"
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
	"github.com/nocturna-ta/ums/pkg/constants"
	"github.com/nocturna-ta/ums/pkg/constants/errorcode"
	"github.com/nocturna-ta/ums/pkg/utils"
	"time"
)

func (m *Module) RegisterUser(ctx context.Context, req *request.UserRegistrationRequest) (*response.UserRegistrationResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.RegisterUser")
	defer span.End()

	var (
		user *model.User
	)

	transaction := func(txCtx context.Context) (any, error) {
		user = model.ConstructUserRegistration(req)

		errTx := m.userRepo.Insert(txCtx, user)
		if errTx != nil {
			if errors.Is(errTx, dao.ErrDuplicate) {
				return nil, &custerr.ErrChain{
					Message: errorcode.UserAlreadyExists.Message,
					Code:    errorcode.UserAlreadyExists.Code,
					Type:    response2.ErrBadRequest,
					Cause:   errTx,
				}
			}

			return nil, errTx
		}

		errTx = m.publisher.Publish(txCtx, m.topics.MasterDataUser.Value, user.ID.String(), user.ToMessageModel(), map[string]any{
			constants.MetaDataOperation: constants.Create,
		})

		if errTx != nil {
			return nil, errTx
		}

		return nil, nil

	}
	_, err := m.txMgr.Execute(ctx, transaction, nil)
	if err != nil {
		return nil, err
	}

	return &response.UserRegistrationResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (m *Module) GetUserByEmail(ctx context.Context, email string) (*response.UserResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetUserByEmail")
	defer span.End()

	user, err := m.userRepo.GetByEmail(ctx, email)
	if err != nil {
		log.WithFields(log.Fields{
			"email": email,
			"error": err,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetUserByEmail] Failed to get user by email")
		return nil, err
	}

	return &response.UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
	}, nil
}

func (m *Module) GetByID(ctx context.Context) (*response.UserResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.GetByID")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println(reqCtx.GetUserId())
	user, err := m.userRepo.GetById(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"id":    reqCtx.UserId,
			"error": err,
		}).ErrorWithCtx(ctx, "[UserUseCases.GetByID] Failed to get user by id")
		return nil, err
	}

	return &response.UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
	}, nil
}

func (m *Module) UpdateUser(ctx context.Context, req *request.UserUpdateRequest) (*response.UserResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.UpdateUser")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return nil, err
	}

	existing, err := m.userRepo.GetById(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"id":    reqCtx.UserId,
			"error": err,
		}).ErrorWithCtx(ctx, "[UserUseCases.UpdateUser] Failed to get user by id")
		return nil, err
	}

	existing.Username = req.Username

	err = m.userRepo.Update(ctx, existing.ID, &model.UserUpdate{
		Username: existing.Username,
	})

	if err != nil {
		log.WithFields(log.Fields{
			"id":      reqCtx.UserId,
			"error":   err,
			"request": req,
		}).ErrorWithCtx(ctx, "[UserUseCases.UpdateUser] Failed to update user")
		return nil, err
	}

	err = m.publisher.Publish(ctx, m.topics.MasterDataUser.Value, existing.ID.String(), existing.ToMessageModel(), map[string]any{
		constants.MetaDataOperation: constants.Update,
	})

	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:       existing.ID.String(),
		Username: existing.Username,
	}, nil
}

func (m *Module) ChangePassword(ctx context.Context, req *request.UserChangePasswordRequest) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.ChangePassword")
	defer span.End()

	reqCtx, err := libCtx.GetRequestContext(ctx)
	if err != nil {
		return err
	}

	existing, err := m.userRepo.GetById(ctx, reqCtx.GetUserId())
	if err != nil {
		log.WithFields(log.Fields{
			"id":    reqCtx.UserId,
			"error": err,
		}).ErrorWithCtx(ctx, "[UserUseCases.ChangePassword] Failed to get user by id")
		return err
	}

	oldHash := utils.PasswordHash(req.Old, existing.PasswordSalt)
	newPassword := utils.PasswordHash(req.Confirm, existing.PasswordSalt)
	if oldHash != existing.Password {
		return &custerr.ErrChain{
			Message: errorcode.WrongPassword.Message,
			Code:    errorcode.WrongPassword.Code,
		}
	}

	err = m.userRepo.ChangePassword(ctx, existing.ID, newPassword)

	if err != nil {
		log.WithFields(log.Fields{
			"id":      reqCtx.UserId,
			"err":     err,
			"request": req,
		}).ErrorWithCtx(ctx, "[UserUseCases.ChangePassword] Failed to change password")
		return err
	}

	return nil
}

func (m *Module) LoginUser(ctx context.Context, req *request.UserLoginRequest) (*response.UserLoginResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "UserUseCases.LoginUser")
	defer span.End()

	existing, err := m.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		log.WithFields(log.Fields{
			"request": req,
			"error":   err,
		}).ErrorWithCtx(ctx, "[UserUseCases.LoginUser] Failed to get user by email")
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
	token, err := m.jwtSvc.GenerateToken(ctx, jwtsvc.AccessClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: expiresAt,
			ID:        uuid.New().String(),
		},
		JwtData: &jwtsvc.JwtData{
			UserID: existing.ID.String(),
			Role:   existing.Role,
		},
	})

	if err != nil {
		log.WithFields(log.Fields{
			"error":   err,
			"request": req,
		}).ErrorWithCtx(ctx, "[UserUseCases.LoginUser] Failed to generate token")
		return nil, err
	}

	return &response.UserLoginResponse{
		Token:     token,
		ExpiresAt: &expiresAt.Time,
	}, nil

}
