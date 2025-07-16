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

func TestAPI_GetAllVoter(t *testing.T) {
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
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.GetAllVoter(tt.args.ctx, tt.args.req)
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

func TestAPI_GetVoterByAddress(t *testing.T) {
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
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.GetVoterByAddress(tt.args.ctx, tt.args.req)
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

func TestAPI_GetVoterByNIK(t *testing.T) {
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
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.GetVoterByNIK(tt.args.ctx, tt.args.req)
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

func TestAPI_GetVoterByRegion(t *testing.T) {
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
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.GetVoterByRegion(tt.args.ctx, tt.args.req)
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

func TestAPI_GetVoterKTPPhoto(t *testing.T) {
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
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.GetVoterKTPPhoto(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoterKTPPhoto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoterKTPPhoto() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_RegisterVoter(t *testing.T) {
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
				voterUc:        tt.fields.voterUc,
				userUc:         tt.fields.userUc,
				kpuProvinsiUc:  tt.fields.kpuProvinsiUc,
				kpuKotaUc:      tt.fields.kpuKotaUc,
			}
			got, err := api.RegisterVoter(tt.args.ctx, tt.args.req)
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
