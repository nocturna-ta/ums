package request

import (
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
)

type KPUProvinsiRegistrationRequest struct {
	Name              string `json:"name"`
	Address           string `json:"address"`
	Region            string `json:"region"`
	IsActive          bool   `json:"is_active"`
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
