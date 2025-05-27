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

func TestAPI_GetKPUPusat(t *testing.T) {
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
			got, err := api.GetKPUPusat(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUPusat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUPusat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_ApproveVerification(t *testing.T) {
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
			got, err := api.ApproveVerification(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApproveVerification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApproveVerification() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetKPUPusat1(t *testing.T) {
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
			got, err := api.GetKPUPusat(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUPusat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUPusat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetPendingVerificationDetails(t *testing.T) {
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
			got, err := api.GetPendingVerificationDetails(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingVerificationDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPendingVerificationDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetPendingVerificationsForRole(t *testing.T) {
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
			got, err := api.GetPendingVerificationsForRole(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingVerificationsForRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPendingVerificationsForRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_RejectVerification(t *testing.T) {
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
			got, err := api.RejectVerification(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RejectVerification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RejectVerification() got = %v, want %v", got, tt.want)
			}
		})
	}
}
