package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"reflect"
	"testing"
)

func TestVoterRepository_InsertVoter(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx   context.Context
		voter *model.Voter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := v.InsertVoter(tt.args.ctx, tt.args.voter); (err != nil) != tt.wantErr {
				t.Errorf("InsertVoter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewVoterRepository(t *testing.T) {
	type args struct {
		opts *OptsVoterRepository
	}
	tests := []struct {
		name string
		args args
		want repository.VoterRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVoterRepository(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVoterRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoterRepository_GetAllVoter(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Voter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := v.GetAllVoter(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllVoter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllVoter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoterRepository_GetVoterByAddress(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx     context.Context
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Voter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := v.GetVoterByAddress(tt.args.ctx, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterByAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoterRepository_GetVoterByID(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Voter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := v.GetVoterByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoterRepository_GetVoterByNIK(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx context.Context
		nik string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Voter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := v.GetVoterByNIK(tt.args.ctx, tt.args.nik)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByNIK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterByNIK() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoterRepository_GetVoterByRegion(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Voter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := v.GetVoterByRegion(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByRegion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterByRegion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoterRepository_GetVoterByUserID(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Voter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := v.GetVoterByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoterRepository_InsertVoter1(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx   context.Context
		voter *model.Voter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := v.InsertVoter(tt.args.ctx, tt.args.voter); (err != nil) != tt.wantErr {
				t.Errorf("InsertVoter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVoterRepository_SendTxVoterBlockchain(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx               context.Context
		signedTransaction string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := v.SendTxVoterBlockchain(tt.args.ctx, tt.args.signedTransaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendTxVoterBlockchain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SendTxVoterBlockchain() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoterRepository_UpdateVoter(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx   context.Context
		voter *model.Voter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := v.UpdateVoter(tt.args.ctx, tt.args.voter); (err != nil) != tt.wantErr {
				t.Errorf("UpdateVoter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
