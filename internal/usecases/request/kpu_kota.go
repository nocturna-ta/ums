package request

import (
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
)

type KPUKotaRegistrationRequest struct {
	Name              string `json:"name"`
	Address           string `json:"address"`
	Region            string `json:"region"`
	IsActive          bool   `json:"is_active"`
	SignedTransaction string `json:"signed_transaction"`
}

func (req *KPUKotaRegistrationRequest) ValidateRegistrationRequest() error {
	if req == nil {
		return &custerr.ErrChain{
			Message: "Request cannot be nil",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	return nil
}
