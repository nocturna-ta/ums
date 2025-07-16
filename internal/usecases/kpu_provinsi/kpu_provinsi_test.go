package kpu_provinsi

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/database/sql"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/ethereum/mocks"
	"github.com/nocturna-ta/golib/event"
	mocksEvent "github.com/nocturna-ta/golib/event/mocks"
	"github.com/nocturna-ta/golib/http"
	"github.com/nocturna-ta/golib/txmanager"
	txSql "github.com/nocturna-ta/golib/txmanager/sql"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/domain/repository/mocks_repository"
	"github.com/nocturna-ta/ums/internal/interfaces/dao"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/interfaces/mocks_interfaces"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/votechain-contract/binding/kpuManager"
	_ "github.com/nocturna-ta/votechain-contract/binding/kpuManager"
	"github.com/nocturna-ta/votechain-contract/interfaces"
	mocksContract "github.com/nocturna-ta/votechain-contract/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"reflect"
	"testing"
	"time"
)

func TestModule_GetAllKPUProvinsi(t *testing.T) {
	mockKPUProvinsiRepo := &mocks_repository.KPUProvinsiRepository{}
	mockKPUContract := &mocksContract.KpuManagerInterface{}

	kpuID1 := uuid.New()
	kpuID2 := uuid.New()
	userID1 := uuid.New()
	userID2 := uuid.New()

	registeredAt := time.Now()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		fn       func()
		assertFn func(t *testing.T, got *[]response.KPUProvinsiResponse)
	}{
		{
			name: "ShouldError_FailedGetFromDB",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			fn: func() {
				mockKPUProvinsiRepo.Mock.On("GetAllKPUProvinsi", mock.Anything).Return(nil, errors.New("database error")).Once()
			},
			assertFn: func(t *testing.T, got *[]response.KPUProvinsiResponse) {
				mockKPUProvinsiRepo.AssertExpectations(t)
			},
		},
		{
			name: "ShouldError_FailedGetFromContract",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			fn: func() {
				kpuProvinsis := []model.KPUProvinsi{
					{
						ID:      kpuID1,
						UserID:  userID1,
						Address: "0x123",
					},
				}
				mockKPUProvinsiRepo.Mock.On("GetAllKPUProvinsi", mock.Anything).Return(kpuProvinsis, nil).Once()
				mockKPUContract.Mock.On("GetAllKPUProvinsi", mock.Anything).Return(nil, errors.New("contract error")).Once()
			},
			assertFn: func(t *testing.T, got *[]response.KPUProvinsiResponse) {
				mockKPUProvinsiRepo.AssertExpectations(t)
				mockKPUContract.AssertExpectations(t)
			},
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
			fn: func() {
				addr1 := common.HexToAddress("0x123")
				addr2 := common.HexToAddress("0x456")

				kpuProvinsis := []model.KPUProvinsi{
					{
						ID:           kpuID1,
						UserID:       userID1,
						Username:     "kpu_user1",
						Name:         "KPU Provinsi Jakarta",
						Address:      addr1.String(),
						Region:       "Jakarta",
						IsActive:     true,
						PhotoPath:    "/path/to/photo1.jpg",
						Telephone:    "021-123456",
						RegisteredAt: registeredAt,
					},
					{
						ID:           kpuID2,
						UserID:       userID2,
						Username:     "kpu_user2",
						Name:         "KPU Provinsi Bandung",
						Address:      addr2.String(),
						Region:       "Bandung",
						IsActive:     true,
						PhotoPath:    "",
						Telephone:    "022-123456",
						RegisteredAt: registeredAt,
					},
				}
				contractKPUs := []kpuManager.IKPUManagerKPUProvinsi{
					{
						Name:    "KPU Provinsi Jakarta",
						Address: addr1,
						Region:  "Jakarta",
					},
					{
						Name:    "KPU Provinsi Bandung",
						Address: addr2,
						Region:  "Bandung",
					},
				}
				mockKPUProvinsiRepo.Mock.On("GetAllKPUProvinsi", mock.Anything).Return(kpuProvinsis, nil).Once()
				mockKPUContract.Mock.On("GetAllKPUProvinsi", mock.Anything).Return(contractKPUs, nil).Once()
			},
			assertFn: func(t *testing.T, got *[]response.KPUProvinsiResponse) {
				require.NotNil(t, got)
				require.Len(t, *got, 2)

				require.Equal(t, kpuID1.String(), (*got)[0].ID)
				require.Equal(t, "KPU Provinsi Jakarta", (*got)[0].Name)
				require.Contains(t, (*got)[0].PhotoURL, "/v1/kpu-provinsi/")

				require.Equal(t, kpuID2.String(), (*got)[1].ID)
				require.Equal(t, "KPU Provinsi Bandung", (*got)[1].Name)
				require.Equal(t, "", (*got)[1].PhotoURL)

				mockKPUProvinsiRepo.AssertExpectations(t)
				mockKPUContract.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			m := &Module{
				kpuProvinsiRepo: mockKPUProvinsiRepo,
				kpuContract:     mockKPUContract,
			}

			got, err := m.GetAllKPUProvinsi(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.assertFn(t, got)
		})
	}
}

