package controller

import (
	"context"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/ums/internal/usecases"
	"reflect"
	"testing"
	"time"
)

func TestAPI_ChangePassword(t *testing.T) {
	type fields struct {
		prefix         string
		port           uint
		readTimeout    time.Duration
		writeTimeout   time.Duration
		requestTimeout time.Duration
		enableSwagger  bool
		corsConfig     *router.CorsConfig
		voterUc        usecases.VoterUseCases
		userUc         usecases.UserUseCases
		kpuProvinsiUc  usecases.KPUProvinsiUseCases
		kpuKotaUc      usecases.KPUKotaUseCases
	}
	type args struct {
		ctx context.Context
		req *router.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rest.JSONResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				prefix:         tt.fields.prefix,
				port:           tt.fields.port,
				readTimeout:    tt.fields.readTimeout,
				writeTimeout:   tt.fields.writeTimeout,
				requestTimeout: tt.fields.requestTimeout,
				enableSwagger:  tt.fields.enableSwagger,
				corsConfig:     tt.fields.corsConfig,
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.ChangePassword(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangePassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_CheckVerificationStatus(t *testing.T) {
	type fields struct {
		prefix         string
		port           uint
		readTimeout    time.Duration
		writeTimeout   time.Duration
		requestTimeout time.Duration
		enableSwagger  bool
		corsConfig     *router.CorsConfig
		voterUc        usecases.VoterUseCases
		userUc         usecases.UserUseCases
		kpuProvinsiUc  usecases.KPUProvinsiUseCases
		kpuKotaUc      usecases.KPUKotaUseCases
	}
	type args struct {
		ctx context.Context
		req *router.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rest.JSONResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				prefix:         tt.fields.prefix,
				port:           tt.fields.port,
				readTimeout:    tt.fields.readTimeout,
				writeTimeout:   tt.fields.writeTimeout,
				requestTimeout: tt.fields.requestTimeout,
				enableSwagger:  tt.fields.enableSwagger,
				corsConfig:     tt.fields.corsConfig,
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.CheckVerificationStatus(tt.args.ctx, tt.args.req)
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

func TestAPI_GetByID(t *testing.T) {
	type fields struct {
		prefix         string
		port           uint
		readTimeout    time.Duration
		writeTimeout   time.Duration
		requestTimeout time.Duration
		enableSwagger  bool
		corsConfig     *router.CorsConfig
		voterUc        usecases.VoterUseCases
		userUc         usecases.UserUseCases
		kpuProvinsiUc  usecases.KPUProvinsiUseCases
		kpuKotaUc      usecases.KPUKotaUseCases
	}
	type args struct {
		ctx context.Context
		req *router.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rest.JSONResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				prefix:         tt.fields.prefix,
				port:           tt.fields.port,
				readTimeout:    tt.fields.readTimeout,
				writeTimeout:   tt.fields.writeTimeout,
				requestTimeout: tt.fields.requestTimeout,
				enableSwagger:  tt.fields.enableSwagger,
				corsConfig:     tt.fields.corsConfig,
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.GetByID(tt.args.ctx, tt.args.req)
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

func TestAPI_GetMyVerificationStatus(t *testing.T) {
	type fields struct {
		prefix         string
		port           uint
		readTimeout    time.Duration
		writeTimeout   time.Duration
		requestTimeout time.Duration
		enableSwagger  bool
		corsConfig     *router.CorsConfig
		voterUc        usecases.VoterUseCases
		userUc         usecases.UserUseCases
		kpuProvinsiUc  usecases.KPUProvinsiUseCases
		kpuKotaUc      usecases.KPUKotaUseCases
	}
	type args struct {
		ctx context.Context
		req *router.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rest.JSONResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				prefix:         tt.fields.prefix,
				port:           tt.fields.port,
				readTimeout:    tt.fields.readTimeout,
				writeTimeout:   tt.fields.writeTimeout,
				requestTimeout: tt.fields.requestTimeout,
				enableSwagger:  tt.fields.enableSwagger,
				corsConfig:     tt.fields.corsConfig,
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.GetMyVerificationStatus(tt.args.ctx, tt.args.req)
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

func TestAPI_GetUserByEmail(t *testing.T) {
	type fields struct {
		prefix         string
		port           uint
		readTimeout    time.Duration
		writeTimeout   time.Duration
		requestTimeout time.Duration
		enableSwagger  bool
		corsConfig     *router.CorsConfig
		voterUc        usecases.VoterUseCases
		userUc         usecases.UserUseCases
		kpuProvinsiUc  usecases.KPUProvinsiUseCases
		kpuKotaUc      usecases.KPUKotaUseCases
	}
	type args struct {
		ctx context.Context
		req *router.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rest.JSONResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				prefix:         tt.fields.prefix,
				port:           tt.fields.port,
				readTimeout:    tt.fields.readTimeout,
				writeTimeout:   tt.fields.writeTimeout,
				requestTimeout: tt.fields.requestTimeout,
				enableSwagger:  tt.fields.enableSwagger,
				corsConfig:     tt.fields.corsConfig,
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.GetUserByEmail(tt.args.ctx, tt.args.req)
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

func TestAPI_LoginUser(t *testing.T) {
	type fields struct {
		prefix         string
		port           uint
		readTimeout    time.Duration
		writeTimeout   time.Duration
		requestTimeout time.Duration
		enableSwagger  bool
		corsConfig     *router.CorsConfig
		voterUc        usecases.VoterUseCases
		userUc         usecases.UserUseCases
		kpuProvinsiUc  usecases.KPUProvinsiUseCases
		kpuKotaUc      usecases.KPUKotaUseCases
	}
	type args struct {
		ctx context.Context
		req *router.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rest.JSONResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				prefix:         tt.fields.prefix,
				port:           tt.fields.port,
				readTimeout:    tt.fields.readTimeout,
				writeTimeout:   tt.fields.writeTimeout,
				requestTimeout: tt.fields.requestTimeout,
				enableSwagger:  tt.fields.enableSwagger,
				corsConfig:     tt.fields.corsConfig,
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.LoginUser(tt.args.ctx, tt.args.req)
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

func TestAPI_RegisterUser(t *testing.T) {
	type fields struct {
		prefix         string
		port           uint
		readTimeout    time.Duration
		writeTimeout   time.Duration
		requestTimeout time.Duration
		enableSwagger  bool
		corsConfig     *router.CorsConfig
		voterUc        usecases.VoterUseCases
		userUc         usecases.UserUseCases
		kpuProvinsiUc  usecases.KPUProvinsiUseCases
		kpuKotaUc      usecases.KPUKotaUseCases
	}
	type args struct {
		ctx context.Context
		req *router.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rest.JSONResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				prefix:         tt.fields.prefix,
				port:           tt.fields.port,
				readTimeout:    tt.fields.readTimeout,
				writeTimeout:   tt.fields.writeTimeout,
				requestTimeout: tt.fields.requestTimeout,
				enableSwagger:  tt.fields.enableSwagger,
				corsConfig:     tt.fields.corsConfig,
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.RegisterUser(tt.args.ctx, tt.args.req)
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
