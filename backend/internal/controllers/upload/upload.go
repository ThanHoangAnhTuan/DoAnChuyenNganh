package upload

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

func (c *Controller) UploadImages(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "UploadImages",
		attribute.String("operation", "upload"),
		attribute.String("resource", "upload"),
	)
	defer span.End()

	var params vo.UploadImages

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UploadImages) error {
		return ctx.ShouldBindQuery(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Upload().UploadImages(ctx, &params)
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
		attribute.String("upload.id", params.ID),
		attribute.Bool("is_detail", params.IsDetail),
	)

	logData := map[string]interface{}{
		"uploadId": params.ID,
		"isDetail": params.IsDetail,
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetImages(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetImages",
		attribute.String("operation", "get"),
		attribute.String("resource", "upload"),
	)
	defer span.End()

	var params vo.GetImagesInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetImagesInput) error {
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

	codeStatus, data, err := services.Upload().GetImages(ctx, &params)
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
		attribute.String("image.id", params.ID),
		attribute.Bool("is_detail", params.IsDetail),
	)

	logData := map[string]interface{}{
		"imageId":  params.ID,
		"isDetail": params.IsDetail,
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}
