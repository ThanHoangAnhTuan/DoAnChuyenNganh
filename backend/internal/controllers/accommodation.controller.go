package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/services"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils"
	"go.uber.org/zap"
)

var Accommodation = new(CAccommodation)

type CAccommodation struct {
}

func (c *CAccommodation) CreateAccommodation(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateAccommodation validation not found")
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
		fmt.Printf("CreateAccommodation validation error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodation validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.Accommodation().CreateAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("CreateAccommodation error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CreateAccommodation success: manager: %s\taccommodation: %s\n", data.ManagerId, data.Id)
	global.Logger.Info("CreateAccommodation success: ",
		zap.String("info", fmt.Sprintf("manager:%s\taccommodation:%s", data.ManagerId, data.Id)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) GetAccommodations(ctx *gin.Context) {
	codeStatus, data, err := services.Accommodation().GetAccommodations(ctx)
	if err != nil {
		fmt.Printf("GetAccommodations error: %s\n", err.Error())
		global.Logger.Error("GetAccommodations error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodations success: %s\n", "Get accommodations success")
	global.Logger.Info("GetAccommodations success: ", zap.String("info", "Get accommodations success"))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) UpdateAccommodation(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("UpdateAccommodation validation not found")
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
		fmt.Printf("UpdateAccommodation validation error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodation validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.Accommodation().UpdateAccommodation(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateAccommodation error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UpdateAccommodation success: manager%s\taccommodation:%s\n", data.ManagerId, data.Id)
	global.Logger.Info("UpdateAccommodation success: ",
		zap.String("info", fmt.Sprintf("manager%s\taccommodation:%s", data.ManagerId, data.Id)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) DeleteAccommodation(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("DeleteAccommodation validation not found")
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
		fmt.Printf("DeleteAccommodation validation error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodation validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
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

	fmt.Printf("DeleteAccommodation success: userId:%s\taccommodationId:%s\n", userId, params.Id)
	global.Logger.Info("DeleteAccommodation success: ",
		zap.String("info", fmt.Sprintf("userId:%s\taccommodationId:%s", userId, params.Id)))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CAccommodation) GetAccommodationByCity(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}
	var params vo.GetAccommodationByCityInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetAccommodationByCity binding error")
		global.Logger.Error("GetAccommodationByCity binding error")
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("GetAccommodationByCity validation error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationByCity validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.Accommodation().GetAccommodationByCity(ctx, &params)
	if err != nil {
		fmt.Printf("GetAccommodationByCity error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationByCity error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodationByCity success: %s\n", "Get accommodation by city success")
	global.Logger.Info("GetAccommodationByCity success: ", zap.String("info", "Get accommodation by city success"))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodation) GetAccommodationById(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}
	var params vo.GetAccommodationByIdInput
	if err := ctx.ShouldBindUri(&params); err != nil {
		fmt.Printf("GetAccommodationById binding error")
		global.Logger.Error("GetAccommodationById binding error")
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("GetAccommodationById validation error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationById validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.Accommodation().GetAccommodationById(ctx, &params)
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
	codeStatus, data, err := services.Accommodation().GetAccommodationsByManager(ctx)
	if err != nil {
		fmt.Printf("GetAccommodationsByManager error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationsByManager error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodationsByManager success: %s\n", "Get accommodations by manager success")
	global.Logger.Info("GetAccommodationsByManager success: ",
		zap.String("info", "Get accommodations by manager success"))
	response.SuccessResponse(ctx, codeStatus, data)
}
