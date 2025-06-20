package response

type TotalDPTResponse struct {
	TotalDPT    int `json:"total_dpt"`
	PendingDPT  int `json:"pending_dpt"`
	ApprovedDPT int `json:"approved_dpt"`
	RejectedDPT int `json:"rejected_dpt"`
}
type ApprovedDPTResponse struct {
	Percentage       float64 `json:"percentage"`
	TotalApprovedDPT int     `json:"total_approved_dpt"`
}

type RejectedDPTResponse struct {
	Percentage       float64 `json:"percentage"`
	TotalRejectedDPT int     `json:"total_rejected_dpt"`
}

type PendingDPTResponse struct {
	Percentage      float64 `json:"percentage"`
	TotalPendingDPT int     `json:"total_pending_dpt"`
}

type StaffKPUResponse struct {
	Count       int `json:"count"`
	RegionCount int `json:"region_count"`
}

type DPTInformationResponse struct {
	KPURegion          string  `json:"kpu_region"`
	StaffCount         int     `json:"staff_count"`
	DPTVotedPercentage float64 `json:"dpt_voted_percentage"`
	TotalDPT           int     `json:"total_dpt"`
}

type VotedStatisticResponse struct {
	Percentage float64 `json:"percentage"`
	TotalDPT   int     `json:"total_dpt"`
}
