package voter

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
	"reflect"
	"testing"
)

func TestModule_GetAllVoter(t *testing.T) {
	type fields struct {
		voterRepo     repository.VoterRepository
		jwtSvc        jwtsvc.JWT
		txMgr         txmanager.TxManager
		publisher     event.MessagePublisher
		topics        config.KafkaTopics
		voterContract interfaces.VoterManagerInterface
		client        ethereum.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]response.VoterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				voterRepo:     tt.fields.voterRepo,
				jwtSvc:        tt.fields.jwtSvc,
				txMgr:         tt.fields.txMgr,
				publisher:     tt.fields.publisher,
				topics:        tt.fields.topics,
				voterContract: tt.fields.voterContract,
				client:        tt.fields.client,
			}
			got, err := m.GetAllVoter(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllVoter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllVoter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetVoterByAddress(t *testing.T) {
	type fields struct {
		voterRepo     repository.VoterRepository
		jwtSvc        jwtsvc.JWT
		txMgr         txmanager.TxManager
		publisher     event.MessagePublisher
		topics        config.KafkaTopics
		voterContract interfaces.VoterManagerInterface
		client        ethereum.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.VoterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				voterRepo:     tt.fields.voterRepo,
				jwtSvc:        tt.fields.jwtSvc,
				txMgr:         tt.fields.txMgr,
				publisher:     tt.fields.publisher,
				topics:        tt.fields.topics,
				voterContract: tt.fields.voterContract,
				client:        tt.fields.client,
			}
			got, err := m.GetVoterByAddress(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterByAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetVoterByNIK(t *testing.T) {
	type fields struct {
		voterRepo     repository.VoterRepository
		jwtSvc        jwtsvc.JWT
		txMgr         txmanager.TxManager
		publisher     event.MessagePublisher
		topics        config.KafkaTopics
		voterContract interfaces.VoterManagerInterface
		client        ethereum.Client
	}
	type args struct {
		ctx context.Context
		nik string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.VoterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				voterRepo:     tt.fields.voterRepo,
				jwtSvc:        tt.fields.jwtSvc,
				txMgr:         tt.fields.txMgr,
				publisher:     tt.fields.publisher,
				topics:        tt.fields.topics,
				voterContract: tt.fields.voterContract,
				client:        tt.fields.client,
			}
			got, err := m.GetVoterByNIK(tt.args.ctx, tt.args.nik)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByNIK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterByNIK() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetVoterByRegion(t *testing.T) {
	type fields struct {
		voterRepo     repository.VoterRepository
		jwtSvc        jwtsvc.JWT
		txMgr         txmanager.TxManager
		publisher     event.MessagePublisher
		topics        config.KafkaTopics
		voterContract interfaces.VoterManagerInterface
		client        ethereum.Client
	}
	type args struct {
		ctx    context.Context
		region string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]response.VoterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				voterRepo:     tt.fields.voterRepo,
				jwtSvc:        tt.fields.jwtSvc,
				txMgr:         tt.fields.txMgr,
				publisher:     tt.fields.publisher,
				topics:        tt.fields.topics,
				voterContract: tt.fields.voterContract,
				client:        tt.fields.client,
			}
			got, err := m.GetVoterByRegion(tt.args.ctx, tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByRegion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterByRegion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetVoterByUserID(t *testing.T) {
	type fields struct {
		voterRepo     repository.VoterRepository
		jwtSvc        jwtsvc.JWT
		txMgr         txmanager.TxManager
		publisher     event.MessagePublisher
		topics        config.KafkaTopics
		voterContract interfaces.VoterManagerInterface
		client        ethereum.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.VoterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				voterRepo:     tt.fields.voterRepo,
				jwtSvc:        tt.fields.jwtSvc,
				txMgr:         tt.fields.txMgr,
				publisher:     tt.fields.publisher,
				topics:        tt.fields.topics,
				voterContract: tt.fields.voterContract,
				client:        tt.fields.client,
			}
			got, err := m.GetVoterByUserID(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterByUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetVoterKTPPhoto(t *testing.T) {
	type fields struct {
		voterRepo     repository.VoterRepository
		jwtSvc        jwtsvc.JWT
		txMgr         txmanager.TxManager
		publisher     event.MessagePublisher
		topics        config.KafkaTopics
		voterContract interfaces.VoterManagerInterface
		client        ethereum.Client
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
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
				voterRepo:     tt.fields.voterRepo,
				jwtSvc:        tt.fields.jwtSvc,
				txMgr:         tt.fields.txMgr,
				publisher:     tt.fields.publisher,
				topics:        tt.fields.topics,
				voterContract: tt.fields.voterContract,
				client:        tt.fields.client,
			}
			got, got1, err := m.GetVoterKTPPhoto(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterKTPPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterKTPPhoto() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetVoterKTPPhoto() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestModule_RegisterVoter(t *testing.T) {
	type fields struct {
		voterRepo     repository.VoterRepository
		jwtSvc        jwtsvc.JWT
		txMgr         txmanager.TxManager
		publisher     event.MessagePublisher
		topics        config.KafkaTopics
		voterContract interfaces.VoterManagerInterface
		client        ethereum.Client
	}
	type args struct {
		ctx context.Context
		req *request.VoterRegistrationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.VoterRegistrationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				voterRepo:     tt.fields.voterRepo,
				jwtSvc:        tt.fields.jwtSvc,
				txMgr:         tt.fields.txMgr,
				publisher:     tt.fields.publisher,
				topics:        tt.fields.topics,
				voterContract: tt.fields.voterContract,
				client:        tt.fields.client,
			}
			got, err := m.RegisterVoter(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterVoter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterVoter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
