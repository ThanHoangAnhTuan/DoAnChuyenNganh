package review

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

func (c *Controller) CreateReview(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "CreateReview",
		attribute.String("operation", "create"),
		attribute.String("resource", "review"),
	)
	defer span.End()

	var params vo.CreateReviewInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateReviewInput) error {
		return ctx.ShouldBindJSON(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, err := services.Review().CreateReview(ctx, &params)
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
		attribute.String("review.id", data.ID),
		attribute.String("accommodation.id", params.AccommodationID),
		attribute.Int("rating", int(params.Rating)),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetReviews(ctx *gin.Context) {
	start := time.Now()

	spanCtx, span := middlewares.StartChildSpan(ctx.Request.Context(), "GetReviews",
		attribute.String("operation", "list"),
		attribute.String("resource", "review"),
	)
	defer span.End()

	var params vo.GetReviewsInput

	if validationErr := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetReviewsInput) error {
		return ctx.ShouldBindQuery(p)
	}); validationErr != nil {
		duration := time.Since(start)
		span.SetAttributes(attribute.String("error", validationErr.Message))
		controllerutil.HandleValidationError(ctx, params, validationErr, duration)
		return
	}

	ctx.Request = ctx.Request.WithContext(spanCtx)

	codeStatus, data, pagination, err := services.Review().GetReviews(ctx, &params)
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
		attribute.String("accommodation.id", params.AccommodationID),
		attribute.Int("reviews.count", len(data)),
	)

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}
