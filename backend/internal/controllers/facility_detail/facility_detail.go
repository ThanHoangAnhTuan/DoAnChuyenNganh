package facility_detail

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

func (c *Controller) CreateFacilityDetail(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "CreateFacilityDetail",
		attribute.String("operation", "create"),
		attribute.String("resource", "facility_detail"),
	)
	defer span.End()

	var params vo.CreateFacilityDetailInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateFacilityDetailInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.FacilityDetail().CreateFacilityDetail(ctx, &params)
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
		attribute.String("facility_detail.id", data.ID),
	)

	logData := map[string]interface{}{
		"facilityDetailId": data.ID,
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetFacilityDetail(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetFacilityDetail",
		attribute.String("operation", "get"),
		attribute.String("resource", "facility_detail"),
	)
	defer span.End()

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.FacilityDetail().GetFacilityDetail(ctx)
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
	)

	controllerutil.HandleStructuredLog(ctx, nil, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) UpdateFacilityDetail(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "UpdateFacilityDetail",
		attribute.String("operation", "update"),
		attribute.String("resource", "facility_detail"),
	)
	defer span.End()

	var params vo.UpdateFacilityDetailInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdateFacilityDetailInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.FacilityDetail().UpdateFacilityDetail(ctx, &params)
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
		attribute.String("facility_detail.id", data.ID),
	)

	logData := map[string]interface{}{
		"facilityDetailId": data.ID,
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteFacilityDetail(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "DeleteFacilityDetail",
		attribute.String("operation", "delete"),
		attribute.String("resource", "facility_detail"),
	)
	defer span.End()

	var params vo.DeleteFacilityDetailInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.DeleteFacilityDetailInput) error {
		return ctx.ShouldBindUri(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.FacilityDetail().DeleteFacilityDetail(ctx, &params)
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
		attribute.String("facility_detail.id", params.ID),
	)

	logData := map[string]interface{}{
		"facilityDetailId": params.ID,
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}
