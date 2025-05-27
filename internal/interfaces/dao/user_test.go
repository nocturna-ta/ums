package dao

import (
	"context"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/ums/internal/domain/model"
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
