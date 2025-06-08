package auth

import (
	"context"
	"errors"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/interfaces/mocks_interfaces"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"github.com/nocturna-ta/ums/pkg/constants"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"testing"
)

func TestModule_ValidateAuthorization(t *testing.T) {
	mockJWTSvc := &mocks_interfaces.JWT{}

	type args struct {
		ctx context.Context
		req *request.ValidateAuthorizationRequest
	}
	tests := []struct {
		name     string
		args     args
		want     *response.ValidateAuthorizationResponse
		wantErr  bool
		fn       func()
		assertFn func(t *testing.T)
	}{
		{
			name: "ShouldError_MissingAuthorizationHeader",
			args: args{
				ctx: context.Background(),
				req: &request.ValidateAuthorizationRequest{
					Headers:       map[string]string{},
					Path:          "/test",
					TargetService: "test-service",
				},
			},
			want:    nil,
			wantErr: true,
			fn:      func() {},
			assertFn: func(t *testing.T) {
				// No JWT service calls expected
			},
		},
		{
			name: "ShouldError_EmptyAuthorizationHeader",
			args: args{
				ctx: context.Background(),
				req: &request.ValidateAuthorizationRequest{
					Headers: map[string]string{
						constants.Authorization: "",
					},
					Path:          "/test",
					TargetService: "test-service",
				},
			},
			want:    nil,
			wantErr: true,
			fn:      func() {},
			assertFn: func(t *testing.T) {
				// No JWT service calls expected
			},
		},
		{
			name: "ShouldError_MissingToken",
			args: args{
				ctx: context.Background(),
				req: &request.ValidateAuthorizationRequest{
					Headers: map[string]string{
						constants.Authorization: "Bearer ",
					},
					Path:          "/test",
					TargetService: "test-service",
				},
			},
			want:    nil,
			wantErr: true,
			fn:      func() {},
			assertFn: func(t *testing.T) {
				// No JWT service calls expected
			},
		},
		{
			name: "ShouldError_TokenValidationFailed",
			args: args{
				ctx: context.Background(),
				req: &request.ValidateAuthorizationRequest{
					Headers: map[string]string{
						constants.Authorization: "Bearer invalid-token",
					},
					Path:          "/test",
					TargetService: "test-service",
				},
			},
			want:    nil,
			wantErr: true,
			fn: func() {
				mockJWTSvc.Mock.On("Validate", mock.Anything, mock.Anything, mock.Anything).
					Return(nil, errors.New("invalid token")).Once()
			},
			assertFn: func(t *testing.T) {
				mockJWTSvc.AssertExpectations(t)
			},
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				req: &request.ValidateAuthorizationRequest{
					Headers: map[string]string{
						constants.Authorization: "Bearer valid-token",
					},
					Path:          "/test",
					TargetService: "test-service",
				},
			},
			want: &response.ValidateAuthorizationResponse{
				IsValid: true,
				ExplodeHeader: map[string]string{
					"X-User-Id": "test-user-id",
					"Role":      "test-role",
				},
			},
			wantErr: false,
			fn: func() {
				claims := &jwtsvc.AccessClaims{
					JwtData: &jwtsvc.JwtData{
						UserID: "test-user-id",
						Role:   "test-role",
					},
				}
				mockJWTSvc.Mock.On("Validate", mock.Anything, mock.Anything, mock.Anything).
					Return(claims, nil).Once()
			},
			assertFn: func(t *testing.T) {
				mockJWTSvc.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			m := &Module{
				jwtSvc: mockJWTSvc,
			}

			got, err := m.ValidateAuthorization(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAuthorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				require.Error(t, err)
				grpcErr, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, grpcErr.Code())
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ValidateAuthorization() got = %v, want %v", got, tt.want)
				}
			}

			tt.assertFn(t)
		})
	}
}
