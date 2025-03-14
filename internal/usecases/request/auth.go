package request

type ValidateAuthorizationRequest struct {
	Headers       map[string]string
	Path          string
	TargetService string
}
