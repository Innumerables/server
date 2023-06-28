package upload

import (
	"mime/multipart"
	"server/global"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewOss() OSS {
	switch global.GVA_CONFIG.System.OssType {
	case "locla":
		return &Local{}
	default:
		return &Local{}
	}
}
