package dao

import (
	"context"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"testing"
)

func TestVoterRepository_InsertVoter(t *testing.T) {
	type fields struct {
		client ethereum.Client
		db     *sql.Store
	}
	type args struct {
		ctx   context.Context
		voter *model.Voter
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
			v := &VoterRepository{
				client: tt.fields.client,
				db:     tt.fields.db,
			}
			if err := v.InsertVoter(tt.args.ctx, tt.args.voter); (err != nil) != tt.wantErr {
				t.Errorf("InsertVoter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
