package facility

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

func (c *Controller) CreateFacility(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "CreateFacility",
		attribute.String("operation", "create"),
		attribute.String("resource", "facility"),
	)
	defer span.End()

	var params vo.CreateFacilityInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateFacilityInput) error {
		return ctx.ShouldBindQuery(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Facility().CreateFacility(ctx, &params)
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
		attribute.String("facility.id", data.ID),
		attribute.String("facility.name", data.Name),
	)

	logData := map[string]interface{}{
		"facilityId": data.ID,
		"name":       data.Name,
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) UpdateFacility(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "UpdateFacility",
		attribute.String("operation", "update"),
		attribute.String("resource", "facility"),
	)
	defer span.End()

	var params vo.UpdateFacilityInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdateFacilityInput) error {
		return ctx.ShouldBindQuery(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Facility().UpdateFacility(ctx, &params)
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
		attribute.String("facility.id", data.ID),
		attribute.String("facility.name", data.Name),
	)

	logData := map[string]interface{}{
		"facilityId": data.ID,
		"name":       data.Name,
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteFacility(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "DeleteFacility",
		attribute.String("operation", "delete"),
		attribute.String("resource", "facility"),
	)
	defer span.End()

	var params vo.DeleteFacilityInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.DeleteFacilityInput) error {
		return ctx.ShouldBindUri(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.Facility().DeleteFacility(ctx, &params)
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
		attribute.String("facility.id", params.ID),
	)

	logData := map[string]interface{}{
		"facilityId": params.ID,
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) GetFacilities(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetFacilities",
		attribute.String("operation", "list"),
		attribute.String("resource", "facility"),
	)
	defer span.End()

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Facility().GetFacilities(ctx)
	duration := time.Since(start)

	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
		controllerutil.HandleStructuredLog(ctx, nil, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	span.SetAttributes(
		attribute.Int("status_code", codeStatus),
		attribute.Int64("duration_ms", duration.Milliseconds()),
		attribute.Int("facilities.count", len(data)),
	)

	controllerutil.HandleStructuredLog(ctx, nil, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}
