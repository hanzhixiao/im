package dig

import (
	"im/apps/upload/internal/service"
)

func provideUpload() {
	Provide(service.NewUploadService)
}
