package review

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"go.uber.org/zap"
)

func (c *Controller) CreateReview(ctx *gin.Context) {
	var params vo.CreateReviewInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateReviewInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, data, err := services.Review().CreateReview(ctx, &params)
	if err != nil {
		fmt.Printf("CreateReview error: %s\n", err.Error())
		global.Logger.Error("CreateReview error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CreateReview success: %s\n", data.ID)
	global.Logger.Info("CreateReview success", zap.String("info", data.ID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetReviews(ctx *gin.Context) {
	var params vo.GetReviewsInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetReviewsInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, data, pagination, err := services.Review().GetReviews(ctx, &params)
	if err != nil {
		fmt.Printf("GetReviews error: %s\n", err.Error())
		global.Logger.Error("GetReviews error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetReviews success: %s\n", params.AccommodationID)
	global.Logger.Info("GetReviews success: ", zap.String("info", params.AccommodationID))
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}
