package upload

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

func (c *Controller) UploadImages(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.UploadImages
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("UploadImages binding error: %s\n", err.Error())
		global.Logger.Error("UploadImages binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("UploadImages validation error: %s\n", validationErrors)
		global.Logger.Error("UploadImages validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Upload().UploadImages(ctx, &params)
	if err != nil {
		fmt.Printf("UploadImages error: %s\n", err.Error())
		global.Logger.Error("UploadImages error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UploadImages success: %s", params.ID)
	global.Logger.Info("UploadImages success", zap.String("info", params.ID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetImages(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.GetImagesInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetImages id binding error: %s\n", err.Error())
		global.Logger.Error("GetImages id binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	if err := ctx.ShouldBindQuery(&params); err != nil {
		fmt.Printf("GetImages is detail binding error: %s\n", err.Error())
		global.Logger.Error("GetImages is detail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("GetImages validation error: %s\n", validationErrors)
		global.Logger.Error("GetImages validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Upload().GetImages(ctx, &params)
	if err != nil {
		fmt.Printf("GetImages error: %s\n", err.Error())
		global.Logger.Error("GetImages error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetImages success")
	global.Logger.Info("GetImages success")
	response.SuccessResponse(ctx, codeStatus, data)
}
