package dao

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"reflect"
	"testing"
)

func TestPendingRegistrationRepository_Insert(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx          context.Context
		registration *model.PendingRegistration
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
			repo := &PendingRegistrationRepository{
				db: tt.fields.db,
			}
			if err := repo.Insert(tt.args.ctx, tt.args.registration); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewPendingRegistrationRepository(t *testing.T) {
	type args struct {
		opts *OptsPendingRegistrationRepository
	}
	tests := []struct {
		name string
		args args
		want repository.PendingRegistrationRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPendingRegistrationRepository(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPendingRegistrationRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPendingRegistrationRepository_Delete(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
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
			repo := &PendingRegistrationRepository{
				db: tt.fields.db,
			}
			if err := repo.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPendingRegistrationRepository_GetByUserID(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.PendingRegistration
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &PendingRegistrationRepository{
				db: tt.fields.db,
			}
			got, err := repo.GetByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPendingRegistrationRepository_Insert1(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx          context.Context
		registration *model.PendingRegistration
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
			repo := &PendingRegistrationRepository{
				db: tt.fields.db,
			}
			if err := repo.Insert(tt.args.ctx, tt.args.registration); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
