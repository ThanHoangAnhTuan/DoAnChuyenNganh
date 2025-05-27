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

var Order = new(COrder)

type COrder struct {
}

func (c *COrder) CreateOrder(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreateOrder validation not found\n")
		global.Logger.Error("CreateOrder validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CreateOrderInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("CreateOrder binding error: %s\n", err.Error())
		global.Logger.Error("CreateOrder binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("CreateOrder validation error: %s\n", err.Error())
		global.Logger.Error("CreateOrder validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.Order().CreateOrder(ctx, &params)
	if err != nil {
		fmt.Printf("CreateOrder error: %s\n", err.Error())
		global.Logger.Error("CreateOrder error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CreateOrder success: %s\n", data.OrderID)
	global.Logger.Info("CreateOrder success: ", zap.String("info", data.OrderID))
	response.SuccessResponse(ctx, codeStatus, data)
}
