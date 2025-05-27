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

func TestKPUProvinsiRepository_InsertKPUProvinsi(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx context.Context
		kpu *model.KPUProvinsi
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
			K := &KPUProvinsiRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := K.InsertKPUProvinsi(tt.args.ctx, tt.args.kpu); (err != nil) != tt.wantErr {
				t.Errorf("InsertKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUProvinsiRepository_GetAllKPUProvinsi(t *testing.T) {
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
		want    []model.KPUProvinsi
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			K := &KPUProvinsiRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.GetAllKPUProvinsi(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllKPUProvinsi() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUProvinsiRepository_GetKPUProvinsiByAddress(t *testing.T) {
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
		want    *model.KPUProvinsi
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			K := &KPUProvinsiRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.GetKPUProvinsiByAddress(tt.args.ctx, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiByAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUProvinsiRepository_GetKPUProvinsiByID(t *testing.T) {
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
		want    *model.KPUProvinsi
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			K := &KPUProvinsiRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.GetKPUProvinsiByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUProvinsiRepository_GetKPUProvinsiByUserID(t *testing.T) {
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
		want    *model.KPUProvinsi
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			K := &KPUProvinsiRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.GetKPUProvinsiByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUProvinsiRepository_InsertKPUProvinsi1(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx context.Context
		kpu *model.KPUProvinsi
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
			K := &KPUProvinsiRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := K.InsertKPUProvinsi(tt.args.ctx, tt.args.kpu); (err != nil) != tt.wantErr {
				t.Errorf("InsertKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUProvinsiRepository_SendTxKPUProvinsiBlockchain(t *testing.T) {
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
			K := &KPUProvinsiRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.SendTxKPUProvinsiBlockchain(tt.args.ctx, tt.args.signedTransaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendTxKPUProvinsiBlockchain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SendTxKPUProvinsiBlockchain() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUProvinsiRepository_UpdateKPUProvinsi(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx context.Context
		kpu *model.KPUProvinsi
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
			K := &KPUProvinsiRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := K.UpdateKPUProvinsi(tt.args.ctx, tt.args.kpu); (err != nil) != tt.wantErr {
				t.Errorf("UpdateKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUProvinsiRepository_UpdateKPUProvinsiPhoto(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx       context.Context
		id        uuid.UUID
		photoPath string
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
			K := &KPUProvinsiRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := K.UpdateKPUProvinsiPhoto(tt.args.ctx, tt.args.id, tt.args.photoPath); (err != nil) != tt.wantErr {
				t.Errorf("UpdateKPUProvinsiPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewKPUProvinsiRepository(t *testing.T) {
	type args struct {
		opts *OptsKPUProvinsiRepository
	}
	tests := []struct {
		name string
		args args
		want repository.KPUProvinsiRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKPUProvinsiRepository(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKPUProvinsiRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
