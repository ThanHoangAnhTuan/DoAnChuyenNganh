package uploader

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

func SaveImageToDisk(ctx *gin.Context, images []*multipart.FileHeader) (codeStatus int, savedImagePaths []string, err error) {
	uploadDir := "storage/uploads"

	// TODO: Make sure the directory exists
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return response.ErrCodeCreateFolderFailed, nil, fmt.Errorf("cannot create upload directory: %s", err)
		}
	}

	for _, file := range images {
		// TODO: create unique file name
		fileName := strconv.Itoa(int(utiltime.GetTimeNow())) + uuid.New().String()
		fileName += filepath.Ext(file.Filename)

		// TODO: create path
		savePath := filepath.Join(uploadDir, fileName)

		// TODO: save file
		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			return response.ErrCodeUploadFileFailed, nil, fmt.Errorf("error upload images: %s", err)
		}
		savedImagePaths = append(savedImagePaths, fileName)
	}

	return response.ErrCodeUploadFileSuccess, savedImagePaths, nil
}

func DeleteImageToDisk(fileNames []string) error {
	uploadDir := "storage/uploads"
	var failedDeletes []string

	for _, name := range fileNames {
		imagePath := filepath.Join(uploadDir, name)
		if err := os.Remove(imagePath); err != nil {
			failedDeletes = append(failedDeletes, name)
			continue
		}
	}

	if len(failedDeletes) > 0 {
		return fmt.Errorf("failed to delete images: %v", failedDeletes)
	}

	return nil
}
