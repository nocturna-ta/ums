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

func TestAPI_GetAllKPUKota(t *testing.T) {
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
			got, err := api.GetAllKPUKota(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKPUKota() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllKPUKota() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetAllKPUKota1(t *testing.T) {
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
			got, err := api.GetAllKPUKota(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllKPUKota() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllKPUKota() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetKPUKotaByUserID(t *testing.T) {
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
			got, err := api.GetKPUKotaByUserID(tt.args.ctx, tt.args.req)
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

func TestAPI_GetKPUKotaPhoto(t *testing.T) {
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
			got, err := api.GetKPUKotaPhoto(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUKotaPhoto() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_UpdateKPUKota(t *testing.T) {
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
			got, err := api.UpdateKPUKota(tt.args.ctx, tt.args.req)
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

func TestAPI_UploadKPUKotaPhoto(t *testing.T) {
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
			got, err := api.UploadKPUKotaPhoto(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadKPUKotaPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UploadKPUKotaPhoto() got = %v, want %v", got, tt.want)
			}
		})
	}
}
