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

func CanVerify(verifierRole, requestedRole string) bool {
	switch verifierRole {
	case RoleKPUPusat:
		return requestedRole == RoleKPUProvinsi
	case RoleKPUProvinsi:
		return requestedRole == RoleKPUKota
	case RoleKPUKota:
		return requestedRole == RoleVoter
	default:
		return false
	}
}

func GetVerifierRoleFor(role string) string {
	switch role {
	case RoleKPUProvinsi:
		return RoleKPUPusat
	case RoleKPUKota:
		return RoleKPUProvinsi
	case RoleVoter:
		return RoleKPUKota
	default:
		return ""
	}
}

func GetRoleLevel(role string) int {
	switch role {
	case RoleKPUPusat:
		return 4
	case RoleKPUProvinsi:
		return 3
	case RoleKPUKota:
		return 2
	case RoleVoter:
		return 1
	default:
		return 0
	}
}
