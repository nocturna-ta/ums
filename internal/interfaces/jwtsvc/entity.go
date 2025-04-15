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
	Role   string `json:"role,omitempty"`
}
type ClaimType string

const (
	AccessType ClaimType = "AccessType"
)

func (c *AccessClaims) HasRole(role string) bool {
	return c.Role == role
}

func (c *AccessClaims) HasAnyRole(roles ...string) bool {
	if c.Role == "" {
		return false
	}

	for _, r := range roles {
		if c.Role == r {
			return true
		}
	}

	return false
}
