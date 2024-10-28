package usecases

import (
	"context"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type UserUseCases interface {
	Register(ctx context.Context, req *request.UserRegisterRequest) (*response.UserRegistrationResponse, error)
	Login(ctx context.Context, req *request.UserLoginRequest) (*response.UserLoginResponse, error)
	GetUserByID(ctx context.Context) (*response.UserResponse, error)
	ChangePassword(ctx context.Context, req request.ChangeUserPasswordRequest) (*response.UserResponse, error)
	GetUserByNIK(ctx context.Context, nik string) (*response.UserResponse, error)
}
