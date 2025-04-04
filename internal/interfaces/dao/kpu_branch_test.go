package dao

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/votechain-contract/binding"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKPUBranchRepository_InsertKPUBranch(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db "sqlmock")

	datastore := &sql.Store{
		Master: &sql.DB{
			DBConnection: sqlxDB,
		},
		Slave: &sql.DB{
			DBConnection: sqlxDB,
		},
	}

	mockDb.ExpectBegin()

	sqlxTx, err := sqlxDB.BeginTxx(context.Background(), nil)
	require.NoError(t, err)

	type args struct {
		ctx               context.Context
		kpuBranch         *model.KPUBranch
		signedTransaction string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn func()
	}{
		{

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			u := NewKPUBranchRepository(&OptsKPUBranchRepository{
				Client:          nil,
				ContractAddress: common.Address{},
				DB:              datastore,
			})

			if err := u.InsertKPUBranch(tt.args.ctx, tt.args.kpuBranch, tt.args.signedTransaction); (err != nil) != tt.wantErr {
				t.Errorf("InsertKPUBranch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
