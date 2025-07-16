package usecases

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/http"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type VoterUseCases interface {
	RegisterVoter(ctx context.Context, req *request.VoterRegistrationRequest) (*response.VoterRegistrationResponse, error)
	GetVoterByNIK(ctx context.Context, nik string) (*response.VoterResponse, error)
	GetVoterByAddress(ctx context.Context) (*response.VoterResponse, error)
	GetVoterByRegion(ctx context.Context, region string) (*[]response.VoterResponse, error)
	GetAllVoter(ctx context.Context) (*[]response.VoterResponse, error)
	GetVoterKTPPhoto(ctx context.Context, id uuid.UUID) (*http.File, string, error)
	GetVoterByUserID(ctx context.Context) (*response.VoterResponse, error)
	GetVoterByProvince(ctx context.Context) (*[]response.VoterResponse, error)
	GetVoterByKPUKota(ctx context.Context) (*[]response.VoterResponse, error)
}
