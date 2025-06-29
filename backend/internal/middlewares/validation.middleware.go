package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/validation"
	"go.uber.org/zap"
)

func ValidatorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validate := validator.New()
		if err := validate.RegisterValidation("strongpassword", validation.ValidateStrongPassword); err != nil {
			// Xử lý lỗi nếu không đăng ký được
			global.Logger.Error("Failed to register strongpassword validation", zap.Error(err))
			response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
			return
		}
		ctx.Set("validation", validate)
		ctx.Next()
	}
}
