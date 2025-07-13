package admin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"go.opentelemetry.io/otel/attribute"
)

type CAdminLogin struct{}

func (c *CAdminLogin) Register(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "AdminRegister",
		attribute.String("operation", "register"),
		attribute.String("resource", "admin"),
	)
	defer span.End()

	var params vo.AdminRegisterInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.AdminRegisterInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.AdminLogin().Register(ctx, &params)
	duration := time.Since(start)

	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	span.SetAttributes(
		attribute.Int("status_code", codeStatus),
		attribute.Int64("duration_ms", duration.Milliseconds()),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CAdminLogin) Login(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "AdminLogin",
		attribute.String("operation", "login"),
		attribute.String("resource", "admin"),
	)
	defer span.End()

	var params vo.AdminLoginInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.AdminLoginInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.AdminLogin().Login(ctx, &params)
	duration := time.Since(start)

	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	span.SetAttributes(
		attribute.Int("status_code", codeStatus),
		attribute.Int64("duration_ms", duration.Milliseconds()),
	)

	logData := map[string]interface{}{
		"userAccount": params.UserAccount,
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}
