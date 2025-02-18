package request

import (
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/ums/pkg/utils"
)

type VoterRegistrationRequest struct {
	NIK          string `json:"nik"`
	VoterAddress string `json:"voter_address"`
	Region       string `json:"region"`
	IsRegistered bool   `json:"is_registered"`
}

func (req *VoterRegistrationRequest) ValidateRegisterRequest() error {
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
