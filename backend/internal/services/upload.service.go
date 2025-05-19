package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
)

type (
	IUpload interface {
		UploadImages(ctx *gin.Context, in *vo.UploadImages) (codeStatus int, savedImagePaths []string, err error)
		GetImages(ctx *gin.Context, in *vo.GetImagesInput) (codeStatus int, imagesPath []string, err error)
		DeleteImage(ctx *gin.Context, fileName string) (err error)
	}
)

var (
	localUpload IUpload
)

func Upload() IUpload {
	if localUpload == nil {
		panic("Implement localUpload not found for interface IUpload")
	}
	return localUpload
}

func InitUpload(i IUpload) {
	localUpload = i
}
