package dao

import (
	"context"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"reflect"
	"testing"
)

func TestNewUserStatisticRepository(t *testing.T) {
	type args struct {
		opts *OptsUserStatisticRepository
	}
	tests := []struct {
		name string
		args args
		want repository.UserStatisticRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserStatisticRepository(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserStatisticRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetCountDPTByStatus(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx    context.Context
		status string
		region *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetCountDPTByStatus(tt.args.ctx, tt.args.status, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCountDPTByStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCountDPTByStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetDPTNotVoted(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx    context.Context
		region *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetDPTNotVoted(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDPTNotVoted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDPTNotVoted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetDPTTotal(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx    context.Context
		region *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetDPTTotal(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDPTTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDPTTotal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetDPTVoted(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx    context.Context
		region *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetDPTVoted(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDPTVoted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDPTVoted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetDistrictCount(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetDistrictCount(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDistrictCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDistrictCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetKPUKotaApprovedUsers(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetKPUKotaApprovedUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaApprovedUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetKPUKotaApprovedUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetKPUKotaStaff(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx    context.Context
		region *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetKPUKotaStaff(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetKPUKotaStaff() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetKPUProvinsiApprovedUsers(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetKPUProvinsiApprovedUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiApprovedUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetKPUProvinsiApprovedUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetKPUProvinsiStaff(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx    context.Context
		region *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetKPUProvinsiStaff(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetKPUProvinsiStaff() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStatisticRepository_GetProvinceCount(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserStatisticRepository{
				db: tt.fields.db,
			}
			got, err := u.GetProvinceCount(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProvinceCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetProvinceCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
