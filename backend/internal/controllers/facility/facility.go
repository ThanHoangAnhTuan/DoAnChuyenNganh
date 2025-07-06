package facility

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

func (c *Controller) CreateFacility(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateFacility validation not found\n")
		global.Logger.Error("CreateFacility validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.CreateFacilityInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreateFacility binding error: %s\n", err.Error())
		global.Logger.Error("CreateFacility binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("CreateFacility validation error: %s\n", validationErrors)
		global.Logger.Error("CreateFacility validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Facility().CreateFacility(ctx, &params)
	if err != nil {
		fmt.Printf("CreateFacility error: %s\n", err.Error())
		global.Logger.Error("CreateFacility error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CreateFacility success: %s\n", data.ID)
	global.Logger.Info("CreateFacility success: ", zap.String("info", fmt.Sprintf("create facility success: %s", data.ID)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) UpdateFacility(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("UpdateFacility validation not found\n")
		global.Logger.Error("UpdateFacility validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.UpdateFacilityInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("UpdateFacility binding error: %s\n", err.Error())
		global.Logger.Error("UpdateFacility binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("UpdateFacility validation error: %s\n", validationErrors)
		global.Logger.Error("UpdateFacility validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Facility().UpdateFacility(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateFacility error: %s\n", err.Error())
		global.Logger.Error("UpdateFacility error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UpdateFacility success: %s\n", data.ID)
	global.Logger.Info("UpdateFacility success: ", zap.String("info", fmt.Sprintf("update facility success: %s", data.ID)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteFacility(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("DeleteFacility validation not found\n")
		global.Logger.Error("DeleteFacility validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.DeleteFacilityInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("DeleteFacility binding error: %s\n", err.Error())
		global.Logger.Error("DeleteFacility binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("DeleteFacility validation error: %s\n", validationErrors)
		global.Logger.Error("DeleteFacility validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.Facility().DeleteFacility(ctx, &params)
	if err != nil {
		fmt.Printf("DeleteFacility error: %s\n", err.Error())
		global.Logger.Error("DeleteFacility error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("DeleteFacility success: %s\n", params.ID)
	global.Logger.Info("DeleteFacility success: ", zap.String("info", fmt.Sprintf("delete facility success: %s", params.ID)))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) GetFacilities(ctx *gin.Context) {
	codeStatus, data, err := services.Facility().GetFacilities(ctx)
	if err != nil {
		fmt.Printf("GetFacilities error: %s\n", err.Error())
		global.Logger.Error("GetFacilities error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetFacilities success\n")
	global.Logger.Info("GetFacilities success")
	response.SuccessResponse(ctx, codeStatus, data)
}
