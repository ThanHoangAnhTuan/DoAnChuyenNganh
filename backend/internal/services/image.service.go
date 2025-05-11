package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
)

type (
	IUploadImage interface {
		UploadImages(ctx *gin.Context, in *vo.UploadImages) (codeStatus int, savedImagePaths []string, err error)
		GetImages(ctx *gin.Context, in *vo.GetImagesInput) (codeStatus int, imagesPath []string, err error)
		DeleteImage(ctx *gin.Context, fileName string) (err error)
	}
)

var (
	localUploadImage IUploadImage
)

func UploadImage() IUploadImage {
	if localUploadImage == nil {
		panic("Implement localUploadImage not found for interface IUploadImage")
	}
	return localUploadImage
}

func InitImage(i IUploadImage) {
	localUploadImage = i
}
