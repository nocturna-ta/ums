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

func TestAPI_GetLogs(t *testing.T) {
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
			got, err := api.GetLogs(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLogs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLogs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
