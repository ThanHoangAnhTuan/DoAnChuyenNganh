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

var UserInfo = new(CUserInfo)

type CUserInfo struct {
}

func (c *CUserInfo) GetUserInfo(ctx *gin.Context) {
	codeStatus, data, err := services.UserInfo().GetUserInfo(ctx)
	if err != nil {
		fmt.Printf("GetUserInfo error: %s\n", err.Error())
		global.Logger.Error("GetUserInfo error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetUserInfo success: %s\n", data.ID)
	global.Logger.Info("GetUserInfo success: ", zap.String("info", data.ID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CUserInfo) UpdateUserInfo(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("UpdateUserInfo validation not found")
		global.Logger.Error("UpdateUserInfo validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.UpdateUserInfoInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("UpdateUserInfo binding error: %s\n", err.Error())
		global.Logger.Error("UpdateUserInfo binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("UpdateUserInfo validation error: %s\n", err.Error())
		global.Logger.Error("UpdateUserInfo validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeResult, data, err := services.UserInfo().UpdateUserInfo(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateUserInfo error: %s\n", err.Error())
		global.Logger.Error("UpdateUserInfo error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult, nil)
		return
	}

	fmt.Printf("UpdateUserInfo success: %s\n", params.Account)
	global.Logger.Info("UpdateUserInfo success: ", zap.String("info", params.Account))
	response.SuccessResponse(ctx, codeResult, data)
}
