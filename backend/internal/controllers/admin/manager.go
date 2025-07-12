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

type CAdminManager struct{}

func (c *CAdminManager) GetManagers(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetManagers",
		attribute.String("operation", "list"),
		attribute.String("resource", "admin_manager"),
	)
	defer span.End()

	var params vo.GetManagerInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetManagerInput) error {
		return ctx.ShouldBindQuery(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, pagination, err := services.AdminManager().GetManagers(ctx, &params)
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
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}

func (c *CAdminManager) GetAccommodationsOfManager(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetAccommodationsOfManager",
		attribute.String("operation", "list_accommodations"),
		attribute.String("resource", "admin_manager"),
	)
	defer span.End()

	var params vo.GetAccommodationsOfManagerInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationsOfManagerInput) error {
		if err := ctx.ShouldBindUri(p); err != nil {
			return err
		}
		return ctx.ShouldBindQuery(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, pagination, err := services.AdminManager().GetAccommodationsOfManager(ctx, &params)
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
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}

func (c *CAdminManager) VerifyAccommodation(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "VerifyAccommodation",
		attribute.String("operation", "verify"),
		attribute.String("resource", "admin_manager"),
	)
	defer span.End()

	var params vo.VerifyAccommodationInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.VerifyAccommodationInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.AdminManager().VerifyAccommodation(ctx, &params)
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

func (c *CAdminManager) SetDeletedAccommodation(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "SetDeletedAccommodation",
		attribute.String("operation", "set_deleted"),
		attribute.String("resource", "admin_manager"),
	)
	defer span.End()

	var params vo.SetDeletedAccommodationInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.SetDeletedAccommodationInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.AdminManager().SetDeletedAccommodation(ctx, &params)
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
		"accommodationId": params.AccommodationID,
	}

	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}
