package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/timezone"
)

func TimezoneMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tz := c.GetHeader("Timezone")
		if tz == "" {
			tz = c.Query("timezone")
		}
		if tz == "" {
			tz, _ = c.Cookie("timezone")
		}
		if tz == "" {
			tz = "UTC"
		}

		ctx := timezone.WithTimezone(c.Request.Context(), tz)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
