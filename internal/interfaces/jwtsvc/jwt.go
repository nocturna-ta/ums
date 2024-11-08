package jwtsvc

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/config"
	"time"
)

type JWT interface {
	GenerateToken(ctx context.Context, claims jwt.Claims) (*string, error)
	Validate(ctx context.Context, token string, claimType ClaimType) (jwt.Claims, error)
}

type jwtsvc struct {
	secret []byte
}

type Options struct {
	Config config.JWTConfig
}

func New(opts *Options) JWT {
	return &jwtsvc{
		secret: []byte(opts.Config.Secret),
	}
}

func (j *jwtsvc) GenerateToken(ctx context.Context, claims jwt.Claims) (*string, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "JWT.GenerateToken")
	defer span.End()

	exp, _ := claims.GetExpirationTime()
	if exp == nil || exp.IsZero() || exp.Before(time.Now()) {
		return nil, fmt.Errorf("token should have valid expiration")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, _ := token.SignedString(j.secret)

	return &tokenStr, nil
}

func (j *jwtsvc) Validate(ctx context.Context, token string, claimType ClaimType) (jwt.Claims, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "JWT.Validate")
	defer span.End()

	var (
		parsedToken *jwt.Token
		err         error
	)

	switch claimType {
	case AccessType:
		parsedToken, err = j.parseAccessClaim(token)
	default:
		return nil, fmt.Errorf("invalid claim type: %v", claimType)
	}

	if err != nil {
		return nil, err
	}

	return parsedToken.Claims, err
}

func (j *jwtsvc) parseAccessClaim(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &AccessClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})
}
