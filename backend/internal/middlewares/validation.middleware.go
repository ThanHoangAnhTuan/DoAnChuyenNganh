package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidatorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validate := validator.New()
		ctx.Set("validation", validate)
		ctx.Next()
	}
}
