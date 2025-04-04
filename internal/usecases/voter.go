package usecases

import (
	"context"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type VoterUseCases interface {
	RegisterVoter(ctx context.Context, req *request.VoterRegistrationRequest) (*response.VoterRegistrationResponse, error)
	GetVoterByNIK(ctx context.Context, nik string) (*response.VoterResponse, error)
	GetVoterByAddress(ctx context.Context) (*response.VoterResponse, error)
	GetVoterByRegion(ctx context.Context, region string) (*[]response.VoterResponse, error)
	GetAllVoter(ctx context.Context) (*[]response.VoterResponse, error)
}
