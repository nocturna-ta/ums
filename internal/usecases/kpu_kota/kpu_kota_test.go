package kpu_kota

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/http"
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/model"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/domain/repository/mocks_repository"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/votechain-contract/binding/kpuManager"
	"github.com/nocturna-ta/votechain-contract/interfaces"
	mocksContract "github.com/nocturna-ta/votechain-contract/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"io"
	"reflect"
	"testing"
)

func TestModule_GetAllKPUKota(t *testing.T) {

	expected := &[]model.KPUKota{}
	expectedContract := []kpuManager.IKPUManagerKPUKota{
		{
			Name:     "Test KPU Kota",
			Region:   "Test Region",
			Address:  common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			IsActive: false,
		},
	}

	reqCtx := libCtx.RequestContext{
		UserId:  uuid.NewString(),
		Address: "0x1234567890abcdef1234567890abcdef12345678",
		Role:    "kpu_kota",
	}

	ctx := context.WithValue(context.Background(), libCtx.RequestContextKey, reqCtx)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		args     args
		fn       func(repo *mocks_repository.KPUKotaRepository, contract *mocksContract.KpuManagerInterface)
		wantResp *[]response.KPUKotaResponse
		wantErr  bool
	}{
		{
			name: "ShouldError_WhenNoRequestContext",
			args: args{
				ctx: context.Background(),
			},
			wantResp: nil,
			wantErr:  true,
			fn: func(repo *mocks_repository.KPUKotaRepository, contract *mocksContract.KpuManagerInterface) {
			},
		},
		{
			name: "ShouldError_WhenFailedGetAllKPUKota",
			args: args{
				ctx: ctx,
			},
			fn: func(repo *mocks_repository.KPUKotaRepository, contract *mocksContract.KpuManagerInterface) {
				repo.On("GetAllKPUKota", mock.Anything).Return(nil, errors.New("Failed")).Once()
				contract.On("GetAllKPUKota", mock.Anything).Return(nil, errors.New("Failed")).Once()
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "ShouldSuccess_WhenGetAllKPUKota",
			args: args{
				ctx: ctx,
			},
			fn: func(repo *mocks_repository.KPUKotaRepository, contract *mocksContract.KpuManagerInterface) {
				repo.On("GetAllKPUKota", mock.Anything).Return(expected, nil).Once()
				contract.On("GetAllKPUKota", mock.Anything).Return(expectedContract, nil).Once()
			},
			wantResp: &[]response.KPUKotaResponse{
				{
					ID:           uuid.NewString(),
					UserID:       uuid.NewString(),
					Username:     "testuser",
					Address:      "0x1234567890abcdef1234567890abcdef12345678",
					Name:         "Test KPU Kota",
					Region:       "Test Region",
					IsActive:     false,
					PhotoURL:     "http://example.com/photo.jpg",
					Telephone:    "1234567890",
					RegisteredAt: "2023-10-01T00:00:00Z",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedKPUKotaRepo := &mocks_repository.KPUKotaRepository{}
			mockedKPUContract := &mocksContract.KpuManagerInterface{}

			if tt.fn != nil {
				tt.fn(mockedKPUKotaRepo, mockedKPUContract)
			}

			m := New(&Opts{
				KpuKotaRepo: mockedKPUKotaRepo,
				KpuContract: mockedKPUContract,
			})

			_, err := m.GetAllKPUKota(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKPUKota() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestModule_GetKPUKotaByAddress(t *testing.T) {
	type fields struct {
		kpuKotaRepo repository.KPUKotaRepository
		jwtSvc      jwtsvc.JWT
		txMgr       txmanager.TxManager
		publisher   event.MessagePublisher
		topics      config.KafkaTopics
		kpuContract interfaces.KpuManagerInterface
		client      ethereum.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.KPUKotaResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				kpuKotaRepo: tt.fields.kpuKotaRepo,
				jwtSvc:      tt.fields.jwtSvc,
				txMgr:       tt.fields.txMgr,
				publisher:   tt.fields.publisher,
				topics:      tt.fields.topics,
				kpuContract: tt.fields.kpuContract,
				client:      tt.fields.client,
			}
			got, err := m.GetKPUKotaByAddress(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUKotaByAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetKPUKotaByID(t *testing.T) {
	type fields struct {
		kpuKotaRepo repository.KPUKotaRepository
		jwtSvc      jwtsvc.JWT
		txMgr       txmanager.TxManager
		publisher   event.MessagePublisher
		topics      config.KafkaTopics
		kpuContract interfaces.KpuManagerInterface
		client      ethereum.Client
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.KPUKotaResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				kpuKotaRepo: tt.fields.kpuKotaRepo,
				jwtSvc:      tt.fields.jwtSvc,
				txMgr:       tt.fields.txMgr,
				publisher:   tt.fields.publisher,
				topics:      tt.fields.topics,
				kpuContract: tt.fields.kpuContract,
				client:      tt.fields.client,
			}
			got, err := m.GetKPUKotaByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUKotaByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetKPUKotaByUserID(t *testing.T) {
	type fields struct {
		kpuKotaRepo repository.KPUKotaRepository
		jwtSvc      jwtsvc.JWT
		txMgr       txmanager.TxManager
		publisher   event.MessagePublisher
		topics      config.KafkaTopics
		kpuContract interfaces.KpuManagerInterface
		client      ethereum.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.KPUKotaResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				kpuKotaRepo: tt.fields.kpuKotaRepo,
				jwtSvc:      tt.fields.jwtSvc,
				txMgr:       tt.fields.txMgr,
				publisher:   tt.fields.publisher,
				topics:      tt.fields.topics,
				kpuContract: tt.fields.kpuContract,
				client:      tt.fields.client,
			}
			got, err := m.GetKPUKotaByUserID(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUKotaByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetKPUKotaPhoto(t *testing.T) {
	type fields struct {
		kpuKotaRepo repository.KPUKotaRepository
		jwtSvc      jwtsvc.JWT
		txMgr       txmanager.TxManager
		publisher   event.MessagePublisher
		topics      config.KafkaTopics
		kpuContract interfaces.KpuManagerInterface
		client      ethereum.Client
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
				kpuKotaRepo: tt.fields.kpuKotaRepo,
				jwtSvc:      tt.fields.jwtSvc,
				txMgr:       tt.fields.txMgr,
				publisher:   tt.fields.publisher,
				topics:      tt.fields.topics,
				kpuContract: tt.fields.kpuContract,
				client:      tt.fields.client,
			}
			got, got1, err := m.GetKPUKotaPhoto(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUKotaPhoto() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetKPUKotaPhoto() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestModule_RegisterKPUKota(t *testing.T) {
	type fields struct {
		kpuKotaRepo repository.KPUKotaRepository
		jwtSvc      jwtsvc.JWT
		txMgr       txmanager.TxManager
		publisher   event.MessagePublisher
		topics      config.KafkaTopics
		kpuContract interfaces.KpuManagerInterface
		client      ethereum.Client
	}
	type args struct {
		ctx context.Context
		req *request.KPUKotaRegistrationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.KPUKotaRegistrationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				kpuKotaRepo: tt.fields.kpuKotaRepo,
				jwtSvc:      tt.fields.jwtSvc,
				txMgr:       tt.fields.txMgr,
				publisher:   tt.fields.publisher,
				topics:      tt.fields.topics,
				kpuContract: tt.fields.kpuContract,
				client:      tt.fields.client,
			}
			got, err := m.RegisterKPUKota(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKPUKota() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterKPUKota() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_UpdateKPUKota(t *testing.T) {
	type fields struct {
		kpuKotaRepo repository.KPUKotaRepository
		jwtSvc      jwtsvc.JWT
		txMgr       txmanager.TxManager
		publisher   event.MessagePublisher
		topics      config.KafkaTopics
		kpuContract interfaces.KpuManagerInterface
		client      ethereum.Client
	}
	type args struct {
		ctx           context.Context
		updateRequest *request.KPUKotaUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.KPUKotaResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				kpuKotaRepo: tt.fields.kpuKotaRepo,
				jwtSvc:      tt.fields.jwtSvc,
				txMgr:       tt.fields.txMgr,
				publisher:   tt.fields.publisher,
				topics:      tt.fields.topics,
				kpuContract: tt.fields.kpuContract,
				client:      tt.fields.client,
			}
			got, err := m.UpdateKPUKota(tt.args.ctx, tt.args.updateRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateKPUKota() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateKPUKota() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_UploadKPUKotaPhoto(t *testing.T) {
	type fields struct {
		kpuKotaRepo repository.KPUKotaRepository
		jwtSvc      jwtsvc.JWT
		txMgr       txmanager.TxManager
		publisher   event.MessagePublisher
		topics      config.KafkaTopics
		kpuContract interfaces.KpuManagerInterface
		client      ethereum.Client
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
				kpuKotaRepo: tt.fields.kpuKotaRepo,
				jwtSvc:      tt.fields.jwtSvc,
				txMgr:       tt.fields.txMgr,
				publisher:   tt.fields.publisher,
				topics:      tt.fields.topics,
				kpuContract: tt.fields.kpuContract,
				client:      tt.fields.client,
			}
			if err := m.UploadKPUKotaPhoto(tt.args.ctx, tt.args.fileData, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("UploadKPUKotaPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
