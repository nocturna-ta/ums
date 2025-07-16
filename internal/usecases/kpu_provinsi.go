package usecases

import (
	"context"
	"github.com/google/uuid"
	"github.com/nocturna-ta/golib/http"
	"github.com/nocturna-ta/ums/internal/usecases/request"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"io"
)

type KPUProvinsiUseCases interface {
	GetKPUPusatByUserID(ctx context.Context) (*response.KPUProvinsiResponse, error)
	RegisterKPUProvinsi(ctx context.Context, req *request.KPUProvinsiRegistrationRequest) (*response.KPUProvinsiRegistrationResponse, error)
	GetAllKPUProvinsi(ctx context.Context) (*[]response.KPUProvinsiResponse, error)
	GetKPUProvinsiByAddress(ctx context.Context) (*response.KPUProvinsiResponse, error)
	GetKPUProvinsiByID(ctx context.Context, id uuid.UUID) (*response.KPUProvinsiResponse, error)
	UploadKPUProvinsiPhoto(ctx context.Context, fileData io.Reader, fileName string) error
	GetKPUProvinsiPhoto(ctx context.Context) (*http.File, string, error)
	GetKPUProvinsiPhotoUseID(ctx context.Context, id uuid.UUID) (*http.File, string, error)
	UpdateKPUProvinsi(ctx context.Context, updateRequest *request.KPUProvinsiUpdateRequest) (*response.KPUProvinsiResponse, error)
	GetKPUProvinsiByUserID(ctx context.Context) (*response.KPUProvinsiResponse, error)
}
