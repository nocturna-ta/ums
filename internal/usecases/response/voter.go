package response

type VoterResponse struct {
	ID           string `json:"id"`
	NIK          string `json:"nik"`
	VoterAddress string `json:"voter_address"`
	Region       string `json:"region"`
	IsRegistered bool   `json:"is_registered"`
	HasVoted     bool   `json:"has_voted"`
}

type VoterRegistrationResponse struct {
	IsRegistered bool `json:"is_registered"`
}
