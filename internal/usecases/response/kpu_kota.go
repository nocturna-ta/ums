package response

type KPUKotaResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Address      string `json:"address"`
	Name         string `json:"name"`
	Region       string `json:"region"`
	IsActive     bool   `json:"is_active"`
	PhotoURL     string `json:"photo_url"`
	Telephone    string `json:"telephone"`
	RegisteredAt string `json:"registered_at"`
}

type KPUKotaRegistrationResponse struct {
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}
