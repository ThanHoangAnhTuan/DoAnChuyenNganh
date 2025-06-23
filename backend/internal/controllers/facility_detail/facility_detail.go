package facility_detail

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

func (c *Controller) CreateFacilityDetail(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateFacilityDetail validation not found\n")
		global.Logger.Error("CreateFacilityDetail validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CreateFacilityDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreateFacilityDetail binding error: %s\n", err.Error())
		global.Logger.Error("CreateFacilityDetail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("CreateFacilityDetail validation error: %s\n", validationErrors)
		global.Logger.Error("CreateFacilityDetail validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.FacilityDetail().CreateFacilityDetail(ctx, &params)
	if err != nil {
		fmt.Printf("CreateFacilityDetail error: %s\n", err.Error())
		global.Logger.Error("CreateFacilityDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CreateFacilityDetail success: %s\n", data)
	global.Logger.Info("CreateFacilityDetail success: ", zap.String("info", data.ID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetFacilityDetail(ctx *gin.Context) {
	codeResult, data, err := services.FacilityDetail().GetFacilityDetail(ctx)
	if err != nil {
		fmt.Printf("GetFacilityDetail error: %s\n", err.Error())
		global.Logger.Error("GetFacilityDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult, nil)
		return
	}

	fmt.Printf("GetFacilityDetail success\n")
	global.Logger.Info("GetFacilityDetail success")
	response.SuccessResponse(ctx, codeResult, data)
}

func (c *Controller) UpdateFacilityDetail(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("UpdateFacilityDetail validation not found\n")
		global.Logger.Error("UpdateFacilityDetail validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.UpdateFacilityDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("UpdateFacilityDetail binding error: %s\n", err.Error())
		global.Logger.Error("UpdateFacilityDetail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("UpdateFacilityDetail validation error: %s\n", validationErrors)
		global.Logger.Error("UpdateFacilityDetail validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.FacilityDetail().UpdateFacilityDetail(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateFacilityDetail error: %s\n", err.Error())
		global.Logger.Error("UpdateFacilityDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UpdateFacilityDetail success: %s\n", data.ID)
	global.Logger.Info("UpdateFacilityDetail success: ", zap.String("info", fmt.Sprintf("update facility detail success: %s", data.ID)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteFacilityDetail(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("DeleteFacilityDetail validation not found\n")
		global.Logger.Error("DeleteFacilityDetail validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.DeleteFacilityDetailInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("DeleteFacilityDetail binding error: %s\n", err.Error())
		global.Logger.Error("DeleteFacilityDetail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("DeleteFacilityDetail validation error: %s\n", validationErrors)
		global.Logger.Error("DeleteFacilityDetail validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.FacilityDetail().DeleteFacilityDetail(ctx, &params)
	if err != nil {
		fmt.Printf("DeleteFacilityDetail error: %s\n", err.Error())
		global.Logger.Error("DeleteFacilityDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("DeleteFacilityDetail success: %s\n", params.ID)
	global.Logger.Info("DeleteFacilityDetail success: ", zap.String("info", fmt.Sprintf("delete facility detail success: %s", params.ID)))
	response.SuccessResponse(ctx, codeStatus, nil)
}
