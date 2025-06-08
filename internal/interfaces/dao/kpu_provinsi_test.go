package dao

import (
	"context"
	sql2 "database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum/mocks"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestKPUProvinsiRepository_InsertKPUProvinsi(t *testing.T) {
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
	username := "test_kpu_provinsi"
	name := "Test KPU Provinsi"
	address := "0x123456789abcdef"
	region := "Test Provinsi"
	photoPath := "/path/to/photo"
	telephone := "08123456789"
	now := time.Now()

	type args struct {
		ctx context.Context
		kpu *model.KPUProvinsi
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
				kpu: &model.KPUProvinsi{
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
				mockDb.ExpectExec(`INSERT INTO kpu_provinsi`).WillReturnError(&pq.Error{Code: "23505"})
			},
		},
		{
			name: "ShouldError_FailedInsert",
			args: args{
				ctx: context.Background(),
				kpu: &model.KPUProvinsi{
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
				mockDb.ExpectExec(`INSERT INTO kpu_provinsi`).WillReturnError(errors.New("failed"))
			},
		},
		{
			name: "Success",
			args: args{
				ctx: utils.SetSqlTx(context.Background(), sqlxTx),
				kpu: &model.KPUProvinsi{
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
				mockDb.ExpectExec(`INSERT INTO kpu_provinsi`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			k := NewKPUProvinsiRepository(&OptsKPUProvinsiRepository{
				Client: nil,
				DB:     dataStore,
			})
			if err := k.InsertKPUProvinsi(tt.args.ctx, tt.args.kpu); (err != nil) != tt.wantErr {
				t.Errorf("InsertKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUProvinsiRepository_GetAllKPUProvinsi(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "username", "name", "address", "region", "is_active",
		"photo_path", "telephone", "registered_at", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), uuid.New(), "test_kpu_provinsi", "Test KPU Provinsi", "0x123", "Test Provinsi", true,
		"/photo", "081234", now, now, now,
	)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []model.KPUProvinsi
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background()},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM kpu_provinsi`).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_QueryFailed",
			args:    args{ctx: context.Background()},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM kpu_provinsi`).WillReturnError(errors.New("query failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			k := NewKPUProvinsiRepository(&OptsKPUProvinsiRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := k.GetAllKPUProvinsi(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUProvinsiRepository_GetKPUProvinsiByAddress(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "username", "name", "address", "region", "is_active",
		"photo_path", "telephone", "registered_at", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), uuid.New(), "test_kpu_provinsi", "Test KPU Provinsi", "0x123", "Test Provinsi", true,
		"/photo", "081234", now, now, now,
	)

	type args struct {
		ctx     context.Context
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.KPUProvinsi
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), address: "0x123"},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM kpu_provinsi`).WithArgs("0x123").WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), address: "0x456"},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM kpu_provinsi`).WithArgs("0x456").WillReturnError(sql2.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			k := NewKPUProvinsiRepository(&OptsKPUProvinsiRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := k.GetKPUProvinsiByAddress(tt.args.ctx, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUProvinsiRepository_GetKPUProvinsiByID(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	id := uuid.New()
	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "username", "name", "address", "region", "is_active",
		"photo_path", "telephone", "registered_at", "created_at", "updated_at",
	}).AddRow(
		id, uuid.New(), "test_kpu_provinsi", "Test KPU Provinsi", "0x123", "Test Provinsi", true,
		"/photo", "081234", now, now, now,
	)

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *model.KPUProvinsi
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), id: id},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM kpu_provinsi`).WithArgs(id).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), id: uuid.New()},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM kpu_provinsi`).WithArgs(sqlmock.AnyArg()).WillReturnError(sql2.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			k := NewKPUProvinsiRepository(&OptsKPUProvinsiRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := k.GetKPUProvinsiByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUProvinsiRepository_GetKPUProvinsiByUserID(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	userID := uuid.New()
	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "username", "name", "address", "region", "is_active",
		"photo_path", "telephone", "registered_at", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), userID, "test_kpu_provinsi", "Test KPU Provinsi", "0x123", "Test Provinsi", true,
		"/photo", "081234", now, now, now,
	)

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *model.KPUProvinsi
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), userID: userID},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM kpu_provinsi`).WithArgs(userID).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), userID: uuid.New()},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM kpu_provinsi`).WithArgs(sqlmock.AnyArg()).WillReturnError(sql2.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			k := NewKPUProvinsiRepository(&OptsKPUProvinsiRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := k.GetKPUProvinsiByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUProvinsiRepository_SendTxKPUProvinsiBlockchain(t *testing.T) {
	const validTxHex = "0x02f9015282053942843b9aca00843b9db93083042ce6946957afa20f78cd0556d57b5ca5506d0b2c81540280b8e495b79907000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000002432633733663862612d383833342d346433372d386637382d6531333061653231623634330000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000023032000000000000000000000000000000000000000000000000000000000000c001a07b4100fff3bd2f748c0ed83bc9bf3ede2634fa65605a1000198a8444c3d30dd5a0019691ecb8b350b04a3291cb520db89e38cdfa0521148087b379b06497a00525"

	mockClient := mocks.NewClient(t)

	type args struct {
		ctx               context.Context
		signedTransaction string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), signedTransaction: validTxHex},
			want:    "0x123abc456def789",
			wantErr: false,
			fn: func() {
				mockClient.On("SendTransaction",
					mock.Anything,
					mock.Anything,
				).Return("0x123abc456def789", nil).Once()
			},
		},
		{
			name: "ShouldError_SendTransactionFailed",
			args: args{
				ctx:               context.Background(),
				signedTransaction: validTxHex,
			},
			wantErr: true,
			fn: func() {
				mockClient.On("SendTransaction",
					mock.Anything,
					mock.Anything,
				).Return("", errors.New("blockchain error")).Once()
			},
		},
		{
			name:    "ShouldError_InvalidTransaction",
			args:    args{ctx: context.Background(), signedTransaction: "invalid"},
			wantErr: true,
			fn: func() {
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			k := NewKPUKotaRepository(&OptsKPUKotaRepository{
				Client: mockClient,
				DB:     nil,
			})

			got, err := k.SendTxKPUKotaBlockchain(tt.args.ctx, tt.args.signedTransaction)

			if (err != nil) != tt.wantErr {
				t.Errorf("SendTxKPUKotaBlockchain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("SendTxKPUKotaBlockchain() got = %v, want %v", got, tt.want)
			}

			if !tt.wantErr {
				mockClient.AssertExpectations(t)
			}
		})
	}
}

