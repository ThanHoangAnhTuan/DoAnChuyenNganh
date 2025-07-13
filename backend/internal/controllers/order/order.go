package order

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

func (c *Controller) CancelOrder(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "CancelOrder",
		attribute.String("operation", "cancel"),
		attribute.String("resource", "order"),
	)
	defer span.End()

	var params vo.CancelOrderInput
	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CancelOrderInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.Order().CancelOrder(ctx, &params)
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
		attribute.String("order.id", params.OrderID),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) CheckIn(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "CheckIn",
		attribute.String("operation", "check_in"),
		attribute.String("resource", "order"),
	)
	defer span.End()

	var params vo.CheckInInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CheckInInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.Order().CheckIn(ctx, &params)
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
		attribute.String("order.id", params.OrderID),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) CheckOut(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "CheckOut",
		attribute.String("operation", "check_out"),
		attribute.String("resource", "order"),
	)
	defer span.End()

	var params vo.CheckOutInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CheckOutInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.Order().CheckOut(ctx, &params)
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
		attribute.String("order.id", params.OrderID),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) GetOrderInfoAfterPayment(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetOrderInfoAfterPayment",
		attribute.String("operation", "get_info"),
		attribute.String("resource", "order"),
	)
	defer span.End()

	var params vo.GetOrderInfoAfterPaymentInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetOrderInfoAfterPaymentInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Order().GetOrderInfoAfterPayment(ctx, &params)
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
		attribute.String("order.id", params.OrderIDExternal),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetOrdersByManager(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetOrdersByManager",
		attribute.String("operation", "list"),
		attribute.String("resource", "order"),
	)
	defer span.End()

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Order().GetOrdersByManager(ctx)
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
		attribute.Int("orders.count", len(data)),
	)

	controllerutil.HandleStructuredLog(ctx, nil, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetOrdersByUser(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetOrdersByUser",
		attribute.String("operation", "list"),
		attribute.String("resource", "order"),
	)
	defer span.End()

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Order().GetOrdersByUser(ctx)
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
		attribute.Int("orders.count", len(data)),
	)

	controllerutil.HandleStructuredLog(ctx, nil, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}
