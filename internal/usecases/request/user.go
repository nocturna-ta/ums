package request

type ChangeUserPasswordRequest struct {
	Old     string `json:"old"`
	New     string `json:"new"`
	Confirm string `json:"confirm"`
}

type UserLoginRequest struct {
	NIK      string `json:"nik"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	NIK  string `json:"nik"`
	Name string `json:"name"`
}
