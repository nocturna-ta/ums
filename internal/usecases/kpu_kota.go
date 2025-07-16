package usecases

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/http"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"io"
)

type KPUKotaUseCases interface {
	RegisterKPUKota(ctx context.Context, req *request.KPUKotaRegistrationRequest) (*response.KPUKotaRegistrationResponse, error)
	GetAllKPUKota(ctx context.Context) (*[]response.KPUKotaResponse, error)
	GetKPUKotaByAddress(ctx context.Context) (*response.KPUKotaResponse, error)
	GetKPUKotaByID(ctx context.Context, id uuid.UUID) (*response.KPUKotaResponse, error)
	UploadKPUKotaPhoto(ctx context.Context, fileData io.Reader, fileName string) error
	GetKPUKotaPhoto(ctx context.Context) (*http.File, string, error)
	GetKPUKotaPhotoUseID(ctx context.Context, id uuid.UUID) (*http.File, string, error)
	UpdateKPUKota(ctx context.Context, updateRequest *request.KPUKotaUpdateRequest) (*response.KPUKotaResponse, error)
	GetKPUKotaByUserID(ctx context.Context) (*response.KPUKotaResponse, error)
}
