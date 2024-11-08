package request

import (
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/ums/pkg/constants"
	"github.com/nocturna-ta/ums/pkg/utils"
)

type ChangeUserPasswordRequest struct {
	Old     string `json:"old"`
	New     string `json:"new"`
	Confirm string `json:"confirm"`
}

type UserLoginRequest struct {
	NIK      string `json:"nik"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	NIK  string `json:"nik"`
	Name string `json:"name"`
}

func (req *UserRegisterRequest) ValidateRegisterRequest() error {
	if req == nil {
		return &custerr.ErrChain{
			Message: "Request cannot be nil",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	if !utils.IsValidNIK(req.NIK) {
		return &custerr.ErrChain{
			Message: "NIK is not valid",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	if req.Name == constants.EmptyString {
		return &custerr.ErrChain{
			Message: "Name cannot be empty",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	return nil
}

func (req *UserLoginRequest) ValidateLoginRequest() error {
	if req == nil {
		return &custerr.ErrChain{
			Message: "Request cannot be nil",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	if !utils.IsValidNIK(req.NIK) {
		return &custerr.ErrChain{
			Message: "NIK is not valid",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}
	return nil
}
