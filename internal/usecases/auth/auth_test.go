package auth

import (
	"context"
	"github.com/nocturna-ta/ums/internal/interfaces/jwtsvc"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"reflect"
	"testing"
)

func TestModule_ValidateAuthorization1(t *testing.T) {
	type fields struct {
		jwtSvc jwtsvc.JWT
	}
	type args struct {
		ctx context.Context
		req *request.ValidateAuthorizationRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.ValidateAuthorizationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				jwtSvc: tt.fields.jwtSvc,
			}
			got, err := m.ValidateAuthorization(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAuthorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateAuthorization() got = %v, want %v", got, tt.want)
			}
		})
	}
}
