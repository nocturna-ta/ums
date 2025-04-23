package response

type KPUProvinsiResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Region       string `json:"region"`
	IsActive     bool   `json:"is_active"`
	PhotoURL     string `json:"photo_url"`
	Telephone    string `json:"telephone"`
	RegisteredAt string `json:"registered_at"`
}

type KPUProvinsiRegistrationResponse struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}
