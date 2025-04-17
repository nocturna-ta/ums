package request

import (
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/ums/pkg/constants"
	"github.com/nocturna-ta/ums/pkg/constants/errorcode"
	"net/mail"
)

type UserRegistrationRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserUpdateRequest struct {
	Username string `json:"username"`
}

type UserChangePasswordRequest struct {
	Old     string `json:"old"`
	New     string `json:"new"`
	Confirm string `json:"confirm"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserLoginRequest) ValidateLogin() error {
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return &custerr.ErrChain{
			Message: errorcode.InvalidEmail.Message,
			Code:    errorcode.InvalidEmail.Code,
			Type:    response.ErrBadRequest,
		}
	}
	return nil
}

func (r *UserChangePasswordRequest) ValidateChangePassword() error {
	if r.New != r.Confirm {
		return &custerr.ErrChain{
			Message: errorcode.NewPasswordMismatch.Message,
			Code:    errorcode.NewPasswordMismatch.Code,
			Type:    response.ErrBadRequest,
		}
	}
	if len(r.New) < 6 {
		return &custerr.ErrChain{
			Message: errorcode.PasswordTooShort.Message,
			Code:    errorcode.PasswordTooShort.Code,
			Type:    response.ErrBadRequest,
		}
	}
	return nil
}

func (r *UserUpdateRequest) ValidateUpdateUser() error {
	if r.Username == constants.EmptyString {
		return &custerr.ErrChain{
			Message: errorcode.UsernameEmpty.Message,
			Code:    errorcode.UsernameEmpty.Code,
			Type:    response.ErrBadRequest,
		}
	}

	return nil
}

func (r *UserRegistrationRequest) ValidateRegistrationUser() error {
	if r.Username == constants.EmptyString {
		return &custerr.ErrChain{
			Message: errorcode.UsernameEmpty.Message,
			Code:    errorcode.UsernameEmpty.Code,
			Type:    response.ErrBadRequest,
		}
	}

	if r.Role == constants.EmptyString {
		return &custerr.ErrChain{
			Message: errorcode.RoleEmpty.Message,
			Code:    errorcode.RoleEmpty.Code,
			Type:    response.ErrBadRequest,
		}
	}

	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return &custerr.ErrChain{
			Message: errorcode.InvalidEmail.Message,
			Code:    errorcode.InvalidEmail.Code,
			Type:    response.ErrBadRequest,
		}
	}

	if len(r.Password) < 6 {
		return &custerr.ErrChain{
			Message: errorcode.PasswordTooShort.Message,
			Code:    errorcode.PasswordTooShort.Code,
			Type:    response.ErrBadRequest,
		}
	}

	return nil
}
