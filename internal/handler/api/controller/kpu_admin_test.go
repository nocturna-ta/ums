package controller

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	response2 "github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/ums/internal/usecases/mocks_usecases"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAPI_GetKPUPusat(t *testing.T) {
	mockedKPUProvinsiUc := &mocks_usecases.KPUProvinsiUseCases{}

	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		KpuProvinsiUc:  mockedKPUProvinsiUc,
	}

	userID := uuid.NewString()
	role := "kpu_pusat"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	res := response.KPUProvinsiResponse{}

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
				mockedKPUProvinsiUc.Mock.On("GetKPUPusatByUserID", mock.Anything, mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedKPUProvinsiUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result response.KPUProvinsiResponse
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
				mockedKPUProvinsiUc.Mock.On("GetKPUPusatByUserID", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedKPUProvinsiUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedKPUProvinsiUc.Mock.On("GetKPUPusatByUserID", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 500, resp.StatusCode)
				mockedKPUProvinsiUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/kpu-pusat/id", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_ApproveVerification(t *testing.T) {
	mockedUserUc := &mocks_usecases.UserUseCases{}

	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		UserUc:         mockedUserUc,
	}

	userID := uuid.NewString()
	role := "kpu_pusat"
	address := "0x1234567890abcdef1234567890abcdef12345678"

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
			name: "ShouldError_WrongRequest",
			args: args{
				req: "test",
			},
			fn: func() {

			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 500, resp.StatusCode, "Want status '%d', got '%d'", 500, resp.StatusCode)
			},
		},
		{
			name: "ShouldError_InvalidRequest",
			args: args{
				req: &request.UserVerificationRequest{},
			},
			fn: func() {

			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
			},
		},
		{
			name: "ShouldError_WhenFailedApprove",
			args: args{
				req: &request.UserVerificationRequest{
					UserID:            uuid.NewString(),
					AdminReason:       "test",
					SignedTransaction: "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
				},
			},
			fn: func() {

				mockedUserUc.Mock.On("ApproveUserVerification", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 500, resp.StatusCode, "Want status '%d', got '%d'", 500, resp.StatusCode)
				mockedUserUc.AssertExpectations(t)
			},
		},
		{
			name: "Success",
			args: args{
				req: &request.UserVerificationRequest{
					UserID:            uuid.NewString(),
					AdminReason:       "test",
					SignedTransaction: "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
				},
			},
			fn: func() {
				mockedUserUc.Mock.On("ApproveUserVerification", mock.Anything, mock.Anything).Return(nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedUserUc.AssertExpectations(t)

				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result string
				config := &mapstructure.DecoderConfig{
					TagName: "json",
				}
				config.Result = &result
				config.DecodeHook = mapstructure.ComposeDecodeHookFunc(toTimeHookFunc())
				decoder, _ := mapstructure.NewDecoder(config)
				err = decoder.Decode(jsonResp.Data)
				require.NoError(t, err)

				require.Equal(t, "User verification approved successfully", result)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			data, _ := json.Marshal(tt.args.req)
			req := httptest.NewRequest("POST", "/v1/verifications/approve", strings.NewReader(string(data)))

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_GetPendingVerificationDetails(t *testing.T) {
	mockUserUC := &mocks_usecases.UserUseCases{}

	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		UserUc:         mockUserUC,
	}

	userID := uuid.NewString()
	role := "kpu_pusat"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	res := response.UserVerificationDetailsResponse{}

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
				mockUserUC.Mock.On("GetVerificationDetails", mock.Anything, mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockUserUC.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result response.UserVerificationDetailsResponse
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
				mockUserUC.Mock.On("GetVerificationDetails", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockUserUC.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockUserUC.Mock.On("GetVerificationDetails", mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 500, resp.StatusCode)
				mockUserUC.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/verifications/details/"+userID, nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_GetPendingVerificationsForRole(t *testing.T) {
	mockedUserUc := &mocks_usecases.UserUseCases{}

	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		UserUc:         mockedUserUc,
	}

	userID := uuid.NewString()
	role := "kpu_pusat"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	var res []response.UserVerificationResponse

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
				mockedUserUc.Mock.On("GetPendingVerificationsByRole", mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedUserUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result []response.UserVerificationResponse
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
				mockedUserUc.Mock.On("GetPendingVerificationsByRole", mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedUserUc.Mock.On("GetPendingVerificationsByRole", mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 500, resp.StatusCode)
				mockedUserUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/verifications/pending", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_RejectVerification(t *testing.T) {
	mockedUserUc := &mocks_usecases.UserUseCases{}

	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		UserUc:         mockedUserUc,
	}

	userID := uuid.NewString()
	role := "kpu_pusat"
	address := "0x1234567890abcdef1234567890abcdef12345678"

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
			name: "ShouldError_WrongRequest",
			args: args{
				req: "test",
			},
			fn: func() {

			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 500, resp.StatusCode, "Want status '%d', got '%d'", 500, resp.StatusCode)
			},
		},
		{
			name: "ShouldError_InvalidRequest",
			args: args{
				req: &request.UserVerificationRequest{},
			},
			fn: func() {

			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
			},
		},
		{
			name: "ShouldError_WhenFailedApprove",
			args: args{
				req: &request.UserVerificationRequest{
					UserID:      uuid.NewString(),
					AdminReason: "test",
				},
			},
			fn: func() {
				mockedUserUc.Mock.On("RejectUserVerification", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 500, resp.StatusCode, "Want status '%d', got '%d'", 500, resp.StatusCode)
				mockedUserUc.AssertExpectations(t)
			},
		},
		{
			name: "Success",
			args: args{
				req: &request.UserVerificationRequest{
					UserID:      uuid.NewString(),
					AdminReason: "test",
				},
			},
			fn: func() {
				mockedUserUc.Mock.On("RejectUserVerification", mock.Anything, mock.Anything).Return(nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedUserUc.AssertExpectations(t)

				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result string
				config := &mapstructure.DecoderConfig{
					TagName: "json",
				}
				config.Result = &result
				config.DecodeHook = mapstructure.ComposeDecodeHookFunc(toTimeHookFunc())
				decoder, _ := mapstructure.NewDecoder(config)
				err = decoder.Decode(jsonResp.Data)
				require.NoError(t, err)

				require.Equal(t, "User verification rejected", result)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			data, _ := json.Marshal(tt.args.req)
			req := httptest.NewRequest("POST", "/v1/verifications/reject", strings.NewReader(string(data)))

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}
