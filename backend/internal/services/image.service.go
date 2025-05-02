package services

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type (
	IImage interface {
		UploadImages(ctx *gin.Context, images []*multipart.FileHeader) (savedImagePaths []string, err error)
		DeleteImage(ctx *gin.Context, fileName string) (err error)
	}
)

var (
	localImage IImage
)

func Image() IImage {
	if localImage == nil {
		panic("Implement localImage not found for interface IImage")
	}
	return localImage
}

func InitImage(i IImage) {
	localImage = i
}
