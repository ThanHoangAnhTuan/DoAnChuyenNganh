package accommodation

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

func (c *Controller) CreateAccommodation(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "CreateAccommodation",
		attribute.String("operation", "create"),
		attribute.String("resource", "accommodation"),
	)
	defer span.End()

	var params vo.CreateAccommodationInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateAccommodationInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Accommodation().CreateAccommodation(ctx, &params)
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
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodations(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetAccommodations",
		attribute.String("operation", "list"),
		attribute.String("resource", "accommodation"),
	)
	defer span.End()

	var params vo.GetAccommodationsInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationsInput) error {
		return ctx.ShouldBindQuery(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, pagination, err := services.Accommodation().GetAccommodations(ctx, &params)
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

func (c *Controller) UpdateAccommodation(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "UpdateAccommodation",
		attribute.String("operation", "update"),
		attribute.String("resource", "accommodation"),
	)
	defer span.End()

	var params vo.UpdateAccommodationInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdateAccommodationInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Accommodation().UpdateAccommodation(ctx, &params)
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
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteAccommodation(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "DeleteAccommodation",
		attribute.String("operation", "delete"),
		attribute.String("resource", "accommodation"),
	)
	defer span.End()

	var params vo.DeleteAccommodationInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.DeleteAccommodationInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.Accommodation().DeleteAccommodation(ctx, &params)
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

func (c *Controller) GetAccommodation(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetAccommodation",
		attribute.String("operation", "get"),
		attribute.String("resource", "accommodation"),
	)
	defer span.End()

	var params vo.GetAccommodationInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationInput) error {
		return ctx.ShouldBindUri(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Accommodation().GetAccommodation(ctx, &params)
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
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodationsByManager(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetAccommodationsByManager",
		attribute.String("operation", "list_by_manager"),
		attribute.String("resource", "accommodation"),
	)
	defer span.End()

	var params vo.GetAccommodationsInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationsInput) error {
		return ctx.ShouldBindQuery(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, pagination, err := services.Accommodation().GetAccommodationsByManager(ctx, &params)
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
