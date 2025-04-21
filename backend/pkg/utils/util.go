package utils

import (
	"fmt"
	"strings"

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
