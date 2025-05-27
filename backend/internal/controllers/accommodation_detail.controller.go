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

var AccommodationDetail = new(CAccommodationDetail)

type CAccommodationDetail struct {
}

func (c *CAccommodationDetail) CreateAccommodationDetail(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateAccommodationDetail Validation not found\n")
		global.Logger.Error("CreateAccommodationDetail Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CreateAccommodationDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreateAccommodationDetail binding error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodationDetail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("CreateAccommodationDetail validation error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodationDetail validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.AccommodationDetail().CreateAccommodationDetail(ctx, &params)
	if err != nil {
		fmt.Printf("CreateAccommodationDetail error: %s\n", err.Error())
		global.Logger.Error("CreateAccommodationDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userId, _ := utils.GetUserIDFromGin(ctx)

	fmt.Printf("CreateAccommodationDetail success: manager: %s\taccommodation detail: %s\n", userId, data.ID)
	global.Logger.Info("CreateAccommodationDetail success: ",
		zap.String("info", fmt.Sprintf("manager:%s\taccommodation detail:%s", userId, data.ID)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodationDetail) GetAccommodationDetails(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetAccommodationDetails Validation not found\n")
		global.Logger.Error("GetAccommodationDetails Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.GetAccommodationDetailsInput
	id := ctx.Param("id")
	if id == "" {
		fmt.Printf("GetAccommodationDetails binding error\n")
		global.Logger.Error("GetAccommodationDetails binding error")
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("GetAccommodationDetails validation error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetails validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}
	params.AccommodationID = id

	codeStatus, data, err := services.AccommodationDetail().GetAccommodationDetails(ctx, &params)
	if err != nil {
		fmt.Printf("GetAccommodationDetails error: %s\n", err.Error())
		global.Logger.Error("GetAccommodationDetails error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetAccommodationDetails success: %s\n", "Get accommodations details success")
	global.Logger.Info("GetAccommodationDetails success: ", zap.String("info", "Get accommodations details success"))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodationDetail) UpdateAccommodationDetail(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetAccommodationDetails Validation not found \n")
		global.Logger.Error("GetAccommodationDetails Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.UpdateAccommodationDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("UpdateAccommodationDetail binding error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodationDetail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("UpdateAccommodationDetail validation error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodationDetail validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.AccommodationDetail().UpdateAccommodationDetail(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateAccommodationDetail error: %s\n", err.Error())
		global.Logger.Error("UpdateAccommodationDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userId, _ := utils.GetUserIDFromGin(ctx)

	fmt.Printf("UpdateAccommodationDetail success: manager%s\taccommodation detail:%s\n", userId, data.ID)
	global.Logger.Info("UpdateAccommodationDetail success: ",
		zap.String("info", fmt.Sprintf("manager%s\taccommodation detail:%s", userId, data.ID)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CAccommodationDetail) DeleteAccommodationDetail(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("DeleteAccommodationDetail Validation not found\n")
		global.Logger.Error("DeleteAccommodationDetail Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.DeleteAccommodationDetailInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("DeleteAccommodationDetail binding error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodationDetail binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("DeleteAccommodationDetail validation error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodationDetail validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, err := services.AccommodationDetail().DeleteAccommodationDetail(ctx, &params)
	if err != nil {
		fmt.Printf("DeleteAccommodationDetail error: %s\n", err.Error())
		global.Logger.Error("DeleteAccommodationDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	userId, _ := utils.GetUserIDFromGin(ctx)

	fmt.Printf("DeleteAccommodationDetail success: userId:%s\taccommodation detail id:%s\n", userId, params.ID)
	global.Logger.Info("DeleteAccommodationDetail success: ",
		zap.String("info", fmt.Sprintf("userId:%s\taccommodation detail id:%s", userId, params.ID)))
	response.SuccessResponse(ctx, codeStatus, nil)
}
