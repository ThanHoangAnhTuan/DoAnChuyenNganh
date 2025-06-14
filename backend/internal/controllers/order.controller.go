package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/consts"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"go.uber.org/zap"
)

var Order = new(COrder)

type COrder struct {
}

func (c *COrder) CancelOrder(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CancelOrder validation not found \n")
		global.Logger.Error("CancelOrder validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CancelOrderInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CancelOrder binding error: %s\n", err.Error())
		global.Logger.Error("CancelOrder binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("CancelOrderInput validation error: %s\n", validationErrors)
		global.Logger.Error("CancelOrderInput validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.Order().CancelOrder(ctx, &params)
	if err != nil {
		fmt.Printf("CancelOrder error: %s\n", err.Error())
		global.Logger.Error("CancelOrder error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CancelOrder success: orderID: %s\n", params.OrderID)
	global.Logger.Info("CancelOrder success", zap.String("orderID", params.OrderID))

	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *COrder) CheckIn(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CheckIn validation not found \n")
		global.Logger.Error("CheckIn validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CheckInInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CheckIn binding error: %s\n", err.Error())
		global.Logger.Error("CheckIn binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("CheckInInput validation error: %s\n", validationErrors)
		global.Logger.Error("CheckInInput validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.Order().CheckIn(ctx, &params)
	if err != nil {
		fmt.Printf("CheckIn error: %s\n", err.Error())
		global.Logger.Error("CheckIn error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CheckIn success: orderID: %s\n", params.OrderID)
	global.Logger.Info("CheckIn success", zap.String("orderID", params.OrderID))

	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *COrder) CheckOut(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CheckOut validation not found \n")
		global.Logger.Error("CheckOut validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CheckOutInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CheckOut binding error: %s\n", err.Error())
		global.Logger.Error("CheckOut binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("CheckOutInput validation error: %s\n", validationErrors)
		global.Logger.Error("CheckOutInput validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.Order().CheckOut(ctx, &params)
	if err != nil {
		fmt.Printf("CheckOut error: %s\n", err.Error())
		global.Logger.Error("CheckOut error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CheckOut success: orderID: %s\n", params.OrderID)
	global.Logger.Info("CheckOut success", zap.String("orderID", params.OrderID))

	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *COrder) GetOrderInfoAfterPayment(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("GetOrderInfoAfterPayment validation not found \n")
		global.Logger.Error("GetOrderInfoAfterPayment validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.GetOrderInfoAfterPaymentInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("GetOrderInfoAfterPayment binding error: %s\n", err.Error())
		global.Logger.Error("GetOrderInfoAfterPayment binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("GetOrderInfoAfterPaymentInput validation error: %s\n", validationErrors)
		global.Logger.Error("GetOrderInfoAfterPaymentInput validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, data, err := services.Order().GetOrderInfoAfterPayment(ctx, &params)
	if err != nil {
		fmt.Printf("GetOrderInfoAfterPayment error: %s\n", err.Error())
		global.Logger.Error("GetOrderInfoAfterPayment error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetOrderInfoAfterPayment success: orderID: %s\n", data.OrderIDExternal)
	global.Logger.Info("GetOrderInfoAfterPayment success",
		zap.String("orderID", data.OrderIDExternal))

	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *COrder) GetOrdersByManager(ctx *gin.Context) {
	codeStatus, data, err := services.Order().GetOrdersByManager(ctx)
	if err != nil {
		fmt.Printf("GetOrdersByManager error: %s\n", err.Error())
		global.Logger.Error("GetOrdersByManager error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	// TODO: get managerID
	managerID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("GetOrdersByManager cannot found managerID")
		global.Logger.Error("GetOrdersByManager cannot found managerID")
		managerID = consts.UNKNOWN
	}

	fmt.Printf("GetOrdersByManager success: %s\n", managerID)
	global.Logger.Info("GetOrdersByManager success", zap.String("info", managerID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *COrder) GetOrdersByUser(ctx *gin.Context) {
	codeStatus, data, err := services.Order().GetOrdersByUser(ctx)
	if err != nil {
		fmt.Printf("GetOrdersByUser error: %s\n", err.Error())
		global.Logger.Error("GetOrdersByUser error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	// TODO: get managerID
	managerID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		fmt.Printf("GetOrdersByUser cannot found managerID")
		global.Logger.Error("GetOrdersByUser cannot found managerID")
		managerID = consts.UNKNOWN
	}

	fmt.Printf("GetOrdersByUser success: %s\n", managerID)
	global.Logger.Info("GetOrdersByUser success", zap.String("info", managerID))
	response.SuccessResponse(ctx, codeStatus, data)
}
