package request

import (
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/ums/pkg/constants"
)

type KPUProvinsiRegistrationRequest struct {
	Name              string `json:"name"`
	Address           string `json:"address"`
	Region            string `json:"region"`
	IsActive          bool   `json:"is_active"`
	SignedTransaction string `json:"signed_transaction"`
}

type KPUProvinsiUpdateRequest struct {
	Name              string `json:"name"`
	Region            string `json:"region"`
	Telephone         string `json:"telephone"`
	SignedTransaction string `json:"signed_transaction"`
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

	if req.SignedTransaction == constants.EmptyString {
		return &custerr.ErrChain{
			Message: "Signed transaction is required",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	return nil
}
