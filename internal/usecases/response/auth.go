package response

type ValidateAuthorizationResponse struct {
	IsValid       bool
	ExplodeHeader map[string]string
}
