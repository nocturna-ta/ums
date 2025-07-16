package controller

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/custerr"
	"github.com/nocturna-ta/golib/fileutils"
	"github.com/nocturna-ta/golib/http/filehandler"
	"github.com/nocturna-ta/golib/response"
	"github.com/nocturna-ta/golib/response/rest"
	"github.com/nocturna-ta/golib/router"
	"github.com/nocturna-ta/golib/tracing"
	"github.com/nocturna-ta/ums/internal/infrastructures/custresp"
	"github.com/nocturna-ta/ums/internal/usecases/request"
)

// GetAllKPUProvinsi godoc
// @Summary Get All KPU Provinsi
// @Description Get All KPU Provinsi
// @Tags kpu_provinsi
// @Param X-User-Id header string false "User"
// @Param X-Address header string false "Address"
// @Param X-Role header string false "Role"
// @Accept json
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUProvinsiResponse}
// @Router /v1/kpu-provinsi [get]
func (api *API) GetAllKPUProvinsi(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAllKPUProvinsi")
	defer span.End()

	res, err := api.kpuProvinsiUc.GetAllKPUProvinsi(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetKPUProvinsiByUserID godoc
// @Summary Get KPU Provinsi By User ID
// @Description Get KPU Provinsi By User ID
// @Tags kpu_provinsi
// @Accept json
// @Param X-User-Id header string true "User ID"
// @Param X-Address-Id header string false "Public Address"
// @Param X-Role header string false "Role"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUProvinsiResponse}
// @Router /v1/kpu-provinsi/id [get]
func (api *API) GetKPUProvinsiByUserID(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetKPUProvinsiByID")
	defer span.End()

	res, err := api.kpuProvinsiUc.GetKPUProvinsiByUserID(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// UploadKPUProvinsiPhoto godoc
// @Summary Upload photo for KPU Provinsi
// @Description Upload a photo for a KPU Provinsi
// @Tags kpu_provinsi
// @Accept multipart/form-data
// @Param X-User-Id header string true "User ID"
// @Param X-Address-Id header string false "Public Address"
// @Param X-Role header string false "Role"
// @Param photo formData file true "Photo file (jpg, jpeg, png only)"
// @Produce json
// @Success 200 {object} jsonResponse{data=string}
// @Router /v1/kpu-provinsi/photo [post]
func (api *API) UploadKPUProvinsiPhoto(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.UploadKPUProvinsiPhoto")
	defer span.End()

	form, err := req.RawRequest().MultipartForm()
	if err != nil {
		return custresp.CustomErrorResponse(&custerr.ErrChain{
			Message: "Failed to parse multipart form",
			Code:    400,
			Type:    response.ErrBadRequest,
			Cause:   err,
		})
	}

	uploadOptions := filehandler.ImageUploadOptions()
	uploadOptions.FieldName = "photo"

	uploadResult, err := filehandler.UploadFile(ctx, form, uploadOptions)
	if err != nil {
		var errorMsg string
		var errorCode int

		switch err {
		case filehandler.ErrNoFile:
			errorMsg = "No photo file provided"
			errorCode = 400
		case filehandler.ErrFileTooLarge:
			errorMsg = "Photo file size exceeds maximum allowed size"
			errorCode = 400
		case filehandler.ErrInvalidFileFormat:
			errorMsg = "Invalid file format. Only JPG, JPEG, and PNG files are allowed"
			errorCode = 400
		default:
			errorMsg = "Failed to process uploaded file"
			errorCode = 500
		}

		return custresp.CustomErrorResponse(&custerr.ErrChain{
			Message: errorMsg,
			Code:    errorCode,
			Type:    response.ErrBadRequest,
			Cause:   err,
		})
	}

	file, err := fileutils.OpenFile(ctx, uploadResult.FilePath)
	if err != nil {
		return custresp.CustomErrorResponse(&custerr.ErrChain{
			Message: "Failed to read uploaded file",
			Code:    500,
			Type:    response.ErrInternalServerError,
			Cause:   err,
		})
	}
	defer file.Close()

	err = api.kpuProvinsiUc.UploadKPUProvinsiPhoto(ctx, file, uploadResult.OriginalFilename)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData("Photo uploaded successfully"), nil
}

// GetKPUProvinsiPhoto godoc
// @Summary Get photo for KPU Provinsi
// @Description Get the photo for a KPU Provinsi
// @Tags kpu_provinsi
// @Param X-User-Id header string true "User ID"
// @Param X-Address-Id header string false "Public Address"
// @Param X-Role header string false "Role"
// @Produce octet-stream
// @Success 200
// @Router /v1/kpu-provinsi/photo [get]
func (api *API) GetKPUProvinsiPhoto(ctx context.Context, req *router.Request) (*rest.AttachmentResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetKPUProvinsiPhoto")
	defer span.End()

	file, contentType, err := api.kpuProvinsiUc.GetKPUProvinsiPhoto(ctx)
	if err != nil {
		return nil, err
	}

	return rest.NewAttachmentResponse().
		SetFile(file).
		SetFileName(file.FileName).
		SetContentType(contentType), nil
}

// GetKPUProvinsiPhotoUseID godoc
// @Summary Get photo for KPU Provinsi
// @Description Get the photo for a KPU Provinsi
// @Tags kpu_provinsi
// @Param X-User-Id header string true "User ID"
// @Param X-Address-Id header string false "Public Address"
// @Param X-Role header string false "Role"
// @Param id path string true "KPU Provinsi ID"
// @Produce octet-stream
// @Success 200
// @Router /v1/kpu-provinsi/photo/{id} [get]
func (api *API) GetKPUProvinsiPhotoUseID(ctx context.Context, req *router.Request) (*rest.AttachmentResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetKPUProvinsiPhoto")
	defer span.End()

	id := req.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	file, contentType, err := api.kpuProvinsiUc.GetKPUProvinsiPhotoUseID(ctx, uuidID)
	if err != nil {
		return nil, err
	}

	return rest.NewAttachmentResponse().
		SetFile(file).
		SetFileName(file.FileName).
		SetContentType(contentType), nil
}

// UpdateKPUProvinsi godoc
// @Summary Update KPU Provinsi
// @Description Update KPU Provinsi information
// @Tags kpu_provinsi
// @Accept json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string true "Address"
// @Param X-Role header string false "Role"
// @Param data body request.KPUProvinsiUpdateRequest true "Update Request"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUProvinsiResponse}
// @Router /v1/kpu-provinsi/update [put]
func (api *API) UpdateKPUProvinsi(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.UpdateKPUProvinsi")
	defer span.End()

	var updateReq request.KPUProvinsiUpdateRequest
	err := json.Unmarshal(req.RawBody(), &updateReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	err = updateReq.ValidateUpdateRequest()
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	res, err := api.kpuProvinsiUc.UpdateKPUProvinsi(ctx, &updateReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
