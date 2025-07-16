package controller

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	response2 "github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/ums/internal/usecases"
	"github.com/nocturna-ta/ums/internal/usecases/mocks_usecases"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestAPI_GetApprovedDPTStatistic(t *testing.T) {
	mockedUserStatisticUc := &mocks_usecases.UserStatisticUseCases{}

	opts := &Options{
		ReadTimeout:     time.Minute,
		WriteTimeout:    time.Minute,
		RequestTimeout:  time.Minute,
		UserStatisticUc: mockedUserStatisticUc,
	}

	userID := uuid.NewString()
	role := "kpu_provinsi"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	var res response.ApprovedDPTResponse

	server = initServer(opts)

	type args struct {
		req any
	}

	tests := []struct {
		name     string
		args     args
		fn       func()
		assertFn func(t *testing.T, resp *http.Response)
	}{
		{
			name: "SuccessGetValue",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetApprovedDPTStatistic", mock.Anything, mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedUserStatisticUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result response.ApprovedDPTResponse
				config := &mapstructure.DecoderConfig{
					TagName: "json",
				}
				config.Result = &result
				config.DecodeHook = mapstructure.ComposeDecodeHookFunc(toTimeHookFunc())
				decoder, _ := mapstructure.NewDecoder(config)
				err = decoder.Decode(jsonResp.Data)
				require.NoError(t, err)

				require.Equal(t, res, result)
			},
		},
		{
			name: "ShouldError_WhenFailedGet",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetApprovedDPTStatistic", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetApprovedDPTStatistic", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/user-statistic/approved-dpt?region=provinsi", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_GetKPUKotaStaffStatistic(t *testing.T) {
	mockedUserStatisticUc := &mocks_usecases.UserStatisticUseCases{}

	opts := &Options{
		ReadTimeout:     time.Minute,
		WriteTimeout:    time.Minute,
		RequestTimeout:  time.Minute,
		UserStatisticUc: mockedUserStatisticUc,
	}

	userID := uuid.NewString()
	role := "kpu_provinsi"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	var res response.StaffKPUResponse

	server = initServer(opts)

	type args struct {
		req any
	}

	tests := []struct {
		name     string
		args     args
		fn       func()
		assertFn func(t *testing.T, resp *http.Response)
	}{
		{
			name: "SuccessGetValue",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetStaffKPUKotaStatistic", mock.Anything, mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedUserStatisticUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result response.StaffKPUResponse
				config := &mapstructure.DecoderConfig{
					TagName: "json",
				}
				config.Result = &result
				config.DecodeHook = mapstructure.ComposeDecodeHookFunc(toTimeHookFunc())
				decoder, _ := mapstructure.NewDecoder(config)
				err = decoder.Decode(jsonResp.Data)
				require.NoError(t, err)

				require.Equal(t, res, result)
			},
		},
		{
			name: "ShouldError_WhenFailedGet",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetStaffKPUKotaStatistic", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetStaffKPUKotaStatistic", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/user-statistic/kpu-kota-staff?region=provinsi", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_GetKPUProvinceStaffStatistic(t *testing.T) {
	mockedUserStatisticUc := &mocks_usecases.UserStatisticUseCases{}

	opts := &Options{
		ReadTimeout:     time.Minute,
		WriteTimeout:    time.Minute,
		RequestTimeout:  time.Minute,
		UserStatisticUc: mockedUserStatisticUc,
	}

	userID := uuid.NewString()
	role := "kpu_pusat"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	var res response.StaffKPUResponse

	server = initServer(opts)

	type args struct {
		req any
	}

	tests := []struct {
		name     string
		args     args
		fn       func()
		assertFn func(t *testing.T, resp *http.Response)
	}{
		{
			name: "SuccessGetValue",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetStaffKPUProvinceStatistic", mock.Anything, mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedUserStatisticUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result response.StaffKPUResponse
				config := &mapstructure.DecoderConfig{
					TagName: "json",
				}
				config.Result = &result
				config.DecodeHook = mapstructure.ComposeDecodeHookFunc(toTimeHookFunc())
				decoder, _ := mapstructure.NewDecoder(config)
				err = decoder.Decode(jsonResp.Data)
				require.NoError(t, err)

				require.Equal(t, res, result)
			},
		},
		{
			name: "ShouldError_WhenFailedGet",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetStaffKPUProvinceStatistic", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetStaffKPUProvinceStatistic", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/user-statistic/kpu-provinsi-staff?region=provinsi", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_GetKotaInformationDPTStatistic(t *testing.T) {
	mockedUserStatisticUc := &mocks_usecases.UserStatisticUseCases{}

	opts := &Options{
		ReadTimeout:     time.Minute,
		WriteTimeout:    time.Minute,
		RequestTimeout:  time.Minute,
		UserStatisticUc: mockedUserStatisticUc,
	}

	userID := uuid.NewString()
	role := "kpu_provinsi"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	var res []response.DPTInformationResponse

	server = initServer(opts)

	type args struct {
		req any
	}

	tests := []struct {
		name     string
		args     args
		fn       func()
		assertFn func(t *testing.T, resp *http.Response)
	}{
		{
			name: "SuccessGetValue",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetKotaInformationDPTStatistic", mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedUserStatisticUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result []response.DPTInformationResponse
				config := &mapstructure.DecoderConfig{
					TagName: "json",
				}
				config.Result = &result
				config.DecodeHook = mapstructure.ComposeDecodeHookFunc(toTimeHookFunc())
				decoder, _ := mapstructure.NewDecoder(config)
				err = decoder.Decode(jsonResp.Data)
				require.NoError(t, err)

				require.Equal(t, res, result)
			},
		},
		{
			name: "ShouldError_WhenFailedGet",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetKotaInformationDPTStatistic", mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetKotaInformationDPTStatistic", mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/user-statistic/kota-information-dpt", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_GetPendingDPTStatistic(t *testing.T) {
	mockedUserStatisticUc := &mocks_usecases.UserStatisticUseCases{}

	opts := &Options{
		ReadTimeout:     time.Minute,
		WriteTimeout:    time.Minute,
		RequestTimeout:  time.Minute,
		UserStatisticUc: mockedUserStatisticUc,
	}

	userID := uuid.NewString()
	role := "kpu_provinsi"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	var res response.PendingDPTResponse

	server = initServer(opts)

	type args struct {
		req any
	}

	tests := []struct {
		name     string
		args     args
		fn       func()
		assertFn func(t *testing.T, resp *http.Response)
	}{
		{
			name: "SuccessGetValue",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetPendingDPTStatistic", mock.Anything, mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedUserStatisticUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result response.PendingDPTResponse
				config := &mapstructure.DecoderConfig{
					TagName: "json",
				}
				config.Result = &result
				config.DecodeHook = mapstructure.ComposeDecodeHookFunc(toTimeHookFunc())
				decoder, _ := mapstructure.NewDecoder(config)
				err = decoder.Decode(jsonResp.Data)
				require.NoError(t, err)

				require.Equal(t, res, result)
			},
		},
		{
			name: "ShouldError_WhenFailedGet",
			args: args{
				req: "",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetPendingDPTStatistic", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedUserStatisticUc.Mock.On("GetPendingDPTStatistic", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserStatisticUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/user-statistic/pending-dpt?region=provinsi", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
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
