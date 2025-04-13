package response

type KPUKotaResponse struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	Region   string `json:"region"`
	IsActive bool   `json:"is_active"`
}

type KPUKotaRegistrationResponse struct {
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}
