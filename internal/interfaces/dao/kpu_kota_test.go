package dao

import (
	"context"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"testing"
)

func TestKPUKotaRepository_InsertKPUKota(t *testing.T) {
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
