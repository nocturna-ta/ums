package dao

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nocturna-ta/golib/database/sql"
	ethMocks "github.com/nocturna-ta/golib/ethereum/mocks"
	"github.com/nocturna-ta/golib/txmanager/utils"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/votechain-contract/binding"
	contractMocks "github.com/nocturna-ta/votechain-contract/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

func TestKPUBranchRepository_InsertKPUBranch(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	datastore := &sql.Store{
		Master: &sql.DB{
			DBConnection: sqlxDB,
		},
		Slave: &sql.DB{
			DBConnection: sqlxDB,
		},
	}

	mockClient := ethMocks.NewClient(t)

	mockDb.ExpectBegin()

	sqlxTx, err := sqlxDB.BeginTxx(context.Background(), nil)
	require.NoError(t, err)

	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	tx := types.NewTransaction(0, common.Address{}, big.NewInt(0), 0, big.NewInt(0), nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privKey)
	require.NoError(t, err)
	txBytes, err := rlp.EncodeToBytes(signedTx)
	require.NoError(t, err)
	signedTxStr := "0x" + common.Bytes2Hex(txBytes)

	type args struct {
		ctx               context.Context
		kpuBranch         *model.KPUBranch
		signedTransaction string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn      func()
	}{
		{
			name: "Success",
			args: args{
				ctx:               utils.SetSqlTx(context.Background(), sqlxTx),
				kpuBranch:         &model.KPUBranch{},
				signedTransaction: signedTxStr,
			},
			wantErr: false,
			fn: func() {

				mockDb.ExpectExec("INSERT INTO kpu_branches").WillReturnResult(sqlmock.NewResult(1, 1))
				mockClient.On("GetEthClient").Return(nil).Once()
				mockClient.On("SendTransaction", mock.Anything, mock.AnythingOfType("*types.Transaction")).Return(nil).Once()

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			u := NewKPUBranchRepository(&OptsKPUBranchRepository{
				Client:          mockClient,
				ContractAddress: common.HexToAddress("0x123456789abcdef"),
				DB:              datastore,
			})

			if err := u.InsertKPUBranch(tt.args.ctx, tt.args.kpuBranch, tt.args.signedTransaction); (err != nil) != tt.wantErr {
				t.Errorf("InsertKPUBranch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKPUBranchRepository_GetAllKPUBranch(t *testing.T) {
	db, mockDb, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	datastore := &sql.Store{
		Master: &sql.DB{
			DBConnection: sqlxDB,
		},
		Slave: &sql.DB{
			DBConnection: sqlxDB,
		},
	}

	mockClient := ethMocks.NewClient(t)
	mockContract := contractMocks.NewIVotechain(t)

	mockDb.ExpectBegin()

	_, err := sqlxDB.BeginTxx(context.Background(), nil)
	require.NoError(t, err)

	expected := model.KPUBranch{
		BaseModel: model.BaseModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ID:            uuid.New(),
		Name:          "test",
		BranchAddress: "0x123456789abcdef",
		Region:        "test",
		IsActive:      true,
	}

	contractResponse := []binding.VotechainKPUBranch{{
		Name:          expected.Name,
		BranchAddress: common.HexToAddress(expected.BranchAddress),
		IsActive:      expected.IsActive,
		Region:        expected.Region,
	}}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		fn      func()
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
			fn: func() {
				rows := mockDb.NewRows([]string{
					"id", "name", "branch_address", "region", "is_active", "created_at", "updated_at",
				}).
					AddRow(
						expected.ID, expected.Name, expected.BranchAddress,
						expected.Region, expected.IsActive, expected.CreatedAt, expected.UpdatedAt,
					)
				mockDb.ExpectQuery("SELECT kpu_branches.id, kpu_branches.name, kpu_branches.branch_address, kpu_branches.region, kpu_branches.is_active, kpu_branches.created_at, kpu_branches.updated_at FROM kpu_branches.*").
					WillReturnRows(rows)
				mockContract.On("GetAllKPUBranches", mock.Anything).Return(contractResponse, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			c := KPUBranchRepository{
				client:   mockClient,
				contract: mockContract,
				db:       datastore,
			}

			_, err := c.GetAllKPUBranch(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKPUBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
