package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"go.uber.org/zap"
)

var Accommodation = new(CAccommodation)

type CAccommodation struct {
}

func (c *CAccommodation) CreateAccommodation(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateAccommodation validation not found\n")
		global.Logger.Error("CreateAccommodation validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CreateAccommodationInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreateAccommodation binding error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodation binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("CreateAccommodation validation error: %s\n", validationErrors)
		global.Logger.Error("CreateAccommodation validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Accommodation().CreateAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("CreateAccommodation error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CreateAccommodation success: manager: %s\naccommodation: %s\n", data.ManagerID, data.ID)
	global.Logger.Info("CreateAccommodation success: ",
		zap.String("info", fmt.Sprintf("manager:%s\naccommodation:%s", data.ManagerID, data.ID)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) GetAccommodations(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetAccommodations validation not found")
		global.Logger.Error("GetAccommodations validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.GetAccommodationsInput
	if err := ctx.ShouldBindQuery(&params); err != nil {
		fmt.Printf("GetAccommodations binding error: %s\n", err.Error())
		global.Logger.Error("GetAccommodations binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("GetAccommodations validation error: %s\n", validationErrors)
		global.Logger.Error("GetAccommodations validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, pagination, err := services.Accommodation().GetAccommodations(ctx, &params)
	if err != nil {
		fmt.Printf("GetAccommodations error: %s\n", err.Error())
		global.Logger.Error("GetAccommodations error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodations success: %s\n", "Get accommodations success")
	global.Logger.Info("GetAccommodations success: ", zap.String("info", "Get accommodations success"))
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}

func (c *CAccommodation) UpdateAccommodation(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("UpdateAccommodation validation not found\n")
		global.Logger.Error("UpdateAccommodation validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.UpdateAccommodationInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("UpdateAccommodation binding error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodation binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("UpdateAccommodation validation error: %s\n", validationErrors)
		global.Logger.Error("UpdateAccommodation validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Accommodation().UpdateAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateAccommodation error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UpdateAccommodation success: manager%s\naccommodation:%s\n", data.ManagerID, data.ID)
	global.Logger.Info("UpdateAccommodation success: ",
		zap.String("info", fmt.Sprintf("manager%s\naccommodation:%s", data.ManagerID, data.ID)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) DeleteAccommodation(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("DeleteAccommodation validation not found\n")
		global.Logger.Error("DeleteAccommodation validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}
	var params vo.DeleteAccommodationInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("DeleteAccommodation binding error")
		global.Logger.Error("DeleteAccommodation binding error")
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("DeleteAccommodation validation error: %s\n", validationErrors)
		global.Logger.Error("DeleteAccommodation validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.Accommodation().DeleteAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("DeleteAccommodation error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userId, _ := utils.GetUserIDFromGin(ctx)

	fmt.Printf("DeleteAccommodation success: userId:%s\naccommodationId:%s\n", userId, params.ID)
	global.Logger.Info("DeleteAccommodation success: ",
		zap.String("info", fmt.Sprintf("userId:%s\naccommodationId:%s", userId, params.ID)))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CAccommodation) GetAccommodation(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetAccommodationById validation not found\n")
		global.Logger.Error("GetAccommodationById validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}
	var params vo.GetAccommodationInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetAccommodationById binding error")
		global.Logger.Error("GetAccommodationById binding error")
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("GetAccommodationById validation error: %s\n", validationErrors)
		global.Logger.Error("GetAccommodationById validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Accommodation().GetAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("GetAccommodationById error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationById error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodationById success: %s\n", "Get accommodation by id success")
	global.Logger.Info("GetAccommodationById success: ", zap.String("info", "Get accommodation by id success"))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) GetAccommodationsByManager(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetAccommodationsByManager validation not found")
		global.Logger.Error("GetAccommodationsByManager validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.GetAccommodationsInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("GetAccommodationsByManager binding error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationsByManager binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("GetAccommodationsByManager validation error: %s\n", validationErrors)
		global.Logger.Error("GetAccommodationsByManager validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, pagination, err := services.Accommodation().GetAccommodationsByManager(ctx, &params)
	if err != nil {
		fmt.Printf("GetAccommodationsByManager error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationsByManager error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodationsByManager success: %s\n", "Get accommodations by manager success")
	global.Logger.Info("GetAccommodationsByManager success: ",
		zap.String("info", "Get accommodations by manager success"))
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)

}
