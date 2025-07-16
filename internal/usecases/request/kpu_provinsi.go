package request

import (
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/ums/pkg/constants"
)

type KPUProvinsiRegistrationRequest struct {
	Username          string `json:"username"`
	Name              string `json:"name"`
	Address           string `json:"address"`
	Region            string `json:"region"`
	IsActive          bool   `json:"is_active"`
	SignedTransaction string `json:"signed_transaction"`
}

type KPUProvinsiUpdateRequest struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Region    string `json:"region"`
	Telephone string `json:"telephone"`
}

func (req *KPUProvinsiRegistrationRequest) ValidateRegistrationRequest() error {
	if req == nil {
		return &custerr.ErrChain{
			Message: "Request cannot be nil",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	return nil
}

func (req *KPUProvinsiUpdateRequest) ValidateUpdateRequest() error {
	if req == nil {
		return &custerr.ErrChain{
			Message: "Request cannot be nil",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	if req.Name == constants.EmptyString || req.Region == constants.EmptyString || req.Telephone == constants.EmptyString {
		return &custerr.ErrChain{
			Message: "Name and region cannot be empty",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	return nil
}
