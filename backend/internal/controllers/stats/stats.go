package stats

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"go.opentelemetry.io/otel/attribute"
)

func (c *Controller) GetMonthlyEarnings(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetMonthlyEarnings",
		attribute.String("operation", "get"),
		attribute.String("resource", "stats"),
	)
	defer span.End()

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Stats().GetMonthlyEarnings(ctx)
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

func (c *Controller) GetDailyEarnings(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetDailyEarnings",
		attribute.String("operation", "get"),
		attribute.String("resource", "stats"),
	)
	defer span.End()

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Stats().GetDailyEarnings(ctx)
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

func (c *Controller) GetDailyEarningsByMonth(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetDailyEarningsByMonth",
		attribute.String("operation", "get"),
		attribute.String("resource", "stats"),
	)
	defer span.End()

	var params vo.GetDailyEarningsByMonthInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetDailyEarningsByMonthInput) error {
		return ctx.ShouldBindUri(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Stats().GetDailyEarningsByMonth(ctx, &params)
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
		attribute.Int("year", params.Year),
		attribute.Int("month", params.Month),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetMonthlyEarningsByYear(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetMonthlyEarningsByYear",
		attribute.String("operation", "get"),
		attribute.String("resource", "stats"),
	)
	defer span.End()

	var params vo.GetMonthlyEarningsByYearInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetMonthlyEarningsByYearInput) error {
		return ctx.ShouldBindUri(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Stats().GetMonthlyEarningsByYear(ctx, &params)
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
		attribute.Int("year", params.Year),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) ExportDailyEarningsCSV(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "ExportDailyEarningsCSV",
		attribute.String("operation", "export"),
		attribute.String("resource", "stats"),
	)
	defer span.End()

	var params vo.GetDailyEarningsByMonthInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetDailyEarningsByMonthInput) error {
		return ctx.ShouldBindUri(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Stats().ExportDailyEarningsCSV(ctx, &params)
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
		attribute.Int("year", params.Year),
		attribute.Int("month", params.Month),
		attribute.Int("csv.size_bytes", len(data)),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)

	filename := fmt.Sprintf("daily_earnings_%d_%02d.csv", params.Year, params.Month)
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Header("Content-Length", strconv.Itoa(len(data)))

	ctx.Data(http.StatusOK, "text/csv", data)
}

func (c *Controller) ExportMonthlyEarningsCSV(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "ExportMonthlyEarningsCSV",
		attribute.String("operation", "export"),
		attribute.String("resource", "stats"),
	)
	defer span.End()

	var params vo.GetMonthlyEarningsByYearInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetMonthlyEarningsByYearInput) error {
		return ctx.ShouldBindUri(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Stats().ExportMonthlyEarningsCSV(ctx, &params)
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
		attribute.Int("year", params.Year),
		attribute.Int("csv.size_bytes", len(data)),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)

	filename := fmt.Sprintf("monthly_earnings_%d.csv", params.Year)
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Header("Content-Length", strconv.Itoa(len(data)))

	ctx.Data(http.StatusOK, "text/csv", data)
}
