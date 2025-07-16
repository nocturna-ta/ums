package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	response2 "github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/ums/internal/usecases/mocks_usecases"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAPI_GetLogs(t *testing.T) {
	mockedUserLogUc := &mocks_usecases.UserLogUseCases{}

	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		UserLogUc:      mockedUserLogUc,
	}

	userID := uuid.NewString()
	role := "kpu_provinsi"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	var res []response.UserLogResponse

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
				mockedUserLogUc.Mock.On("GetAllUserLog", mock.Anything, mock.Anything, mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedUserLogUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result []response.UserLogResponse
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
				mockedUserLogUc.Mock.On("GetAllUserLog", mock.Anything, mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserLogUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedUserLogUc.Mock.On("GetAllUserLog", mock.Anything, mock.Anything, mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedUserLogUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/user-logs?limit=abc&offset=10", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}
