package request

import (
	"encoding/json"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/pkg/utils"
	"mime/multipart"
)

func ParseRegistrationRequest(form *multipart.Form, files map[string]utils.UploadedFile) (*request.UserRegistrationRequest, error) {
	userValues := form.Value["user"]
	if len(userValues) == 0 {
		return nil, &custerr.ErrChain{
			Message: "Missing registration request data in 'user' field",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	userJSON := userValues[0]
	var regReq request.UserRegistrationRequest
	if err := json.Unmarshal([]byte(userJSON), &regReq); err != nil {
		return nil, &custerr.ErrChain{
			Message: "Invalid JSON in registration request",
			Code:    400,
			Type:    response.ErrBadRequest,
			Cause:   err,
		}
	}

	if err := regReq.ValidateRegistrationUser(); err != nil {
		return nil, err
	}

	if ktpPhoto, exists := files["ktp_photo"]; exists {
		regReq.KTPPhotoName = ktpPhoto.OriginalFilename
		regReq.KTPPhotoFile = ktpPhoto.File
	}

	return &regReq, nil
}
