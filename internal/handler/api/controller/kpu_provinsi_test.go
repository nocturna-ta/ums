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

func TestAPI_GetAllKPUProvinsi(t *testing.T) {
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
			got, err := api.GetAllKPUProvinsi(tt.args.ctx, tt.args.req)
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

func TestAPI_GetAllKPUProvinsi1(t *testing.T) {
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
			got, err := api.GetAllKPUProvinsi(tt.args.ctx, tt.args.req)
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

func TestAPI_GetKPUProvinsiByUserID(t *testing.T) {
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
			got, err := api.GetKPUProvinsiByUserID(tt.args.ctx, tt.args.req)
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

func TestAPI_GetKPUProvinsiPhoto(t *testing.T) {
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
		want    *rest.AttachmentResponse
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
			got, err := api.GetKPUProvinsiPhoto(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinsiPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinsiPhoto() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_UpdateKPUProvinsi(t *testing.T) {
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
			got, err := api.UpdateKPUProvinsi(tt.args.ctx, tt.args.req)
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

func TestAPI_UploadKPUProvinsiPhoto(t *testing.T) {
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
			got, err := api.UploadKPUProvinsiPhoto(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadKPUProvinsiPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadKPUProvinsiPhoto() got = %v, want %v", got, tt.want)
			}
		})
	}
}
