package response

import "time"

type UserRegistrationResponse struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
}

type UserResponse struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

type UserLoginResponse struct {
	Token     *string    `json:"token"`
	ExpiresAt *time.Time `json:"expires_at"`
}
