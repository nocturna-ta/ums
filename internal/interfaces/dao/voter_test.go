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

func TestVoterRepository_InsertVoter(t *testing.T) {
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
	nik := "1234567890123456"
	fullName := "Test Voter"
	gender := "Male"
	birthPlace := "Jakarta"
	birthDate := time.Now().AddDate(-25, 0, 0)
	residentialAddress := "Jl. Test No. 123"
	region := "Jakarta Selatan"
	voterAddress := "0x123456789abcdef"
	ktpPhotoPath := "/path/to/ktp.jpg"
	now := time.Now()

	type args struct {
		ctx   context.Context
		voter *model.Voter
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
				voter: &model.Voter{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:                 id,
					UserID:             userID,
					NIK:                nik,
					FullName:           fullName,
					Gender:             gender,
					BirthPlace:         birthPlace,
					BirthDate:          birthDate,
					ResidentialAddress: residentialAddress,
					Region:             region,
					VoterAddress:       voterAddress,
					IsRegistered:       true,
					KTPPhotoPath:       ktpPhotoPath,
					HasVoted:           false,
					VotedAt:            nil,
					LastLogin:          nil,
				},
			},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO voters`).WillReturnError(&pq.Error{Code: "23505"})
			},
		},
		{
			name: "ShouldError_FailedInsert",
			args: args{
				ctx: context.Background(),
				voter: &model.Voter{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:                 id,
					UserID:             userID,
					NIK:                nik,
					FullName:           fullName,
					Gender:             gender,
					BirthPlace:         birthPlace,
					BirthDate:          birthDate,
					ResidentialAddress: residentialAddress,
					Region:             region,
					VoterAddress:       voterAddress,
					IsRegistered:       true,
					KTPPhotoPath:       ktpPhotoPath,
					HasVoted:           false,
					VotedAt:            nil,
					LastLogin:          nil,
				},
			},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO voters`).WillReturnError(errors.New("failed"))
			},
		},
		{
			name: "Success",
			args: args{
				ctx: utils.SetSqlTx(context.Background(), sqlxTx),
				voter: &model.Voter{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:                 id,
					UserID:             userID,
					NIK:                nik,
					FullName:           fullName,
					Gender:             gender,
					BirthPlace:         birthPlace,
					BirthDate:          birthDate,
					ResidentialAddress: residentialAddress,
					Region:             region,
					VoterAddress:       voterAddress,
					IsRegistered:       true,
					KTPPhotoPath:       ktpPhotoPath,
					HasVoted:           false,
					VotedAt:            nil,
					LastLogin:          nil,
				},
			},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO voters`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			v := NewVoterRepository(&OptsVoterRepository{
				Client: nil,
				DB:     dataStore,
			})
			if err := v.InsertVoter(tt.args.ctx, tt.args.voter); (err != nil) != tt.wantErr {
				t.Errorf("InsertVoter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVoterRepository_GetAllVoter(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	now := time.Now()
	birthDate := time.Now().AddDate(-25, 0, 0)
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "nik", "full_name", "gender", "birth_place",
		"birth_date", "residential_address", "region", "voter_address", "is_registered", "ktp_photo_path",
		"has_voted", "voted_at", "last_login", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), uuid.New(), "1234567890123456", "Test Voter", "Male", "Jakarta",
		birthDate, "Jl. Test", "Jakarta", "0x123", true, "/ktp.jpg",
		false, nil, nil, now, now,
	)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Voter
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background()},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_QueryFailed",
			args:    args{ctx: context.Background()},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WillReturnError(errors.New("query failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			v := NewVoterRepository(&OptsVoterRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := v.GetAllVoter(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllVoter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVoterRepository_GetVoterByAddress(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	now := time.Now()
	birthDate := time.Now().AddDate(-25, 0, 0)
	address := "0x123456789abcdef"
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "nik", "full_name", "gender", "birth_place",
		"birth_date", "residential_address", "region", "voter_address", "is_registered", "ktp_photo_path",
		"has_voted", "voted_at", "last_login", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), uuid.New(), "1234567890123456", "Test Voter", "Male", "Jakarta",
		birthDate, "Jl. Test", "Jakarta", address, true, "/ktp.jpg",
		false, nil, nil, now, now,
	)

	type args struct {
		ctx     context.Context
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Voter
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), address: address},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs(address).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), address: "0x456"},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs("0x456").WillReturnError(sql2.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			v := NewVoterRepository(&OptsVoterRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := v.GetVoterByAddress(tt.args.ctx, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVoterRepository_GetVoterByID(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	id := uuid.New()
	now := time.Now()
	birthDate := time.Now().AddDate(-25, 0, 0)
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "nik", "full_name", "gender", "birth_place",
		"birth_date", "residential_address", "region", "voter_address", "is_registered", "ktp_photo_path",
		"has_voted", "voted_at", "last_login", "created_at", "updated_at",
	}).AddRow(
		id, uuid.New(), "1234567890123456", "Test Voter", "Male", "Jakarta",
		birthDate, "Jl. Test", "Jakarta", "0x123", true, "/ktp.jpg",
		false, nil, nil, now, now,
	)

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Voter
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), id: id},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs(id).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), id: uuid.New()},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs(sqlmock.AnyArg()).WillReturnError(sql2.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			v := NewVoterRepository(&OptsVoterRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := v.GetVoterByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVoterRepository_GetVoterByNIK(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	now := time.Now()
	birthDate := time.Now().AddDate(-25, 0, 0)
	nik := "1234567890123456"
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "nik", "full_name", "gender", "birth_place",
		"birth_date", "residential_address", "region", "voter_address", "is_registered", "ktp_photo_path",
		"has_voted", "voted_at", "last_login", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), uuid.New(), nik, "Test Voter", "Male", "Jakarta",
		birthDate, "Jl. Test", "Jakarta", "0x123", true, "/ktp.jpg",
		false, nil, nil, now, now,
	)

	type args struct {
		ctx context.Context
		nik string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Voter
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), nik: nik},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs(nik).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), nik: "9999999999999999"},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs("9999999999999999").WillReturnError(sql2.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			v := NewVoterRepository(&OptsVoterRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := v.GetVoterByNIK(tt.args.ctx, tt.args.nik)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByNIK() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVoterRepository_GetVoterByRegion(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	now := time.Now()
	birthDate := time.Now().AddDate(-25, 0, 0)
	region := "Jakarta Selatan"
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "nik", "full_name", "gender", "birth_place",
		"birth_date", "residential_address", "region", "voter_address", "is_registered", "ktp_photo_path",
		"has_voted", "voted_at", "last_login", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), uuid.New(), "1234567890123456", "Test Voter", "Male", "Jakarta",
		birthDate, "Jl. Test", region, "0x123", true, "/ktp.jpg",
		false, nil, nil, now, now,
	)

	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Voter
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), region: region},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs(region).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_QueryFailed",
			args:    args{ctx: context.Background(), region: "Unknown Region"},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs("Unknown Region").WillReturnError(errors.New("query failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			v := NewVoterRepository(&OptsVoterRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := v.GetVoterByRegion(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByRegion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVoterRepository_GetVoterByUserID(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	userID := uuid.New()
	now := time.Now()
	birthDate := time.Now().AddDate(-25, 0, 0)
	rows := sqlmock.NewRows([]string{
		"id", "user_id", "nik", "full_name", "gender", "birth_place",
		"birth_date", "residential_address", "region", "voter_address", "is_registered", "ktp_photo_path",
		"has_voted", "voted_at", "last_login", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), userID, "1234567890123456", "Test Voter", "Male", "Jakarta",
		birthDate, "Jl. Test", "Jakarta", "0x123", true, "/ktp.jpg",
		false, nil, nil, now, now,
	)

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Voter
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), userID: userID},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs(userID).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), userID: uuid.New()},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM voters`).WithArgs(sqlmock.AnyArg()).WillReturnError(sql2.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			v := NewVoterRepository(&OptsVoterRepository{
				Client: nil,
				DB:     dataStore,
			})
			_, err := v.GetVoterByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVoterRepository_SendTxVoterBlockchain(t *testing.T) {
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

func TestVoterRepository_UpdateVoter(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	id := uuid.New()
	now := time.Now()
	voter := &model.Voter{
		BaseModel: model.BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
		},
		ID:       id,
		HasVoted: true,
		VotedAt:  &now,
	}

	type args struct {
		ctx   context.Context
		voter *model.Voter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), voter: voter},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`UPDATE voters SET`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "ShouldError_NoRowsAffected",
			args:    args{ctx: context.Background(), voter: voter},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`UPDATE voters SET`).WillReturnResult(sqlmock.NewResult(1, 0))
			},
		},
		{
			name:    "ShouldError_UpdateFailed",
			args:    args{ctx: context.Background(), voter: voter},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`UPDATE voters SET`).WillReturnError(errors.New("update failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			v := NewVoterRepository(&OptsVoterRepository{
				Client: nil,
				DB:     dataStore,
			})
			if err := v.UpdateVoter(tt.args.ctx, tt.args.voter); (err != nil) != tt.wantErr {
				t.Errorf("UpdateVoter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
