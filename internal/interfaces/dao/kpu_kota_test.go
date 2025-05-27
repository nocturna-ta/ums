package dao

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

func TestKPUKotaRepository_InsertKPUKota(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	mockDb.ExpectBegin()
	sqlxTx, err := sqlxDB.BeginTxx(context.Background(), nil)
	require.NoError(t, err)

	id := uuid.New()
	userID := uuid.New()
	username := "test_kpu"
	name := "Test KPU Kota"
	address := "0x123456789"
	region := "Test Region"
	photoPath := "/path/to/photo"
	telephone := "08123456789"
	now := time.Now()

	type args struct {
		ctx context.Context
		kpu *model.KPUKota
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn      func()
	}{
		{
			name: "ShouldError_Duplicate",
			args: args{
				ctx: context.Background(),
				kpu: &model.KPUKota{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:           id,
					UserID:       userID,
					Username:     username,
					Name:         name,
					Address:      address,
					Region:       region,
					IsActive:     true,
					PhotoPath:    photoPath,
					Telephone:    telephone,
					RegisteredAt: now,
				},
			},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO kpu_kota`).WillReturnError(&pq.Error{Code: "23505"})
			},
		},
		{
			name: "ShouldError_FailedInsert",
			args: args{
				ctx: context.Background(),
				kpu: &model.KPUKota{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:           id,
					UserID:       userID,
					Username:     username,
					Name:         name,
					Address:      address,
					Region:       region,
					IsActive:     true,
					PhotoPath:    photoPath,
					Telephone:    telephone,
					RegisteredAt: now,
				},
			},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO kpu_kota`).WillReturnError(errors.New("failed"))
			},
		},
		{
			name: "Success",
			args: args{
				ctx: utils.SetSqlTx(context.Background(), sqlxTx),
				kpu: &model.KPUKota{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:           id,
					UserID:       userID,
					Username:     username,
					Name:         name,
					Address:      address,
					Region:       region,
					IsActive:     true,
					PhotoPath:    photoPath,
					Telephone:    telephone,
					RegisteredAt: now,
				},
			},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO kpu_kota`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			k := NewKPUKotaRepository(&OptsKPUKotaRepository{
				Client: nil,
				DB:     dataStore,
			})
			if err := k.InsertKPUKota(tt.args.ctx, tt.args.kpu); (err != nil) != tt.wantErr {
				t.Errorf("InsertKPUKota() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUKotaRepository_GetAllKPUKota(t *testing.T) {
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
		want    []model.KPUKota
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			K := &KPUKotaRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.GetAllKPUKota(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKPUKota() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllKPUKota() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUKotaRepository_GetKPUKotaByAddress(t *testing.T) {
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
		want    *model.KPUKota
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			K := &KPUKotaRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.GetKPUKotaByAddress(tt.args.ctx, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUKotaByAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUKotaRepository_GetKPUKotaByID(t *testing.T) {
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
		want    *model.KPUKota
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			K := &KPUKotaRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.GetKPUKotaByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUKotaByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUKotaRepository_GetKPUKotaByUserID(t *testing.T) {
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
		want    *model.KPUKota
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			K := &KPUKotaRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.GetKPUKotaByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUKotaByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUKotaRepository_InsertKPUKota1(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx context.Context
		kpu *model.KPUKota
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
			K := &KPUKotaRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := K.InsertKPUKota(tt.args.ctx, tt.args.kpu); (err != nil) != tt.wantErr {
				t.Errorf("InsertKPUKota() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUKotaRepository_SendTxKPUKotaBlockchain(t *testing.T) {
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
			K := &KPUKotaRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			got, err := K.SendTxKPUKotaBlockchain(tt.args.ctx, tt.args.signedTransaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendTxKPUKotaBlockchain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SendTxKPUKotaBlockchain() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKPUKotaRepository_UpdateKPUKota(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx context.Context
		kpu *model.KPUKota
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
			K := &KPUKotaRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := K.UpdateKPUKota(tt.args.ctx, tt.args.kpu); (err != nil) != tt.wantErr {
				t.Errorf("UpdateKPUKota() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUKotaRepository_UpdateKPUKotaPhoto(t *testing.T) {
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
			K := &KPUKotaRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := K.UpdateKPUKotaPhoto(tt.args.ctx, tt.args.id, tt.args.photoPath); (err != nil) != tt.wantErr {
				t.Errorf("UpdateKPUKotaPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewKPUKotaRepository(t *testing.T) {
	type args struct {
		opts *OptsKPUKotaRepository
	}
	tests := []struct {
		name string
		args args
		want repository.KPUKotaRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKPUKotaRepository(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKPUKotaRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
