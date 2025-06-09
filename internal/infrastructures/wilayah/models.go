package wilayah

type WilayahMeta struct {
	AdministrativeAreaLevel int    `json:"administrative_area_level"`
	UpdatedAt               string `json:"updated_at"`
}

// ProvinceData represents province data from wilayah.id
type ProvinceData struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// RegencyData represents regency/city data from wilayah.id
type RegencyData struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// ProvinceAPIResponse represents the provinces response
type ProvinceAPIResponse struct {
	Data []ProvinceData `json:"data"`
	Meta WilayahMeta    `json:"meta"`
}

// RegencyAPIResponse represents the regencies response
type RegencyAPIResponse struct {
	Data []RegencyData `json:"data"`
	Meta WilayahMeta   `json:"meta"`
}
