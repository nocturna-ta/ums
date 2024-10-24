package jwtsvc

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nocturna-ta/ums/config"
	"reflect"
	"testing"
	"time"
)

func Test_jwtsvc_GenerateToken(t *testing.T) {
	type fields struct {
		secret string
	}
	type args struct {
		ctx    context.Context
		claims jwt.Claims
	}

	res := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMyNTExNzE4ODYxfQ.6MU0Da_ksCbX74nEf7iPOkp4S09Iqe1cb40o8sJLUOpfijrAH040Ona3DL703CXjc5fPRlHdseXDbO9P_sMdxw"

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		{
			name: "ShouldError_NoExpiration",
			fields: fields{
				secret: "test",
			},
			args: args{
				ctx:    context.Background(),
				claims: jwt.RegisteredClaims{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ShouldError_InvalidExpiration",
			fields: fields{
				secret: "test",
			},
			args: args{
				ctx: context.Background(),
				claims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				secret: "012dasd",
			},
			args: args{
				ctx: context.Background(),
				claims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Date(3000, 4, 4, 1, 1, 1, 0, time.UTC)),
				},
			},
			want:    &res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := New(&Options{
				Config: config.JWTConfig{
					Secret: tt.fields.secret,
				},
			})
			got, err := j.GenerateToken(tt.args.ctx, tt.args.claims)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jwtsvc_Validate(t *testing.T) {
	token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjMyNTExNzE4ODYxfQ.6MU0Da_ksCbX74nEf7iPOkp4S09Iqe1cb40o8sJLUOpfijrAH040Ona3DL703CXjc5fPRlHdseXDbO9P_sMdxw"

	type fields struct {
		secret []byte
	}
	type args struct {
		ctx       context.Context
		token     string
		claimType ClaimType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    jwt.Claims
		wantErr bool
	}{
		{
			name:   "ShouldError_InvalidClaim",
			fields: fields{secret: []byte("012dasd")},
			args: args{
				ctx:       context.Background(),
				token:     token,
				claimType: ClaimType("reg"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "ShouldError_InvalidToken",
			fields: fields{secret: []byte("012dasd")},
			args: args{
				ctx:       context.Background(),
				token:     "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTA4MTQ2ODF9.x1x_LDLLMRZN4KfyHJzujaJReSaZEfpCAfqKcBXKQfUizlCust1a5PmkidfHuk_pbqltbGHXrhbDtQsgtRsTbg",
				claimType: AccessType,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "Success",
			fields: fields{secret: []byte("012dasd")},
			args: args{
				ctx:       context.Background(),
				token:     token,
				claimType: AccessType,
			},
			want: &AccessClaims{
				RegisteredClaims: &jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Date(3000, 4, 4, 8, 1, 1, 0, time.UTC))},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwtsvc{
				secret: tt.fields.secret,
			}
			_, err := j.Validate(tt.args.ctx, tt.args.token, tt.args.claimType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
