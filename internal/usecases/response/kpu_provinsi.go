package response

type KPUProvinsiResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Region   string `json:"region"`
	IsActive bool   `json:"is_active"`
	PhotoURL string `json:"photo_url"`
}

type KPUProvinsiRegistrationResponse struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
}
