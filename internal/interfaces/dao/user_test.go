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
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/pkg/constants"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserRepository_Insert(t *testing.T) {
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
	email := "test@example.com"
	password := "hashedpassword"
	passwordSalt := "salt123"
	role := "voter"
	requestedRole := "voter"
	isActive := true
	verificationStatus := model.VerificationStatusPending
	now := time.Now()

	type args struct {
		ctx  context.Context
		user *model.User
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
				user: &model.User{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:                 id,
					Email:              email,
					Password:           password,
					PasswordSalt:       passwordSalt,
					Role:               role,
					RequestedRole:      requestedRole,
					IsActive:           isActive,
					VerificationStatus: verificationStatus,
				},
			},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO users`).WillReturnError(&pq.Error{Code: "23505"})
			},
		},
		{
			name: "ShouldError_FailedInsert",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:                 id,
					Email:              email,
					Password:           password,
					PasswordSalt:       passwordSalt,
					Role:               role,
					RequestedRole:      requestedRole,
					IsActive:           isActive,
					VerificationStatus: verificationStatus,
				},
			},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO users`).WillReturnError(errors.New("failed"))
			},
		},
		{
			name: "Success",
			args: args{
				ctx: utils.SetSqlTx(context.Background(), sqlxTx),
				user: &model.User{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:                 id,
					Email:              email,
					Password:           password,
					PasswordSalt:       passwordSalt,
					Role:               role,
					RequestedRole:      requestedRole,
					IsActive:           isActive,
					VerificationStatus: verificationStatus,
				},
			},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO users`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewUserRepository(&OptsUserRepository{
				DB: dataStore,
			})
			if err := repo.Insert(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_ChangePassword(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	id := uuid.New()
	newPassword := "newhashedpassword"

	type args struct {
		ctx     context.Context
		id      uuid.UUID
		newPass string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), id: id, newPass: newPassword},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`UPDATE users SET`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "ShouldError_EmptyPassword",
			args:    args{ctx: context.Background(), id: id, newPass: constants.EmptyString},
			wantErr: true,
			fn: func() {
				// No DB mock needed, error occurs before DB call
			},
		},
		{
			name:    "ShouldError_UpdateFailed",
			args:    args{ctx: context.Background(), id: id, newPass: newPassword},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`UPDATE users SET`).WillReturnError(errors.New("update failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewUserRepository(&OptsUserRepository{
				DB: dataStore,
			})
			if err := repo.ChangePassword(tt.args.ctx, tt.args.id, tt.args.newPass); (err != nil) != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	email := "test@example.com"
	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "email", "password", "password_salt", "role", "requested_role",
		"is_active", "verification_status", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), email, "hashedpassword", "salt123", "voter", "voter",
		true, model.VerificationStatusApproved, now, now,
	)

	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), email: email},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WithArgs(email).WillReturnRows(rows)
			},
		},
		{
			name:    "Success_URLEncoded",
			args:    args{ctx: context.Background(), email: "test%40example.com"},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WithArgs("test@example.com").WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), email: "notfound@example.com"},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WithArgs("notfound@example.com").WillReturnError(sql2.ErrNoRows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewUserRepository(&OptsUserRepository{
				DB: dataStore,
			})
			_, err := repo.GetByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetById(t *testing.T) {
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
		"id", "email", "password", "password_salt", "role", "requested_role",
		"is_active", "verification_status", "created_at", "updated_at",
	}).AddRow(
		id, "test@example.com", "hashedpassword", "salt123", "voter", "voter",
		true, model.VerificationStatusApproved, now, now,
	)

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), id: id},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WithArgs(id).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), id: uuid.New()},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WithArgs(sqlmock.AnyArg()).WillReturnError(sql2.ErrNoRows)
			},
		},
		{
			name:    "ShouldError_QueryFailed",
			args:    args{ctx: context.Background(), id: id},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WithArgs(id).WillReturnError(errors.New("query failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewUserRepository(&OptsUserRepository{
				DB: dataStore,
			})
			_, err := repo.GetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetPendingVerificationUsers(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "email", "password", "password_salt", "role", "requested_role",
		"is_active", "verification_status", "created_at", "updated_at",
	}).
		AddRow(uuid.New(), "user1@example.com", "hash1", "salt1", "voter", "voter", false, model.VerificationStatusPending, now, now).
		AddRow(uuid.New(), "user2@example.com", "hash2", "salt2", "kpu_kota", "kpu_kota", false, model.VerificationStatusPending, now, now)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []model.User
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background()},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_QueryFailed",
			args:    args{ctx: context.Background()},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WillReturnError(errors.New("query failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewUserRepository(&OptsUserRepository{
				DB: dataStore,
			})
			_, err := repo.GetPendingVerificationUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingVerificationUsers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetPendingVerificationUsersByRequestedRole(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	requestedRole := "voter"
	now := time.Now()
	rows := sqlmock.NewRows([]string{
		"id", "email", "password", "password_salt", "role", "requested_role",
		"is_active", "verification_status", "created_at", "updated_at",
	}).AddRow(
		uuid.New(), "voter@example.com", "hash", "salt", "voter", requestedRole,
		false, model.VerificationStatusPending, now, now,
	)

	type args struct {
		ctx           context.Context
		requestedRole string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.User
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), requestedRole: requestedRole},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WithArgs(model.VerificationStatusPending, requestedRole).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_QueryFailed",
			args:    args{ctx: context.Background(), requestedRole: requestedRole},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM users`).WithArgs(model.VerificationStatusPending, requestedRole).WillReturnError(errors.New("query failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewUserRepository(&OptsUserRepository{
				DB: dataStore,
			})
			_, err := repo.GetPendingVerificationUsersByRequestedRole(tt.args.ctx, tt.args.requestedRole)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingVerificationUsersByRequestedRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_UpdateVerificationStatus(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	id := uuid.New()
	status := model.VerificationStatusApproved
	role := "voter"

	type args struct {
		ctx    context.Context
		id     uuid.UUID
		status string
		role   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), id: id, status: status, role: role},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`UPDATE users SET`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "ShouldError_UpdateFailed",
			args:    args{ctx: context.Background(), id: id, status: status, role: role},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`UPDATE users SET`).WillReturnError(errors.New("update failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewUserRepository(&OptsUserRepository{
				DB: dataStore,
			})
			if err := repo.UpdateVerificationStatus(tt.args.ctx, tt.args.id, tt.args.status, tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("UpdateVerificationStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
