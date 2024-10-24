package response

type UserResponse struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	NIK         string `json:"nik,omitempty"`
	NoTelephone string `json:"no_telephone,omitempty"`
	Email       string `json:"email,omitempty"`
}