func TestKPUProvinsiRepository_UpdateKPUProvinsi(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	id := uuid.New()
	now := time.Now()
	kpu := &model.KPUProvinsi{
		BaseModel: model.BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
		},
		ID:        id,
		Name:      "Updated Name",
		Region:    "Updated Region",
		Telephone: "081234567890",
		Username:  "updated_username",
	}

	type args struct {
		ctx context.Context
		kpu *model.KPUProvinsi
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), kpu: kpu},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`UPDATE kpu_provinsi SET`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "ShouldError_NoRowsAffected",
			args:    args{ctx: context.Background(), kpu: kpu},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`UPDATE kpu_provinsi SET`).WillReturnResult(sqlmock.NewResult(1, 0))
			},
		},
		{
			name:    "ShouldError_UpdateFailed",
			args:    args{ctx: context.Background(), kpu: kpu},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`UPDATE kpu_provinsi SET`).WillReturnError(errors.New("update failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			k := NewKPUProvinsiRepository(&OptsKPUProvinsiRepository{
				Client: nil,
				DB:     dataStore,
			})
			if err := k.UpdateKPUProvinsi(tt.args.ctx, tt.args.kpu); (err != nil) != tt.wantErr {
				t.Errorf("UpdateKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUProvinsiRepository_UpdateKPUProvinsiPhoto(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	id := uuid.New()
	photoPath := "/new/photo/path.jpg"

	type args struct {
		ctx       context.Context
		id        uuid.UUID
		photoPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), id: id, photoPath: photoPath},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`UPDATE kpu_provinsi SET`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "ShouldError_UpdateFailed",
			args:    args{ctx: context.Background(), id: id, photoPath: photoPath},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`UPDATE kpu_provinsi SET`).WillReturnError(errors.New("update failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			k := NewKPUProvinsiRepository(&OptsKPUProvinsiRepository{
				Client: nil,
				DB:     dataStore,
			})
			if err := k.UpdateKPUProvinsiPhoto(tt.args.ctx, tt.args.id, tt.args.photoPath); (err != nil) != tt.wantErr {
				t.Errorf("UpdateKPUProvinsiPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
