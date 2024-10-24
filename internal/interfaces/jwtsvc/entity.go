package jwtsvc

import (
	"github.com/golang-jwt/jwt/v5"
)

type AccessClaims struct {
	*jwt.RegisteredClaims
	*JwtData
}

type JwtData struct {
	UserID string `json:"user_id,omitempty"`
}
type ClaimType string

const (
	AccessType ClaimType = "AccessType"
)
