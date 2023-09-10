package ctrl

import (
	"im/apps/upload/internal/service"
)

type UploadCtrl struct {
	svc service.UploadService
}

func NewUploadCtrl(svc service.UploadService) *UploadCtrl {
	return &UploadCtrl{svc: svc}
}
