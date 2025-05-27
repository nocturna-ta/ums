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

func TestUserRepository_Insert(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx  context.Context
		user *model.User
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
			repo := &UserRepository{
				db: tt.fields.db,
			}
			if err := repo.Insert(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewUserRepository(t *testing.T) {
	type args struct {
		opts *OptsUserRepository
	}
	tests := []struct {
		name string
		args args
		want repository.UserRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_ChangePassword(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx     context.Context
		id      uuid.UUID
		newPass string
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
			repo := &UserRepository{
				db: tt.fields.db,
			}
			if err := repo.ChangePassword(tt.args.ctx, tt.args.id, tt.args.newPass); (err != nil) != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetByEmail(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserRepository{
				db: tt.fields.db,
			}
			got, err := repo.GetByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetById(t *testing.T) {
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
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserRepository{
				db: tt.fields.db,
			}
			got, err := repo.GetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetPendingVerificationUsers(t *testing.T) {
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
		want    []model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserRepository{
				db: tt.fields.db,
			}
			got, err := repo.GetPendingVerificationUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingVerificationUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPendingVerificationUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetPendingVerificationUsersByRequestedRole(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx           context.Context
		requestedRole string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &UserRepository{
				db: tt.fields.db,
			}
			got, err := repo.GetPendingVerificationUsersByRequestedRole(tt.args.ctx, tt.args.requestedRole)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingVerificationUsersByRequestedRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPendingVerificationUsersByRequestedRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Insert1(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx  context.Context
		user *model.User
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
			repo := &UserRepository{
				db: tt.fields.db,
			}
			if err := repo.Insert(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_UpdateVerificationStatus(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx    context.Context
		id     uuid.UUID
		status string
		role   string
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
			repo := &UserRepository{
				db: tt.fields.db,
			}
			if err := repo.UpdateVerificationStatus(tt.args.ctx, tt.args.id, tt.args.status, tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("UpdateVerificationStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
