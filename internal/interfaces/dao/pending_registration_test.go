package dao

import (
	"context"
	sql2 "database/sql"
	"encoding/json"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPendingRegistrationRepository_Insert(t *testing.T) {
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
	role := "voter"
	entityData := `{"nik":"1234567890123456","full_name":"Test User"}`
	rawMessage := json.RawMessage(entityData)
	now := time.Now()

	type args struct {
		ctx          context.Context
		registration *model.PendingRegistration
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
				registration: &model.PendingRegistration{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:         id,
					UserID:     userID,
					Role:       role,
					EntityData: rawMessage,
				},
			},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO pending_registrations`).WillReturnError(&pq.Error{Code: "23505"})
			},
		},
		{
			name: "ShouldError_FailedInsert",
			args: args{
				ctx: context.Background(),
				registration: &model.PendingRegistration{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:         id,
					UserID:     userID,
					Role:       role,
					EntityData: rawMessage,
				},
			},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO pending_registrations`).WillReturnError(errors.New("failed"))
			},
		},
		{
			name: "Success",
			args: args{
				ctx: utils.SetSqlTx(context.Background(), sqlxTx),
				registration: &model.PendingRegistration{
					BaseModel: model.BaseModel{
						CreatedAt: now,
						UpdatedAt: now,
					},
					ID:         id,
					UserID:     userID,
					Role:       role,
					EntityData: rawMessage,
				},
			},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`INSERT INTO pending_registrations`).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewPendingRegistrationRepository(&OptsPendingRegistrationRepository{
				DB: dataStore,
			})
			if err := repo.Insert(tt.args.ctx, tt.args.registration); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPendingRegistrationRepository_Delete(t *testing.T) {
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

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), id: id},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`UPDATE pending_registrations SET is_deleted`).WithArgs(sqlmock.AnyArg(), id).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "Success_WithTransaction",
			args:    args{ctx: utils.SetSqlTx(context.Background(), sqlxTx), id: id},
			wantErr: false,
			fn: func() {
				mockDb.ExpectExec(`UPDATE pending_registrations SET is_deleted`).WithArgs(sqlmock.AnyArg(), id).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "ShouldError_DeleteFailed",
			args:    args{ctx: context.Background(), id: id},
			wantErr: true,
			fn: func() {
				mockDb.ExpectExec(`UPDATE pending_registrations SET is_deleted`).WithArgs(sqlmock.AnyArg(), id).WillReturnError(errors.New("delete failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewPendingRegistrationRepository(&OptsPendingRegistrationRepository{
				DB: dataStore,
			})
			if err := repo.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPendingRegistrationRepository_GetByUserID(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	userID := uuid.New()
	id := uuid.New()
	role := "voter"
	entityData := `{"nik":"1234567890123456","full_name":"Test User"}`
	rawMessage := json.RawMessage(entityData)
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "user_id", "role", "entity_data", "created_at", "updated_at",
	}).AddRow(
		id, userID, role, rawMessage, now, now,
	)

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    *model.PendingRegistration
		wantErr bool
		fn      func()
	}{
		{
			name:    "Success",
			args:    args{ctx: context.Background(), userID: userID},
			wantErr: false,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM pending_registrations`).WithArgs(userID).WillReturnRows(rows)
			},
		},
		{
			name:    "ShouldError_NotFound",
			args:    args{ctx: context.Background(), userID: uuid.New()},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM pending_registrations`).WithArgs(sqlmock.AnyArg()).WillReturnError(sql2.ErrNoRows)
			},
		},
		{
			name:    "ShouldError_QueryFailed",
			args:    args{ctx: context.Background(), userID: userID},
			wantErr: true,
			fn: func() {
				mockDb.ExpectQuery(`SELECT .+ FROM pending_registrations`).WithArgs(userID).WillReturnError(errors.New("query failed"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			repo := NewPendingRegistrationRepository(&OptsPendingRegistrationRepository{
				DB: dataStore,
			})
			_, err := repo.GetByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUserID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
