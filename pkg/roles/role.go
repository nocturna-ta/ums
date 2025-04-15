package roles

const (
	RoleVoter       = "voter"
	RoleKPUKota     = "kpu_kota"
	RoleKPUProvinsi = "kpu_provinsi"
	RoleKPUPusat    = "kpu_pusat"
)

func IsValidRole(role string) bool {
	switch role {
	case RoleVoter, RoleKPUKota, RoleKPUProvinsi, RoleKPUPusat:
		return true
	default:
		return false
	}
}

func GetRoleDisplayName(role string) string {
	switch role {
	case RoleVoter:
		return "Voter"
	case RoleKPUKota:
		return "KPU Kota"
	case RoleKPUProvinsi:
		return "KPU Provinsi"
	case RoleKPUPusat:
		return "KPU Pusat"
	default:
		return "Unknown"
	}
}
