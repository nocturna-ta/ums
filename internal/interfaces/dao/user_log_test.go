package dao

import (
	"context"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"reflect"
	"testing"
)

func TestNewUserLogRepository(t *testing.T) {
	type args struct {
		opts *OptsUserLogRepository
	}
	tests := []struct {
		name string
		args args
		want repository.UserLogRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserLogRepository(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserLogRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLogRepository_GetUserLogs(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.UserLogs
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserLogRepository{
				db: tt.fields.db,
			}
			got, err := u.GetUserLogs(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserLogs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserLogRepository_InsertLog(t *testing.T) {
	type fields struct {
		db *sql.Store
	}
	type args struct {
		ctx context.Context
		log *model.UserLogs
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
			u := &UserLogRepository{
				db: tt.fields.db,
			}
			if err := u.InsertLog(tt.args.ctx, tt.args.log); (err != nil) != tt.wantErr {
				t.Errorf("InsertLog() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
