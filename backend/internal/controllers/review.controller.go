package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"go.uber.org/zap"
)

var Review = new(CReview)

type CReview struct {
}

func (c *CReview) CreateReview(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateReview validation not found\n")
		global.Logger.Error("CreateReview validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CreateReviewInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreateReview binding error: %s\n", err.Error())
		global.Logger.Error("CreateReview binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("CreateReview validation error: %s\n", validationErrors)
		global.Logger.Error("CreateReview validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
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

func (c *CReview) GetReviews(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetReviews validation not found\n")
		global.Logger.Error("GetReviews validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.GetReviewsInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("GetReviews binding error: %s\n", err.Error())
		global.Logger.Error("GetReviews binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("GetReviews validation error: %s\n", validationErrors)
		global.Logger.Error("GetReviews validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
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
