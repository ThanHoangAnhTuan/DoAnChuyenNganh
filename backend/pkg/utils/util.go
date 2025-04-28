package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GenerateCliTokenUUID(userId string) string {
	newUUID := uuid.New()

	uuidString := strings.ReplaceAll(newUUID.String(), "", "")
	return userId + "clitoken" + uuidString
}

func SaveImagesToLocal(ctx *gin.Context, files *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(files.Filename)
	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	err := ctx.SaveUploadedFile(files, "./uploads/"+newFilename)
	if err != nil {
		return "", err
	}
	return newFilename, nil
}
