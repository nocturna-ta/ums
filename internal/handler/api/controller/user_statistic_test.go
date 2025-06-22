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

func TestAPI_GetApprovedDPTStatistic(t *testing.T) {
	type fields struct {
		prefix          string
		port            uint
		readTimeout     time.Duration
		writeTimeout    time.Duration
		requestTimeout  time.Duration
		enableSwagger   bool
		voterUc         usecases.VoterUseCases
		userUc          usecases.UserUseCases
		kpuProvinsiUc   usecases.KPUProvinsiUseCases
		kpuKotaUc       usecases.KPUKotaUseCases
		userLogUc       usecases.UserLogUseCases
		userStatisticUc usecases.UserStatisticUseCases
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
				prefix:          tt.fields.prefix,
				port:            tt.fields.port,
				readTimeout:     tt.fields.readTimeout,
				writeTimeout:    tt.fields.writeTimeout,
				requestTimeout:  tt.fields.requestTimeout,
				enableSwagger:   tt.fields.enableSwagger,
				voterUc:         tt.fields.voterUc,
				userUc:          tt.fields.userUc,
				kpuProvinsiUc:   tt.fields.kpuProvinsiUc,
				kpuKotaUc:       tt.fields.kpuKotaUc,
				userLogUc:       tt.fields.userLogUc,
				userStatisticUc: tt.fields.userStatisticUc,
			}
			got, err := api.GetApprovedDPTStatistic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApprovedDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApprovedDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetKPUKotaStaffStatistic(t *testing.T) {
	type fields struct {
		prefix          string
		port            uint
		readTimeout     time.Duration
		writeTimeout    time.Duration
		requestTimeout  time.Duration
		enableSwagger   bool
		voterUc         usecases.VoterUseCases
		userUc          usecases.UserUseCases
		kpuProvinsiUc   usecases.KPUProvinsiUseCases
		kpuKotaUc       usecases.KPUKotaUseCases
		userLogUc       usecases.UserLogUseCases
		userStatisticUc usecases.UserStatisticUseCases
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
				prefix:          tt.fields.prefix,
				port:            tt.fields.port,
				readTimeout:     tt.fields.readTimeout,
				writeTimeout:    tt.fields.writeTimeout,
				requestTimeout:  tt.fields.requestTimeout,
				enableSwagger:   tt.fields.enableSwagger,
				voterUc:         tt.fields.voterUc,
				userUc:          tt.fields.userUc,
				kpuProvinsiUc:   tt.fields.kpuProvinsiUc,
				kpuKotaUc:       tt.fields.kpuKotaUc,
				userLogUc:       tt.fields.userLogUc,
				userStatisticUc: tt.fields.userStatisticUc,
			}
			got, err := api.GetKPUKotaStaffStatistic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUKotaStaffStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUKotaStaffStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetKPUProvinceStaffStatistic(t *testing.T) {
	type fields struct {
		prefix          string
		port            uint
		readTimeout     time.Duration
		writeTimeout    time.Duration
		requestTimeout  time.Duration
		enableSwagger   bool
		voterUc         usecases.VoterUseCases
		userUc          usecases.UserUseCases
		kpuProvinsiUc   usecases.KPUProvinsiUseCases
		kpuKotaUc       usecases.KPUKotaUseCases
		userLogUc       usecases.UserLogUseCases
		userStatisticUc usecases.UserStatisticUseCases
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
				prefix:          tt.fields.prefix,
				port:            tt.fields.port,
				readTimeout:     tt.fields.readTimeout,
				writeTimeout:    tt.fields.writeTimeout,
				requestTimeout:  tt.fields.requestTimeout,
				enableSwagger:   tt.fields.enableSwagger,
				voterUc:         tt.fields.voterUc,
				userUc:          tt.fields.userUc,
				kpuProvinsiUc:   tt.fields.kpuProvinsiUc,
				kpuKotaUc:       tt.fields.kpuKotaUc,
				userLogUc:       tt.fields.userLogUc,
				userStatisticUc: tt.fields.userStatisticUc,
			}
			got, err := api.GetKPUProvinceStaffStatistic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKPUProvinceStaffStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKPUProvinceStaffStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetKotaInformationDPTStatistic(t *testing.T) {
	type fields struct {
		prefix          string
		port            uint
		readTimeout     time.Duration
		writeTimeout    time.Duration
		requestTimeout  time.Duration
		enableSwagger   bool
		voterUc         usecases.VoterUseCases
		userUc          usecases.UserUseCases
		kpuProvinsiUc   usecases.KPUProvinsiUseCases
		kpuKotaUc       usecases.KPUKotaUseCases
		userLogUc       usecases.UserLogUseCases
		userStatisticUc usecases.UserStatisticUseCases
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
				prefix:          tt.fields.prefix,
				port:            tt.fields.port,
				readTimeout:     tt.fields.readTimeout,
				writeTimeout:    tt.fields.writeTimeout,
				requestTimeout:  tt.fields.requestTimeout,
				enableSwagger:   tt.fields.enableSwagger,
				voterUc:         tt.fields.voterUc,
				userUc:          tt.fields.userUc,
				kpuProvinsiUc:   tt.fields.kpuProvinsiUc,
				kpuKotaUc:       tt.fields.kpuKotaUc,
				userLogUc:       tt.fields.userLogUc,
				userStatisticUc: tt.fields.userStatisticUc,
			}
			got, err := api.GetKotaInformationDPTStatistic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKotaInformationDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKotaInformationDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetPendingDPTStatistic(t *testing.T) {
	type fields struct {
		prefix          string
		port            uint
		readTimeout     time.Duration
		writeTimeout    time.Duration
		requestTimeout  time.Duration
		enableSwagger   bool
		voterUc         usecases.VoterUseCases
		userUc          usecases.UserUseCases
		kpuProvinsiUc   usecases.KPUProvinsiUseCases
		kpuKotaUc       usecases.KPUKotaUseCases
		userLogUc       usecases.UserLogUseCases
		userStatisticUc usecases.UserStatisticUseCases
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
				prefix:          tt.fields.prefix,
				port:            tt.fields.port,
				readTimeout:     tt.fields.readTimeout,
				writeTimeout:    tt.fields.writeTimeout,
				requestTimeout:  tt.fields.requestTimeout,
				enableSwagger:   tt.fields.enableSwagger,
				voterUc:         tt.fields.voterUc,
				userUc:          tt.fields.userUc,
				kpuProvinsiUc:   tt.fields.kpuProvinsiUc,
				kpuKotaUc:       tt.fields.kpuKotaUc,
				userLogUc:       tt.fields.userLogUc,
				userStatisticUc: tt.fields.userStatisticUc,
			}
			got, err := api.GetPendingDPTStatistic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPendingDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPendingDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetProvinceInformationDPTStatistic(t *testing.T) {
	type fields struct {
		prefix          string
		port            uint
		readTimeout     time.Duration
		writeTimeout    time.Duration
		requestTimeout  time.Duration
		enableSwagger   bool
		voterUc         usecases.VoterUseCases
		userUc          usecases.UserUseCases
		kpuProvinsiUc   usecases.KPUProvinsiUseCases
		kpuKotaUc       usecases.KPUKotaUseCases
		userLogUc       usecases.UserLogUseCases
		userStatisticUc usecases.UserStatisticUseCases
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
				prefix:          tt.fields.prefix,
				port:            tt.fields.port,
				readTimeout:     tt.fields.readTimeout,
				writeTimeout:    tt.fields.writeTimeout,
				requestTimeout:  tt.fields.requestTimeout,
				enableSwagger:   tt.fields.enableSwagger,
				voterUc:         tt.fields.voterUc,
				userUc:          tt.fields.userUc,
				kpuProvinsiUc:   tt.fields.kpuProvinsiUc,
				kpuKotaUc:       tt.fields.kpuKotaUc,
				userLogUc:       tt.fields.userLogUc,
				userStatisticUc: tt.fields.userStatisticUc,
			}
			got, err := api.GetProvinceInformationDPTStatistic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProvinceInformationDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProvinceInformationDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetRejectedDPTStatistic(t *testing.T) {
	type fields struct {
		prefix          string
		port            uint
		readTimeout     time.Duration
		writeTimeout    time.Duration
		requestTimeout  time.Duration
		enableSwagger   bool
		voterUc         usecases.VoterUseCases
		userUc          usecases.UserUseCases
		kpuProvinsiUc   usecases.KPUProvinsiUseCases
		kpuKotaUc       usecases.KPUKotaUseCases
		userLogUc       usecases.UserLogUseCases
		userStatisticUc usecases.UserStatisticUseCases
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
				prefix:          tt.fields.prefix,
				port:            tt.fields.port,
				readTimeout:     tt.fields.readTimeout,
				writeTimeout:    tt.fields.writeTimeout,
				requestTimeout:  tt.fields.requestTimeout,
				enableSwagger:   tt.fields.enableSwagger,
				voterUc:         tt.fields.voterUc,
				userUc:          tt.fields.userUc,
				kpuProvinsiUc:   tt.fields.kpuProvinsiUc,
				kpuKotaUc:       tt.fields.kpuKotaUc,
				userLogUc:       tt.fields.userLogUc,
				userStatisticUc: tt.fields.userStatisticUc,
			}
			got, err := api.GetRejectedDPTStatistic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRejectedDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRejectedDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetTotalDPTStatistic(t *testing.T) {
	type fields struct {
		prefix          string
		port            uint
		readTimeout     time.Duration
		writeTimeout    time.Duration
		requestTimeout  time.Duration
		enableSwagger   bool
		voterUc         usecases.VoterUseCases
		userUc          usecases.UserUseCases
		kpuProvinsiUc   usecases.KPUProvinsiUseCases
		kpuKotaUc       usecases.KPUKotaUseCases
		userLogUc       usecases.UserLogUseCases
		userStatisticUc usecases.UserStatisticUseCases
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
				prefix:          tt.fields.prefix,
				port:            tt.fields.port,
				readTimeout:     tt.fields.readTimeout,
				writeTimeout:    tt.fields.writeTimeout,
				requestTimeout:  tt.fields.requestTimeout,
				enableSwagger:   tt.fields.enableSwagger,
				voterUc:         tt.fields.voterUc,
				userUc:          tt.fields.userUc,
				kpuProvinsiUc:   tt.fields.kpuProvinsiUc,
				kpuKotaUc:       tt.fields.kpuKotaUc,
				userLogUc:       tt.fields.userLogUc,
				userStatisticUc: tt.fields.userStatisticUc,
			}
			got, err := api.GetTotalDPTStatistic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTotalDPTStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTotalDPTStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_GetVotedStatistic(t *testing.T) {
	type fields struct {
		prefix          string
		port            uint
		readTimeout     time.Duration
		writeTimeout    time.Duration
		requestTimeout  time.Duration
		enableSwagger   bool
		voterUc         usecases.VoterUseCases
		userUc          usecases.UserUseCases
		kpuProvinsiUc   usecases.KPUProvinsiUseCases
		kpuKotaUc       usecases.KPUKotaUseCases
		userLogUc       usecases.UserLogUseCases
		userStatisticUc usecases.UserStatisticUseCases
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
				prefix:          tt.fields.prefix,
				port:            tt.fields.port,
				readTimeout:     tt.fields.readTimeout,
				writeTimeout:    tt.fields.writeTimeout,
				requestTimeout:  tt.fields.requestTimeout,
				enableSwagger:   tt.fields.enableSwagger,
				voterUc:         tt.fields.voterUc,
				userUc:          tt.fields.userUc,
				kpuProvinsiUc:   tt.fields.kpuProvinsiUc,
				kpuKotaUc:       tt.fields.kpuKotaUc,
				userLogUc:       tt.fields.userLogUc,
				userStatisticUc: tt.fields.userStatisticUc,
			}
			got, err := api.GetVotedStatistic(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVotedStatistic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVotedStatistic() got = %v, want %v", got, tt.want)
			}
		})
	}
}
