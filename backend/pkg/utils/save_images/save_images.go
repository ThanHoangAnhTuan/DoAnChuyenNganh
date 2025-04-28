package saveimages

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
)

func SaveImages(ctx *gin.Context, images []*multipart.FileHeader) (savedImagePaths []string, err error) {
	uploadDir := "storage/uploads"

	for _, file := range images {
		// !1. create file name
		fileName := strconv.Itoa(int(utiltime.GetTimeNow())) + uuid.New().String()
		fileName += filepath.Ext(file.Filename)

		// !2. create file path
		savePath := filepath.Join(uploadDir, fileName)

		// !3. save file to disk
		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			return nil, fmt.Errorf("error for save image: %s", err)
		}
		savedImagePaths = append(savedImagePaths, fileName)
	}

	fmt.Println("savedImagePaths: ", savedImagePaths)

	return savedImagePaths, nil
}
