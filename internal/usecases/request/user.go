package request

import (
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/ums/pkg/common"
	"github.com/nocturna-ta/ums/pkg/constants"
	"github.com/nocturna-ta/ums/pkg/constants/errorcode"
	"github.com/nocturna-ta/ums/pkg/roles"
	"io"
	"net/mail"
)

type UserRegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Address  string `json:"address,omitempty"`

	// Role-specific fields
	// For KPU Provinsi and KPU Kota
	KPUName string `json:"kpu-name,omitempty"`
	Region  string `json:"region,omitempty"`

	// For Voter
	NIK                string    `json:"nik,omitempty"`
	FullName           string    `json:"full_name,omitempty"`
	Gender             string    `json:"gender,omitempty"`
	BirthPlace         string    `json:"birth_place,omitempty"`
	BirthDate          string    `json:"birth_date,omitempty"`
	ResidentialAddress string    `json:"residential_address,omitempty"`
	KTPPhotoPath       string    `json:"ktp_photo_path,omitempty"`
	KTPPhotoFile       io.Reader `json:"-" swaggerignore:"true"`
	KTPPhotoName       string    `json:"-" swaggerignore:"true"`
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

type UserVerificationRequest struct {
	UserID            string `json:"user_id"`
	AdminReason       string `json:"admin_reason"`
	SignedTransaction string `json:"signed_transaction"`
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

func (r *UserVerificationRequest) ValidateVerificationRequest() error {
	_, err := uuid.Parse(r.UserID)
	if err != nil {
		return &custerr.ErrChain{
			Message: errorcode.InvalidUUID.Message,
			Code:    errorcode.InvalidUUID.Code,
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

func (r *UserRegistrationRequest) ValidateRegistrationUser() error {
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

	if roles.IsValidRole(r.Role) {
		switch r.Role {
		case roles.RoleKPUProvinsi, roles.RoleKPUKota:
			if r.KPUName == constants.EmptyString {
				return &custerr.ErrChain{
					Message: "Name is required for KPU roles",
					Code:    400,
					Type:    response.ErrBadRequest,
				}
			}
			if r.Address == constants.EmptyString {
				return &custerr.ErrChain{
					Message: "Address is required for KPU roles",
					Code:    400,
					Type:    response.ErrBadRequest,
				}
			}
			if r.Region == constants.EmptyString {
				return &custerr.ErrChain{
					Message: "Region is required for KPU roles",
					Code:    400,
					Type:    response.ErrBadRequest,
				}
			}
		case roles.RoleVoter:
			if r.NIK == constants.EmptyString {
				return &custerr.ErrChain{
					Message: "NIK is required for voter role",
					Code:    400,
					Type:    response.ErrBadRequest,
				}
			}
			if r.Address == constants.EmptyString {
				return &custerr.ErrChain{
					Message: "Voter address is required for voter role",
					Code:    400,
					Type:    response.ErrBadRequest,
				}
			}
			if !common.IsValidNIK(r.NIK) {
				return &custerr.ErrChain{
					Message: "NIK is not valid",
					Code:    400,
					Type:    response.ErrBadRequest,
				}
			}
		}
	}

	return nil
}
