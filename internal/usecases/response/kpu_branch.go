package response

type KPUBranchResponse struct {
	ID            string `json:"id"`
	BranchAddress string `json:"branch_address"`
	Region        string `json:"region"`
	IsActive      bool   `json:"is_active"`
}

type KPUBranchRegistrationResponse struct {
	BranchAddress string `json:"branch_address"`
	IsActive      bool   `json:"is_active"`
}
