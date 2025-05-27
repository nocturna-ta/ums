package kpu_provinsi

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/ethereum"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/http"
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/votechain-contract/interfaces"
	"io"
	"reflect"
	"testing"
)

func TestModule_GetAllKPUProvinsi(t *testing.T) {
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
		want    *[]response.KPUProvinsiResponse
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
			got, err := m.GetAllKPUProvinsi(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllKPUProvinsi() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetKPUProvinsiByAddress(t *testing.T) {
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
			got, err := m.GetKPUProvinsiByAddress(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiByAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetKPUProvinsiByID(t *testing.T) {
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
		id  uuid.UUID
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
			got, err := m.GetKPUProvinsiByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiByID() got = %v, want %v", got, tt.want)
			}
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
		req *request.KPUProvinsiRegistrationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.KPUProvinsiRegistrationResponse
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
			got, err := m.RegisterKPUProvinsi(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterKPUProvinsi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterKPUProvinsi() got = %v, want %v", got, tt.want)
			}
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
