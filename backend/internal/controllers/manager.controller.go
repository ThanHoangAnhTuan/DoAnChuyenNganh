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

var ManagerLogin = new(CManagerLogin)

type CManagerLogin struct {
}

func (c *CManagerLogin) Register(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Manager register validation not found\n")
		global.Logger.Error("Manager register validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.ManagerRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("Manager register binding error: %s\n", err.Error())
		global.Logger.Error("Manager register binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("Manager register validation error: %s\n", err.Error())
		global.Logger.Error("Manager register validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, err := services.ManagerLogin().Register(ctx, &params)
	if err != nil {
		fmt.Printf("Manager register error: %s\n", err.Error())
		global.Logger.Error("Manager register error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("Manager register success: %s\n", params.UserAccount)
	global.Logger.Info("Manager register success: ", zap.String("info", params.UserAccount))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CManagerLogin) Login(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Manager login validation not found\n")
		global.Logger.Error("Manager login validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.ManagerLoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("Manager login binding error\n: %s", err.Error())
		global.Logger.Error("Manager login binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("Manager login validation error\n: %s", err.Error())
		global.Logger.Error("Manager login validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeResult, data, err := services.ManagerLogin().Login(ctx, &params)
	if err != nil {
		fmt.Printf("Manager login error\n: %s", err.Error())
		global.Logger.Error("Manager login error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult, nil)
		return
	}

	fmt.Printf("Manager login success\n: %s", data.Token)
	global.Logger.Info("Manager login success: ", zap.String("info", data.Token))
	response.SuccessResponse(ctx, codeResult, data)
}
