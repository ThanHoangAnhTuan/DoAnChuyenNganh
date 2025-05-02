package impl

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
)

type ImageImpl struct {
}

// DeleteImage implements services.IImage.
func (i *ImageImpl) DeleteImage(ctx *gin.Context, fileName string) (err error) {
	panic("unimplemented")
}

func (i *ImageImpl) UploadImages(ctx *gin.Context, images []*multipart.FileHeader) (savedImagePaths []string, err error) {
	uploadDir := "storage/uploads"

	// Make sure the directory exists
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("cannot create upload directory: %s", err)
		}
	}

	for _, file := range images {
		// 1. create unique file name
		fileName := strconv.Itoa(int(utiltime.GetTimeNow())) + uuid.New().String()
		fileName += filepath.Ext(file.Filename)

		// 2. create path
		savePath := filepath.Join(uploadDir, fileName)

		// 3. save file
		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			return nil, fmt.Errorf("error saving image: %s", err)
		}

		savedImagePaths = append(savedImagePaths, fileName)
	}

	return savedImagePaths, nil
}

func NewImageImpl() *ImageImpl {
	return &ImageImpl{}
}
