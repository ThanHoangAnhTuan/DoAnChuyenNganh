package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GenerateCliTokenUUID(userId string) string {
	newUUID := uuid.New()

	uuidString := strings.ReplaceAll(newUUID.String(), "", "")
	return userId + "clitoken" + uuidString
}

func GetUserIDFromGin(ctx *gin.Context) (string, bool) {
	id, ok := ctx.Get("userId")
	if !ok {
		return "", false
	}
	uid, ok := id.(string)
	return uid, ok
}

// func GetUserIDFromContext(c context.Context) (string, bool) {
// 	id := c.Value("userId")
// 	uid, ok := id.(string)
// 	return uid, ok
// }

func FormatCurrency(amount uint32) string {
	p := message.NewPrinter(language.Vietnamese)
	return p.Sprintf("%d", amount)
}

func GetCurrentUTCTimestamp() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func NowMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GenerateRequestID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
