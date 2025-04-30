package response

type VoterResponse struct {
	ID                 string `json:"id"`
	UserID             string `json:"user_id"`
	NIK                string `json:"nik"`
	FullName           string `json:"full_name"`
	Gender             string `json:"gender"`
	BirthPlace         string `json:"birth_place"`
	BirthDate          string `json:"birth_date"`
	ResidentialAddress string `json:"residential_address"`
	VoterAddress       string `json:"voter_address"`
	Region             string `json:"region"`
	IsRegistered       bool   `json:"is_registered"`
	HasVoted           bool   `json:"has_voted"`
}

type VoterRegistrationResponse struct {
	IsRegistered bool `json:"is_registered"`
}
