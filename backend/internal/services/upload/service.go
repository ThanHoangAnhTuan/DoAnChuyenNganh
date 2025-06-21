package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	UploadImages(ctx *gin.Context, in *vo.UploadImages) (codeStatus int, savedImagePaths []string, err error)
	GetImages(ctx *gin.Context, in *vo.GetImagesInput) (codeStatus int, imagesPath []string, err error)
	DeleteImage(ctx *gin.Context, fileName string) (err error)
}
