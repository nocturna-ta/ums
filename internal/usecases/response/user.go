package response

import "time"

type UserResponse struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	NIK         string `json:"nik,omitempty"`
	NoTelephone string `json:"no_telephone,omitempty"`
	Email       string `json:"email,omitempty"`
}

type UserRegistrationResponse struct {
	NIK string `json:"nik,omitempty"`
}

type UserLoginResponse struct {
	Token     *string    `json:"token"`
	ExpiresAt *time.Time `json:"expires_at"`
}
