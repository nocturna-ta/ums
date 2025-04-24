package response

import (
	"time"
)

type UserRegistrationResponse struct {
	ID                 string `json:"id,omitempty"`
	Email              string `json:"email,omitempty"`
	Username           string `json:"username,omitempty"`
	VerificationStatus string `json:"verification_status,omitempty"`
	RequestedRole      string `json:"requested_role,omitempty"`
	Message            string `json:"message,omitempty"`
}

type UserResponse struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type UserLoginResponse struct {
	Token              *string    `json:"token,omitempty"`
	ExpiresAt          *time.Time `json:"expires_at,omitempty"`
	VerificationStatus string     `json:"verification_status,omitempty"`
	IsActive           bool       `json:"is_active"`
	RequestedRole      string     `json:"requested_role,omitempty"`
	Message            string     `json:"message,omitempty"`
}

type UserVerificationResponse struct {
	ID                 string    `json:"id"`
	Email              string    `json:"email"`
	Username           string    `json:"username"`
	RequestedRole      string    `json:"requested_role"`
	VerificationStatus string    `json:"verification_status"`
	CreatedAt          time.Time `json:"created_at"`
}

type UserVerificationDetailsResponse struct {
	ID                 string                 `json:"id"`
	Email              string                 `json:"email"`
	Username           string                 `json:"username"`
	RequestedRole      string                 `json:"requested_role"`
	VerificationStatus string                 `json:"verification_status"`
	CreatedAt          time.Time              `json:"created_at"`
	EntityData         map[string]interface{} `json:"entity_data,omitempty"`
	SignedTransaction  string                 `json:"signed_transaction,omitempty"`
}

type UserVerificationStatusResponse struct {
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	RequestedRole      string    `json:"requested_role"`
	Role               string    `json:"role"`
	VerificationStatus string    `json:"verification_status"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`
	Message            string    `json:"message,omitempty"`
}

type EnhancedUserVerificationStatusResponse struct {
	Username           string    `json:"username"`
	Email              string    `json:"email"`
	RequestedRole      string    `json:"requested_role"`
	Role               string    `json:"role"`
	VerificationStatus string    `json:"verification_status"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`
	Message            string    `json:"message,omitempty"`
	VerifierRole       string    `json:"verifier_role,omitempty"`
	HierarchyLevel     int       `json:"hierarchy_level"`
}
