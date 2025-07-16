package usecases

import (
	"context"
	"github.com/nocturna-ta/ums/internal/usecases/response"
)

type UserStatisticUseCases interface {
	GetApprovedDPTStatistic(ctx context.Context, region string) (*response.ApprovedDPTResponse, error)
	GetRejectedDPTStatistic(ctx context.Context, region string) (*response.RejectedDPTResponse, error)
	GetPendingDPTStatistic(ctx context.Context, region string) (*response.PendingDPTResponse, error)
	GetTotalDPTStatistic(ctx context.Context, region string) (*response.TotalDPTResponse, error)
	GetStaffKPUProvinceStatistic(ctx context.Context, region string) (*response.StaffKPUResponse, error)
	GetStaffKPUKotaStatistic(ctx context.Context, region string) (*response.StaffKPUResponse, error)
	GetProvinceInformationDPTStatistic(ctx context.Context) (*[]response.DPTInformationResponse, error)
	GetKotaInformationDPTStatistic(ctx context.Context) (*[]response.DPTInformationResponse, error)
	GetVotedStatistic(ctx context.Context, region string) (*response.VotedStatisticResponse, error)
}
