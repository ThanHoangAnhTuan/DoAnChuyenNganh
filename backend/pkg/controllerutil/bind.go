package controllerutil

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/tracing"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"go.uber.org/zap"
)

type ValidationError struct {
	Type    string      `json:"type"`              // "internal_error", "binding_error", "validation_error"
	Message string      `json:"message"`           // Human readable error message
	Details interface{} `json:"details,omitempty"` // Additional error details (e.g., field validation errors)
}

func BindAndValidate[T any](ctx *gin.Context, params *T, bindFunc func(*T) error) *ValidationError {
	validation, exists := ctx.Get("validation")
	if !exists {
		HandleStructuredLog(ctx, nil, "error", response.ErrCodeValidator, 0, fmt.Errorf("Validation middleware not found"))
		return &ValidationError{
			Type:    "internal_error",
			Message: "Validation middleware not found",
		}
	}

	if err := bindFunc(params); err != nil {
		HandleStructuredLog(ctx, nil, "error", response.ErrCodeValidator, 0, fmt.Errorf("Invalid request format: %s", err.Error()))
		return &ValidationError{
			Type:    "binding_error",
			Message: "Invalid request format",
		}
	}

	if err := validation.(*validator.Validate).Struct(params); err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		HandleStructuredLog(ctx, validationErrors, "error", response.ErrCodeValidator, 0, fmt.Errorf("Validation failed"))
		return &ValidationError{
			Type:    "validation_error",
			Message: "Validation failed",
			Details: validationErrors,
		}
	}

	return nil
}

func HandleValidationError(ctx *gin.Context, params interface{}, validationErr *ValidationError, duration time.Duration) {
	HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, duration, fmt.Errorf(validationErr.Message))
	response.ErrorResponse(ctx, response.ErrCodeValidator, validationErr)
}

func HandleStructuredLog(ctx *gin.Context, params interface{}, status string, code int, duration time.Duration, err error) {
	userId, _ := utils.GetUserIDFromGin(ctx)

	// Get method and path from context
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	api := method + " " + path

	// Get trace information
	traceID := tracing.TraceIDFromContext(ctx.Request.Context())
	spanID := tracing.SpanIDFromContext(ctx.Request.Context())

	fields := []zap.Field{
		zap.String("timestamp", utils.GetCurrentUTCTimestamp()),
		zap.String("api", api),
		zap.String("userId", userId),
		zap.Any("params", params),
		zap.String("status", status),
		zap.Int("code", code),
		zap.Int64("durationMs", duration.Milliseconds()),
	}

	// Add trace information if available
	if traceID != "" {
		fields = append(fields, zap.String("trace_id", traceID))
	}
	if spanID != "" {
		fields = append(fields, zap.String("span_id", spanID))
	}

	if err != nil {
		fields = append(fields, zap.String("error", err.Error()))
		global.Logger.Error("API Log", fields...)
	} else {
		global.Logger.Info("API Log", fields...)
	}
}