func TestModule_GetKPUProvinsiByAddress(t *testing.T) {
	mockKPUProvinsiRepo := &mocks_repository.KPUProvinsiRepository{}
	mockKPUContract := &mocksContract.KpuManagerInterface{}

	kpuID := uuid.New()
	userID := uuid.New()
	address := "0x123"
	registeredAt := time.Now()

	reqCtx := libCtx.RequestContext{
		Address: address,
	}
	ctx := context.WithValue(context.Background(), libCtx.RequestContextKey, reqCtx)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		args     args
		want     *response.KPUProvinsiResponse
		wantErr  bool
		fn       func()
		assertFn func(t *testing.T)
	}{
		{
			name: "ShouldError_KPUProvinsiNotFoundInDB",
			args: args{
				ctx: ctx,
			},
			want:    nil,
			wantErr: true,
			fn: func() {
				mockKPUProvinsiRepo.Mock.On("GetKPUProvinsiByAddress", mock.Anything, address).Return(nil, errors.New("kpu provinsi not found")).Once()
			},
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
			},
		},
		{
			name: "ShouldError_KPUProvinsiNotFoundInContract",
			args: args{
				ctx: ctx,
			},
			want:    nil,
			wantErr: true,
			fn: func() {
				kpuProvinsi := &model.KPUProvinsi{
					ID:      kpuID,
					UserID:  userID,
					Address: address,
				}
				mockKPUProvinsiRepo.Mock.On("GetKPUProvinsiByAddress", mock.Anything, reqCtx.GetAddress()).Return(kpuProvinsi, nil).Once()
				mockKPUContract.Mock.On("GetKpuProvinsiByAddress", mock.Anything, common.HexToAddress(reqCtx.GetAddress())).Return(kpuManager.IKPUManagerKPUProvinsi{}, errors.New("kpu provinsi not found in contract")).Once()
			},
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
				mockKPUContract.AssertExpectations(t)
			},
		},
		{
			name: "ShouldError_AddressMismatch",
			args: args{
				ctx: ctx,
			},
			want:    nil,
			wantErr: true,
			fn: func() {
				kpuProvinsi := &model.KPUProvinsi{
					ID:      kpuID,
					UserID:  userID,
					Address: address,
				}

				// Mock contract KPU Provinsi - using actual contract binding type
				contractKPU := kpuManager.IKPUManagerKPUProvinsi{
					Address: common.HexToAddress("0x456"),
					Name:    "Different KPU",
					Region:  "Different Region",
				}

				mockKPUProvinsiRepo.Mock.On("GetKPUProvinsiByAddress", mock.Anything, address).Return(kpuProvinsi, nil).Once()
				mockKPUContract.Mock.On("GetKpuProvinsiByAddress", mock.Anything, common.HexToAddress(address)).Return(contractKPU, nil).Once()
			},
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
				mockKPUContract.AssertExpectations(t)
			},
		},
		{
			name: "Success",
			args: args{
				ctx: ctx,
			},
			want: &response.KPUProvinsiResponse{
				ID:           kpuID.String(),
				UserID:       userID.String(),
				Username:     "kpu_user",
				Name:         "KPU Provinsi Jakarta",
				Address:      common.HexToAddress(address).String(), // Use formatted address
				Region:       "Jakarta",
				IsActive:     true,
				PhotoURL:     "/v1/kpu-provinsi/" + kpuID.String() + "/photo",
				Telephone:    "021-123456",
				RegisteredAt: registeredAt.String(),
			},
			wantErr: false,
			fn: func() {
				// Ensure address format matches
				addr := common.HexToAddress(address)

				kpuProvinsi := &model.KPUProvinsi{
					ID:           kpuID,
					UserID:       userID,
					Username:     "kpu_user",
					Name:         "KPU Provinsi Jakarta",
					Address:      addr.String(), // Use the same string format as contract
					Region:       "Jakarta",
					IsActive:     true,
					PhotoPath:    "/path/to/photo.jpg",
					Telephone:    "021-123456",
					RegisteredAt: registeredAt,
				}

				// Mock contract KPU Provinsi - using actual contract binding type
				contractKPU := kpuManager.IKPUManagerKPUProvinsi{
					Address: addr,
					Name:    "KPU Provinsi Jakarta",
					Region:  "Jakarta",
				}

				mockKPUProvinsiRepo.Mock.On("GetKPUProvinsiByAddress", mock.Anything, address).Return(kpuProvinsi, nil).Once()
				mockKPUContract.Mock.On("GetKpuProvinsiByAddress", mock.Anything, common.HexToAddress(address)).Return(contractKPU, nil).Once()
			},
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
				mockKPUContract.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			m := &Module{
				kpuProvinsiRepo: mockKPUProvinsiRepo,
				kpuContract:     mockKPUContract,
			}

			got, err := m.GetKPUProvinsiByAddress(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiByAddress() got = %v, want %v", got, tt.want)
			}

			tt.assertFn(t)
		})
	}
}

