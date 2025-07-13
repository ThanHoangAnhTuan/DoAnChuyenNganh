package payment

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

func (c *Controller) CreatePaymentURL(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "CreatePaymentURL",
		attribute.String("operation", "create"),
		attribute.String("resource", "payment"),
	)
	defer span.End()

	var params vo.CreatePaymentURLInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreatePaymentURLInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Payment().CreatePaymentURL(ctx, &params)
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

func (c *Controller) VNPayReturn(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "VNPayReturn",
		attribute.String("operation", "return"),
		attribute.String("resource", "payment"),
	)
	defer span.End()

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, err := services.Payment().VNPayReturn(ctx)
	duration := time.Since(start)

	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
		controllerutil.HandleStructuredLog(ctx, nil, "error", codeStatus, duration, err)
		return
	}

	span.SetAttributes(
		attribute.Int("status_code", codeStatus),
		attribute.Int64("duration_ms", duration.Milliseconds()),
	)

	controllerutil.HandleStructuredLog(ctx, nil, "success", codeStatus, duration, nil)
}

func (c *Controller) VNPayIPN(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "VNPayIPN",
		attribute.String("operation", "ipn"),
		attribute.String("resource", "payment"),
	)
	defer span.End()

	ctx.Request = ctx.Request.WithContext(spanCtx)

	services.Payment().VNPayIPN(ctx)
	duration := time.Since(start)

	span.SetAttributes(
		attribute.Int64("duration_ms", duration.Milliseconds()),
	)

	controllerutil.HandleStructuredLog(ctx, nil, "success", 200, duration, nil)
}

func (c *Controller) PostQueryDR(ctx *gin.Context) {

}

func (c *Controller) PostRefund(ctx *gin.Context) {

}
