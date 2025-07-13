package user

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

type CUserLogin struct {
}

func (c *CUserLogin) Register(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "Register",
		attribute.String("operation", "register"),
		attribute.String("resource", "user"),
	)
	defer span.End()

	var params vo.RegisterInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.RegisterInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.UserLogin().Register(ctx, &params)
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
		attribute.String("email", params.VerifyKey),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CUserLogin) VerifyOTP(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "VerifyOTP",
		attribute.String("operation", "verify"),
		attribute.String("resource", "user"),
	)
	defer span.End()

	var params vo.VerifyOTPInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.VerifyOTPInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.UserLogin().VerifyOTP(ctx, &params)
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
		attribute.String("email", params.VerifyKey),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "UpdatePasswordRegister",
		attribute.String("operation", "update_password"),
		attribute.String("resource", "user"),
	)
	defer span.End()

	var params vo.UpdatePasswordRegisterInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdatePasswordRegisterInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.UserLogin().UpdatePasswordRegister(ctx, &params)
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
		attribute.String("token", params.Token),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CUserLogin) Login(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "Login",
		attribute.String("operation", "login"),
		attribute.String("resource", "user"),
	)
	defer span.End()

	var params vo.LoginInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.LoginInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.UserLogin().Login(ctx, &params)
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
		attribute.String("email", params.UserAccount),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}
