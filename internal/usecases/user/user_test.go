package user

import (
	"context"
	"github.com/nocturna-ta/golib/event"
	"github.com/nocturna-ta/golib/txmanager"
	"github.com/nocturna-ta/ums/config"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"reflect"
	"testing"
)

func TestModule_ApproveUserVerification(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
		req *request.UserVerificationRequest
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
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			if err := m.ApproveUserVerification(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("ApproveUserVerification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModule_ChangePassword(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
		req *request.UserChangePasswordRequest
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
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			if err := m.ChangePassword(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModule_CheckVerificationStatus(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserVerificationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.CheckVerificationStatus(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckVerificationStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckVerificationStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetByID(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.GetByID(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetEnhancedVerificationStatus(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.EnhancedUserVerificationStatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.GetEnhancedVerificationStatus(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEnhancedVerificationStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnhancedVerificationStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetMyVerificationStatus(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserVerificationStatusResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.GetMyVerificationStatus(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMyVerificationStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMyVerificationStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetPendingVerificationUsers(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]response.UserVerificationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.GetPendingVerificationUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingVerificationUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPendingVerificationUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetPendingVerificationsByRole(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]response.UserVerificationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.GetPendingVerificationsByRole(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingVerificationsByRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPendingVerificationsByRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetUserByEmail(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_GetVerificationDetails(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx       context.Context
		userIDStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserVerificationDetailsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.GetVerificationDetails(tt.args.ctx, tt.args.userIDStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVerificationDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVerificationDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_LoginUser(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
		req *request.UserLoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserLoginResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.LoginUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_RegisterUser(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
		req *request.UserRegistrationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.UserRegistrationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			got, err := m.RegisterUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_RejectUserVerification(t *testing.T) {
	type fields struct {
		userRepo            repository.UserRepository
		pendingRegRepo      repository.PendingRegistrationRepository
		kpuProvinsiRepo     repository.KPUProvinsiRepository
		kpuKotaRepo         repository.KPUKotaRepository
		voterRepo           repository.VoterRepository
		kpuProvinsiUsecases usecases.KPUProvinsiUseCases
		kpuKotaUsecases     usecases.KPUKotaUseCases
		voterUsecases       usecases.VoterUseCases
		txMgr               txmanager.TxManager
		jwtSvc              jwtsvc.JWT
		publisher           event.MessagePublisher
		topics              config.KafkaTopics
	}
	type args struct {
		ctx context.Context
		req *request.UserVerificationRequest
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
				userRepo:            tt.fields.userRepo,
				pendingRegRepo:      tt.fields.pendingRegRepo,
				kpuProvinsiRepo:     tt.fields.kpuProvinsiRepo,
				kpuKotaRepo:         tt.fields.kpuKotaRepo,
				voterRepo:           tt.fields.voterRepo,
				kpuProvinsiUsecases: tt.fields.kpuProvinsiUsecases,
				kpuKotaUsecases:     tt.fields.kpuKotaUsecases,
				voterUsecases:       tt.fields.voterUsecases,
				txMgr:               tt.fields.txMgr,
				jwtSvc:              tt.fields.jwtSvc,
				publisher:           tt.fields.publisher,
				topics:              tt.fields.topics,
			}
			if err := m.RejectUserVerification(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("RejectUserVerification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
