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

// RegisterKPUKota godoc
// @Summary Register KPU Kota
// @Description Register KPU Kota
// @Tags kpu_kota
// @Accept json
// @Param users body request.KPUKotaRegistrationRequest true "Register Request"
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUKotaRegistrationResponse}
// @Router /v1/kpu-kota/register [post]
func (api *API) RegisterKPUKota(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.RegisterKPUKota")
	defer span.End()

	var regisReq request.KPUKotaRegistrationRequest
	err := json.Unmarshal(req.RawBody(), &regisReq)

	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	err = regisReq.ValidateRegistrationRequest()

	res, err := api.kpuKotaUc.RegisterKPUKota(ctx, &regisReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetAllKPUKota godoc
// @Summary Get All KPU Kota
// @Description Get All KPU Kota
// @Tags kpu_kota
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Accept json
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUKotaResponse}
// @Router /v1/kpu-kota [get]
func (api *API) GetAllKPUKota(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetAllKPUKota")
	defer span.End()

	res, err := api.kpuKotaUc.GetAllKPUKota(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetKPUKotaByAddress godoc
// @Summary Get KPU Kota By Address
// @Description Get KPU Kota By Address
// @Tags kpu_kota
// @Accept json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUKotaResponse}
// @Router /v1/kpu-kota/address [get]
func (api *API) GetKPUKotaByAddress(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetKPUKotaByAddress")
	defer span.End()

	res, err := api.kpuKotaUc.GetKPUKotaByAddress(ctx)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// GetKPUKotaByID godoc
// @Summary Get KPU Kota By ID
// @Description Get KPU Kota By ID
// @Tags kpu_kota
// @Accept json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Param id path string true "KPU Kota ID"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUKotaResponse}
// @Router /v1/kpu-kota/{id} [get]
func (api *API) GetKPUKotaByID(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetKPUKotaByID")
	defer span.End()

	kpuKotaIDStr := req.Params("id")
	kpuKotaID, err := uuid.Parse(kpuKotaIDStr)
	if err != nil {
		return custresp.CustomErrorResponse(&custerr.ErrChain{
			Message: "Invalid KPU Kota ID",
			Code:    400,
			Type:    response.ErrBadRequest,
		})
	}

	res, err := api.kpuKotaUc.GetKPUKotaByID(ctx, kpuKotaID)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}

// UploadKPUKotaPhoto godoc
// @Summary Upload photo for KPU Kota
// @Description Upload a photo for a KPU Kota
// @Tags kpu_kota
// @Accept multipart/form-data
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Param id path string true "KPU Kota ID"
// @Param photo formData file true "Photo file (jpg, jpeg, png only)"
// @Produce json
// @Success 200 {object} jsonResponse{data=string}
// @Router /v1/kpu-kota/{id}/photo [post]
func (api *API) UploadKPUKotaPhoto(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.UploadKPUKotaPhoto")
	defer span.End()

	kpuKotaIDStr := req.Params("id")
	kpuKotaID, err := uuid.Parse(kpuKotaIDStr)
	if err != nil {
		return custresp.CustomErrorResponse(&custerr.ErrChain{
			Message: "Invalid KPU Kota ID",
			Code:    400,
			Type:    response.ErrBadRequest,
		})
	}

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

	err = api.kpuKotaUc.UploadKPUKotaPhoto(ctx, kpuKotaID, file, uploadResult.OriginalFilename)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData("Photo uploaded successfully"), nil
}

// GetKPUKotaPhoto godoc
// @Summary Get photo for KPU Kota
// @Description Get the photo for a KPU Kota
// @Tags kpu_kota
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string false "Address"
// @Param X-Role header string false "Role"
// @Param id path string true "KPU Kota ID"
// @Produce octet-stream
// @Success 200
// @Router /v1/kpu-kota/{id}/photo [get]
func (api *API) GetKPUKotaPhoto(ctx context.Context, req *router.Request) (*rest.AttachmentResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.GetKPUKotaPhoto")
	defer span.End()

	kpuKotaIDStr := req.Params("id")
	kpuKotaID, err := uuid.Parse(kpuKotaIDStr)
	if err != nil {
		return nil, &custerr.ErrChain{
			Message: "Invalid KPU Kota ID",
			Code:    400,
			Type:    response.ErrBadRequest,
		}
	}

	file, contentType, err := api.kpuKotaUc.GetKPUKotaPhoto(ctx, kpuKotaID)
	if err != nil {
		return nil, err
	}

	return rest.NewAttachmentResponse().
		SetFile(file).
		SetFileName(file.FileName).
		SetContentType(contentType), nil
}

// UpdateKPUKota godoc
// @Summary Update KPU Kota
// @Description Update KPU Kota information
// @Tags kpu_kota
// @Accept json
// @Param X-User-Id header string false "User"
// @Param X-Address-Id header string true "Address"
// @Param X-Role header string false "Role"
// @Param data body request.KPUKotaUpdateRequest true "Update Request"
// @Produce json
// @Success 200 {object} jsonResponse{data=response.KPUKotaResponse}
// @Router /v1/kpu-kota/update [put]
func (api *API) UpdateKPUKota(ctx context.Context, req *router.Request) (*rest.JSONResponse, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "Controller.UpdateKPUKota")
	defer span.End()

	var updateReq request.KPUKotaUpdateRequest
	err := json.Unmarshal(req.RawBody(), &updateReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	err = updateReq.ValidateUpdateRequest()
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	res, err := api.kpuKotaUc.UpdateKPUKota(ctx, &updateReq)
	if err != nil {
		return custresp.CustomErrorResponse(err)
	}

	return rest.NewJSONResponse().SetData(res), nil
}
