package usecases

import (
	"context"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type UserUseCases interface {
	RegisterUser(ctx context.Context, req *request.UserRegistrationRequest) (*response.UserRegistrationResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*response.UserResponse, error)
	GetByID(ctx context.Context, id string) (*response.UserResponse, error)
	UpdateUser(ctx context.Context, id string, req *request.UserUpdateRequest) (*response.UserResponse, error)
	//ChangePassword(ctx context.Context, id string, req *request.UserChangePasswordRequest) error
	//LoginUser(ctx context.Context, req *request.UserLoginRequest) (*response.UserLoginResponse, error)
}
