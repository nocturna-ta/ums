package user_statistic

import (
	"context"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/infrastructures/wilayah"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"reflect"
	"testing"
)

func TestModule_GetApprovedDPTStatistic(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.ApprovedDPTResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.GetApprovedDPTStatistic(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApprovedDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApprovedDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetKotaInformationDPTStatistic(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*response.DPTInformationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.GetKotaInformationDPTStatistic(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKotaInformationDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKotaInformationDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetPendingDPTStatistic(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.PendingDPTResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.GetPendingDPTStatistic(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPendingDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetProvinceInformationDPTStatistic(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*response.DPTInformationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.GetProvinceInformationDPTStatistic(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProvinceInformationDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProvinceInformationDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetRejectedDPTStatistic(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.RejectedDPTResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.GetRejectedDPTStatistic(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRejectedDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRejectedDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetStaffKPUKotaStatistic(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.StaffKPUResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.GetStaffKPUKotaStatistic(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaffKPUKotaStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaffKPUKotaStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetStaffKPUProvinceStatistic(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.StaffKPUResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.GetStaffKPUProvinceStatistic(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaffKPUProvinceStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaffKPUProvinceStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetTotalDPTStatistic(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.TotalDPTResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.GetTotalDPTStatistic(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTotalDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTotalDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetVotedStatistic(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.VotedStatisticResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.GetVotedStatistic(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVotedStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVotedStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_getCitiesInProvince(t *testing.T) {
	type fields struct {
		userStatisticRepo repository.UserStatisticRepository
		kpuProvinsiRepo   repository.KPUProvinsiRepository
		kpuKotaRepo       repository.KPUKotaRepository
		wilayahAPIClient  *wilayah.WilayahAPIClient
	}
	type args struct {
		ctx      context.Context
		province string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userStatisticRepo: tt.fields.userStatisticRepo,
				kpuProvinsiRepo:   tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:       tt.fields.kpuKotaRepo,
				wilayahAPIClient:  tt.fields.wilayahAPIClient,
			}
			got, err := m.getCitiesInProvince(tt.args.ctx, tt.args.province)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCitiesInProvince() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCitiesInProvince() got = %v, want %v", got, tt.want)
			}
		})
	}
}
