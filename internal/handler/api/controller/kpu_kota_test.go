package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	libCtx "github.com/nocturna-ta/golib/context"
	"github.com/nocturna-ta/golib/custerr"
	http2 "github.com/nocturna-ta/golib/http"
	response2 "github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/ums/internal/usecases/mocks_usecases"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAPI_GetAllKPUKota(t *testing.T) {
	mockedKPUKotaUc := &mocks_usecases.KPUKotaUseCases{}

	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		KpuKotaUc:      mockedKPUKotaUc,
	}

	userID := uuid.NewString()
	role := "kpu_kota"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	var res []response.KPUKotaResponse

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
				mockedKPUKotaUc.Mock.On("GetAllKPUKota", mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedKPUKotaUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result []response.KPUKotaResponse
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
				mockedKPUKotaUc.Mock.On("GetAllKPUKota", mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedKPUKotaUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedKPUKotaUc.Mock.On("GetAllKPUKota", mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 500, resp.StatusCode)
				mockedKPUKotaUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/kpu-kota", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_GetKPUKotaByUserID(t *testing.T) {
	mockedKPUKotaUc := &mocks_usecases.KPUKotaUseCases{}

	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		KpuKotaUc:      mockedKPUKotaUc,
	}

	userID := uuid.NewString()
	role := "kpu_kota"
	address := "0x1234567890abcdef1234567890abcdef12345678"
	var res response.KPUKotaResponse

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
				mockedKPUKotaUc.Mock.On("GetKPUKotaByUserID", mock.Anything).Return(&res, nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				mockedKPUKotaUc.AssertExpectations(t)
				require.Equalf(t, 200, resp.StatusCode, "Want status '%d', got '%d'", 200, resp.StatusCode)

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)

				var result response.KPUKotaResponse
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
				mockedKPUKotaUc.Mock.On("GetKPUKotaByUserID", mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedKPUKotaUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldBadRequest_WhenInvalidId",
			args: args{
				req: "12asf34",
			},
			fn: func() {
				mockedKPUKotaUc.Mock.On("GetKPUKotaByUserID", mock.Anything).Return(nil, &custerr.ErrChain{
					Message: "invalid user ID",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 500, resp.StatusCode)
				mockedKPUKotaUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/kpu-kota/id", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_GetKPUKotaPhoto(t *testing.T) {
	mockedKPUKotaUc := &mocks_usecases.KPUKotaUseCases{}

	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		KpuKotaUc:      mockedKPUKotaUc,
	}

	userID := uuid.NewString()
	role := "kpu_kota"
	address := "0x1234567890abcdef1234567890abcdef12345678"

	server = initServer(opts)

	tests := []struct {
		name     string
		fn       func()
		assertFn func(t *testing.T, resp *http.Response)
	}{
		{
			name: "SuccessGetValue",
			fn: func() {
				fileContent := []byte("fake-image-bytes")
				reader := bytes.NewReader(fileContent)

				file := &http2.File{
					Reader:      reader,
					FileName:    "kpu_photo.jpg",
					DisplayMode: "inline",
				}

				mockedKPUKotaUc.Mock.On("GetKPUKotaPhoto", mock.Anything).Return(file, "image/jpeg", nil).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				defer resp.Body.Close()

				require.Equal(t, 200, resp.StatusCode)
				require.Equal(t, "image/jpeg", resp.Header.Get("Content-Type"))
				require.Contains(t, resp.Header.Get("Content-Disposition"), `filename="kpu_photo.jpg"`)

				body, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				require.Equal(t, []byte("fake-image-bytes"), body)
			},
		},
		{
			name: "ShouldError_WhenFailedGet",
			fn: func() {
				mockedKPUKotaUc.Mock.On("GetKPUKotaPhoto", mock.Anything).Return(nil, "", &custerr.ErrChain{
					Message: "not found",
					Type:    response2.ErrBadRequest,
				}).Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equalf(t, 400, resp.StatusCode, "Want status '%d', got '%d'", 400, resp.StatusCode)
				mockedKPUKotaUc.AssertExpectations(t)
			},
		},
		{
			name: "ShouldError_WhenInternalFailure",
			fn: func() {
				mockedKPUKotaUc.Mock.
					On("GetKPUKotaPhoto", mock.Anything).
					Return(nil, "", errors.New("something went wrong")).
					Once()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equal(t, 500, resp.StatusCode)
				mockedKPUKotaUc.AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req := httptest.NewRequest("GET", "/v1/kpu-kota/photo", nil)

			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}

func TestAPI_UploadKPUKotaPhoto(t *testing.T) {
	mockedKPUKotaUc := &mocks_usecases.KPUKotaUseCases{}
	opts := &Options{
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		RequestTimeout: time.Minute,
		KpuKotaUc:      mockedKPUKotaUc,
	}

	userID := uuid.NewString()
	role := "kpu_kota"
	address := "0x1234567890abcdef1234567890abcdef12345678"

	server = initServer(opts)

	tests := []struct {
		name     string
		fn       func()
		reqBody  func() (*bytes.Buffer, string)
		assertFn func(t *testing.T, resp *http.Response)
	}{
		{
			name: "Success_UploadPhoto",
			fn: func() {
				mockedKPUKotaUc.
					Mock.
					On("UploadKPUKotaPhoto", mock.Anything, mock.Anything, "test.jpg").
					Return(nil).
					Once()
			},
			reqBody: func() (*bytes.Buffer, string) {
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)

				part, _ := writer.CreateFormFile("photo", "test.jpg")
				part.Write([]byte("fake image data"))
				writer.Close()

				return body, writer.FormDataContentType()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusOK, resp.StatusCode)
				defer resp.Body.Close()

				var jsonResp rest.JSONResponse
				err := json.NewDecoder(resp.Body).Decode(&jsonResp)
				require.NoError(t, err)
				require.Equal(t, "Photo uploaded successfully", jsonResp.Data)
			},
		},
		{
			name: "Fail_NoFileProvided",
			fn:   func() {}, // no UC call
			reqBody: func() (*bytes.Buffer, string) {
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)
				writer.Close()
				return body, writer.FormDataContentType()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
		{
			name: "Fail_InvalidFileFormat",
			fn:   func() {}, // UC not reached
			reqBody: func() (*bytes.Buffer, string) {
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)

				part, _ := writer.CreateFormFile("photo", "malicious.exe")
				part.Write([]byte("not a valid image"))
				writer.Close()

				return body, writer.FormDataContentType()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			},
		},
		{
			name: "Fail_InternalUploadError",
			fn: func() {
				mockedKPUKotaUc.
					Mock.
					On("UploadKPUKotaPhoto", mock.Anything, mock.Anything, "test.jpg").
					Return(&custerr.ErrChain{
						Message: "storage failed",
						Type:    response2.ErrInternalServerError,
					}).
					Once()
			},
			reqBody: func() (*bytes.Buffer, string) {
				body := &bytes.Buffer{}
				writer := multipart.NewWriter(body)

				part, _ := writer.CreateFormFile("photo", "test.jpg")
				part.Write([]byte("fake image data"))
				writer.Close()

				return body, writer.FormDataContentType()
			},
			assertFn: func(t *testing.T, resp *http.Response) {
				require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			body, contentType := tt.reqBody()

			req := httptest.NewRequest("POST", "/v1/kpu-kota/photo", body)

			req.Header.Set("Content-Type", contentType)
			req.Header.Add(libCtx.XUserId, userID)
			req.Header.Add(libCtx.XRole, role)
			req.Header.Add(libCtx.XAddressId, address)

			resp, err := server.Test(req)
			require.NoError(t, err)
			tt.assertFn(t, resp)
		})
	}
}
