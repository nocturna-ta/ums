package repository

import (
	"context"
)

type UserStatisticRepository interface {
	GetCountDPTByStatus(ctx context.Context, status string, region *string) (int, error)
	GetDPTTotal(ctx context.Context, region *string) (int, error)
	GetDPTVoted(ctx context.Context, region *string) (int, error)
	GetDPTNotVoted(ctx context.Context, region *string) (int, error)
	GetKPUProvinsiApprovedUsers(ctx context.Context) (int, error)
	GetKPUProvinsiStaff(ctx context.Context, region *string) (int, error)
	GetKPUKotaApprovedUsers(ctx context.Context) (int, error)
	GetKPUKotaStaff(ctx context.Context, region *string) (int, error)
	GetProvinceCount(ctx context.Context) (int, error)
	GetDistrictCount(ctx context.Context) (int, error)
}
