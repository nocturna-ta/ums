package response

type KPUKotaResponse struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Region   string `json:"region"`
	IsActive bool   `json:"is_active"`
	PhotoURL string `json:"photo_url"`
}

type KPUKotaRegistrationResponse struct {
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}
