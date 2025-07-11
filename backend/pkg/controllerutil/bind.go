package controllerutil

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"go.uber.org/zap"
)

// Helper function for binding and validating
func BindAndValidate[T any](ctx *gin.Context, params *T, bindFunc func(*T) error) error {
	validation, exists := ctx.Get("validation")
	if !exists {
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return fmt.Errorf("validation not found")
	}
	if err := bindFunc(params); err != nil {
		global.Logger.Error("Binding error", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return err
	}
	if err := validation.(*validator.Validate).Struct(params); err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		global.Logger.Error("Validation error", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return err
	}
	return nil
}

func HandleStructuredLog(ctx *gin.Context, params interface{}, status string, code int, durationMs int64, err error) {
	userId, _ := utils.GetUserIDFromGin(ctx)
	requestId := ctx.GetHeader("X-Request-Id")
	if requestId == "" {
		requestId = utils.GenerateRequestID()
	}

	// Get method and path from context
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	api := method + " " + path

	fields := []zap.Field{
		zap.String("timestamp", utils.GetCurrentUTCTimestamp()),
		zap.String("api", api),
		zap.String("userId", userId),
		zap.String("requestId", requestId),
		zap.Any("params", params),
		zap.String("status", status),
		zap.Int("code", code),
		zap.Int64("durationMs", durationMs),
	}

	if err != nil {
		fields = append(fields, zap.String("error", err.Error()))
		global.Logger.Error("API Log", fields...)
	} else {
		global.Logger.Info("API Log", fields...)
	}
}