func TestModule_GetKPUProvinsiByID(t *testing.T) {
	mockKPUProvinsiRepo := &mocks_repository.KPUProvinsiRepository{}
	mockKPUContract := &mocksContract.KpuManagerInterface{}

	kpuID := uuid.New()
	userID := uuid.New()
	address := "0x123"
	registeredAt := time.Now()

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name     string
		args     args
		want     *response.KPUProvinsiResponse
		wantErr  bool
		fn       func()
		assertFn func(t *testing.T)
	}{
		{
			name: "ShouldError_KPUProvinsiNotFound",
			args: args{
				ctx: context.Background(),
				id:  kpuID,
			},
			want:    nil,
			wantErr: true,
			fn: func() {
				mockKPUProvinsiRepo.Mock.On("GetKPUProvinsiByID", mock.Anything, kpuID).Return(nil, errors.New("kpu provinsi not found")).Once()
			},
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
			},
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				id:  kpuID,
			},
			want: &response.KPUProvinsiResponse{
				ID:           kpuID.String(),
				UserID:       userID.String(),
				Username:     "kpu_user",
				Name:         "KPU Provinsi Jakarta",
				Address:      common.HexToAddress(address).String(), // Use formatted address
				Region:       "Jakarta",
				IsActive:     true,
				PhotoURL:     "/v1/kpu-provinsi/" + kpuID.String() + "/photo",
				Telephone:    "021-123456",
				RegisteredAt: registeredAt.String(),
			},
			wantErr: false,
			fn: func() {
				// Ensure address format matches
				addr := common.HexToAddress(address)

				kpuProvinsi := &model.KPUProvinsi{
					ID:           kpuID,
					UserID:       userID,
					Username:     "kpu_user",
					Name:         "KPU Provinsi Jakarta",
					Address:      addr.String(), // Use the same string format as contract
					Region:       "Jakarta",
					IsActive:     true,
					PhotoPath:    "/path/to/photo.jpg",
					Telephone:    "021-123456",
					RegisteredAt: registeredAt,
				}

				// Mock contract KPU Provinsi - using actual contract binding type
				contractKPU := kpuManager.IKPUManagerKPUProvinsi{
					Address: addr,
					Name:    "KPU Provinsi Jakarta",
					Region:  "Jakarta",
				}

				mockKPUProvinsiRepo.Mock.On("GetKPUProvinsiByID", mock.Anything, kpuID).Return(kpuProvinsi, nil).Once()
				mockKPUContract.Mock.On("GetKpuProvinsiByAddress", mock.Anything, addr).Return(contractKPU, nil).Once()
			},
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
				mockKPUContract.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			m := &Module{
				kpuProvinsiRepo: mockKPUProvinsiRepo,
				kpuContract:     mockKPUContract,
			}

			got, err := m.GetKPUProvinsiByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiByID() got = %v, want %v", got, tt.want)
			}

			tt.assertFn(t)
		})
	}
}

func TestModule_GetKPUProvinsiByUserID(t *testing.T) {
	type fields struct {
		kpuProvinsiRepo repository.KPUProvinsiRepository
		jwtSvc          jwtsvc.JWT
		txMgr           txmanager.TxManager
		publisher       event.MessagePublisher
		topics          config.KafkaTopics
		kpuContract     interfaces.KpuManagerInterface
		client          ethereum.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.KPUProvinsiResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				kpuProvinsiRepo: tt.fields.kpuProvinsiRepo,
				jwtSvc:          tt.fields.jwtSvc,
				txMgr:           tt.fields.txMgr,
				publisher:       tt.fields.publisher,
				topics:          tt.fields.topics,
				kpuContract:     tt.fields.kpuContract,
				client:          tt.fields.client,
			}
			got, err := m.GetKPUProvinsiByUserID(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetKPUProvinsiPhoto(t *testing.T) {
	type fields struct {
		kpuProvinsiRepo repository.KPUProvinsiRepository
		jwtSvc          jwtsvc.JWT
		txMgr           txmanager.TxManager
		publisher       event.MessagePublisher
		topics          config.KafkaTopics
		kpuContract     interfaces.KpuManagerInterface
		client          ethereum.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.File
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				kpuProvinsiRepo: tt.fields.kpuProvinsiRepo,
				jwtSvc:          tt.fields.jwtSvc,
				txMgr:           tt.fields.txMgr,
				publisher:       tt.fields.publisher,
				topics:          tt.fields.topics,
				kpuContract:     tt.fields.kpuContract,
				client:          tt.fields.client,
			}
			got, got1, err := m.GetKPUProvinsiPhoto(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiPhoto() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetKPUProvinsiPhoto() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestModule_GetKPUPusatByUserID(t *testing.T) {
	type fields struct {
		kpuProvinsiRepo repository.KPUProvinsiRepository
		jwtSvc          jwtsvc.JWT
		txMgr           txmanager.TxManager
		publisher       event.MessagePublisher
		topics          config.KafkaTopics
		kpuContract     interfaces.KpuManagerInterface
		client          ethereum.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.KPUProvinsiResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				kpuProvinsiRepo: tt.fields.kpuProvinsiRepo,
				jwtSvc:          tt.fields.jwtSvc,
				txMgr:           tt.fields.txMgr,
				publisher:       tt.fields.publisher,
				topics:          tt.fields.topics,
				kpuContract:     tt.fields.kpuContract,
				client:          tt.fields.client,
			}
			got, err := m.GetKPUPusatByUserID(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUPusatByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUPusatByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_RegisterKPUProvinsi(t *testing.T) {
	mockKPUProvinsiRepo := &mocks_repository.KPUProvinsiRepository{}
	mockJWTSvc := &mocks_interfaces.JWT{}
	mockPublisher := &mocksEvent.Publisher{}
	mockKPUContract := &mocksContract.KpuManagerInterface{}
	mockEthClient := &mocks.Client{}

	db, mockDb, _ := sqlmock.New()
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	dataStore := &sql.Store{
		Master: &sql.DB{DBConnection: sqlxDB},
		Slave:  &sql.DB{DBConnection: sqlxDB},
	}

	txMgr, err := txmanager.New(context.Background(), &txmanager.DriverConfig{
		Type:   "sql",
		Config: txSql.Config{DB: dataStore},
	})
	require.NoError(t, err)

	userID := uuid.New()
	reqCtx := libCtx.RequestContext{
		UserId: userID.String(),
	}
	ctx := context.WithValue(context.Background(), libCtx.RequestContextKey, reqCtx)

	type args struct {
		ctx context.Context
		req *request.KPUProvinsiRegistrationRequest
	}
	tests := []struct {
		name           string
		args           args
		wantErr        bool
		fn             func()
		expectBegin    bool
		expectRollback bool
		expectCommit   bool
		assertFn       func(t *testing.T)
	}{
		{
			name: "ShouldError_DuplicateKPUProvinsi",
			args: args{
				ctx: ctx,
				req: &request.KPUProvinsiRegistrationRequest{
					Username:          "kpu_user",
					Name:              "KPU Provinsi Jakarta",
					Address:           "0x123",
					Region:            "Jakarta",
					IsActive:          true,
					SignedTransaction: "0xsignedtx",
				},
			},
			wantErr: true,
			fn: func() {
				mockKPUProvinsiRepo.Mock.On("InsertKPUProvinsi", mock.Anything, mock.Anything).Return(dao.ErrDuplicate).Once()
			},
			expectBegin:    true,
			expectRollback: true,
			expectCommit:   false,
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
			},
		},
		{
			name: "ShouldError_FailedInsertKPUProvinsi",
			args: args{
				ctx: ctx,
				req: &request.KPUProvinsiRegistrationRequest{
					Username:          "kpu_user",
					Name:              "KPU Provinsi Jakarta",
					Address:           "0x123",
					Region:            "Jakarta",
					IsActive:          true,
					SignedTransaction: "0xsignedtx",
				},
			},
			wantErr: true,
			fn: func() {
				mockKPUProvinsiRepo.Mock.On("InsertKPUProvinsi", mock.Anything, mock.Anything).Return(errors.New("database error")).Once()
			},
			expectBegin:    true,
			expectRollback: true,
			expectCommit:   false,
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
			},
		},
		{
			name: "ShouldError_FailedBlockchainTransaction",
			args: args{
				ctx: ctx,
				req: &request.KPUProvinsiRegistrationRequest{
					Username:          "kpu_user",
					Name:              "KPU Provinsi Jakarta",
					Address:           "0x123",
					Region:            "Jakarta",
					IsActive:          true,
					SignedTransaction: "0xsignedtx",
				},
			},
			wantErr: true,
			fn: func() {
				mockKPUProvinsiRepo.Mock.On("InsertKPUProvinsi", mock.Anything, mock.Anything).Return(nil).Once()
				mockKPUProvinsiRepo.Mock.On("SendTxKPUProvinsiBlockchain", mock.Anything, "0xsignedtx").Return("", errors.New("blockchain error")).Once()
			},
			expectBegin:    true,
			expectRollback: true,
			expectCommit:   false,
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
			},
		},
		{
			name: "ShouldError_EmptyTransactionHash",
			args: args{
				ctx: ctx,
				req: &request.KPUProvinsiRegistrationRequest{
					Username:          "kpu_user",
					Name:              "KPU Provinsi Jakarta",
					Address:           "0x123",
					Region:            "Jakarta",
					IsActive:          true,
					SignedTransaction: "0xsignedtx",
				},
			},
			wantErr: true,
			fn: func() {
				mockKPUProvinsiRepo.Mock.On("InsertKPUProvinsi", mock.Anything, mock.Anything).Return(nil).Once()
				mockKPUProvinsiRepo.Mock.On("SendTxKPUProvinsiBlockchain", mock.Anything, "0xsignedtx").Return("", nil).Once()
			},
			expectBegin:    true,
			expectRollback: true,
			expectCommit:   false,
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
			},
		},
		{
			name: "ShouldError_FailedPublishEvent",
			args: args{
				ctx: ctx,
				req: &request.KPUProvinsiRegistrationRequest{
					Username:          "kpu_user",
					Name:              "KPU Provinsi Jakarta",
					Address:           "0x123",
					Region:            "Jakarta",
					IsActive:          true,
					SignedTransaction: "0xsignedtx",
				},
			},
			wantErr: true,
			fn: func() {
				mockKPUProvinsiRepo.Mock.On("InsertKPUProvinsi", mock.Anything, mock.Anything).Return(nil).Once()
				mockKPUProvinsiRepo.Mock.On("SendTxKPUProvinsiBlockchain", mock.Anything, "0xsignedtx").Return("0xtxhash", nil).Once()
				mockPublisher.Mock.On("Publish", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("publish error")).Once()
			},
			expectBegin:    true,
			expectRollback: true,
			expectCommit:   false,
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
				mockPublisher.AssertExpectations(t)
			},
		},
		{
			name: "Success",
			args: args{
				ctx: ctx,
				req: &request.KPUProvinsiRegistrationRequest{
					Username:          "kpu_user",
					Name:              "KPU Provinsi Jakarta",
					Address:           "0x123",
					Region:            "Jakarta",
					IsActive:          true,
					SignedTransaction: "0xsignedtx",
				},
			},
			wantErr: false,
			fn: func() {
				mockKPUProvinsiRepo.Mock.On("InsertKPUProvinsi", mock.Anything, mock.Anything).Return(nil).Once()
				mockKPUProvinsiRepo.Mock.On("SendTxKPUProvinsiBlockchain", mock.Anything, "0xsignedtx").Return("0xtxhash", nil).Once()
				mockPublisher.Mock.On("Publish", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectBegin:    true,
			expectRollback: false,
			expectCommit:   true,
			assertFn: func(t *testing.T) {
				mockKPUProvinsiRepo.AssertExpectations(t)
				mockPublisher.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			if tt.expectBegin {
				mockDb.ExpectBegin()
			}
			if tt.expectCommit {
				mockDb.ExpectCommit()
			}
			if tt.expectRollback {
				mockDb.ExpectRollback()
			}

			m := &Module{
				kpuProvinsiRepo: mockKPUProvinsiRepo,
				jwtSvc:          mockJWTSvc,
				txMgr:           txMgr,
				publisher:       mockPublisher,
				kpuContract:     mockKPUContract,
				client:          mockEthClient,
			}

			got, err := m.RegisterKPUProvinsi(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				require.NotNil(t, got)
				require.Equal(t, "0x123", got.Address)
				require.True(t, got.IsActive)
			}

			tt.assertFn(t)
		})
	}
}

func TestModule_UpdateKPUProvinsi(t *testing.T) {
	type fields struct {
		kpuProvinsiRepo repository.KPUProvinsiRepository
		jwtSvc          jwtsvc.JWT
		txMgr           txmanager.TxManager
		publisher       event.MessagePublisher
		topics          config.KafkaTopics
		kpuContract     interfaces.KpuManagerInterface
		client          ethereum.Client
	}
	type args struct {
		ctx           context.Context
		updateRequest *request.KPUProvinsiUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.KPUProvinsiResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				kpuProvinsiRepo: tt.fields.kpuProvinsiRepo,
				jwtSvc:          tt.fields.jwtSvc,
				txMgr:           tt.fields.txMgr,
				publisher:       tt.fields.publisher,
				topics:          tt.fields.topics,
				kpuContract:     tt.fields.kpuContract,
				client:          tt.fields.client,
			}
			got, err := m.UpdateKPUProvinsi(tt.args.ctx, tt.args.updateRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateKPUProvinsi() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_UploadKPUProvinsiPhoto(t *testing.T) {
	type fields struct {
		kpuProvinsiRepo repository.KPUProvinsiRepository
		jwtSvc          jwtsvc.JWT
		txMgr           txmanager.TxManager
		publisher       event.MessagePublisher
		topics          config.KafkaTopics
		kpuContract     interfaces.KpuManagerInterface
		client          ethereum.Client
	}
	type args struct {
		ctx      context.Context
		fileData io.Reader
		fileName string
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
			m := &Module{
				kpuProvinsiRepo: tt.fields.kpuProvinsiRepo,
				jwtSvc:          tt.fields.jwtSvc,
				txMgr:           tt.fields.txMgr,
				publisher:       tt.fields.publisher,
				topics:          tt.fields.topics,
				kpuContract:     tt.fields.kpuContract,
				client:          tt.fields.client,
			}
			if err := m.UploadKPUProvinsiPhoto(tt.args.ctx, tt.args.fileData, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("UploadKPUProvinsiPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
