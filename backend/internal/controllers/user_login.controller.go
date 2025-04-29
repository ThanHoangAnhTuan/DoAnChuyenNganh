package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/services"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

var UserLogin = new(CUserLogin)

type CUserLogin struct {
}

func (c *CUserLogin) Register(ctx *gin.Context) {
	var params vo.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	codeStatus, err := services.UserLogin().Register(ctx, &params)
	if err != nil {
		fmt.Printf("Register error: %s\n", err.Error())
		global.Logger.Error("Register error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}

	fmt.Printf("Register success: %s\n", params.VerifyKey)
	global.Logger.Info("Register success: ", zap.String("info", params.VerifyKey))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CUserLogin) VerifyOTP(ctx *gin.Context) {
	var params vo.VerifyOTPInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	codeResult, data, err := services.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		fmt.Printf("VerifyOTP error: %s\n", err.Error())
		global.Logger.Error("VerifyOTP error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult)
		return
	}

	fmt.Printf("VerifyOTP success: %s\n", params.VerifyKey)
	global.Logger.Info("VerifyOTP success: ", zap.String("info", params.VerifyKey))
	response.SuccessResponse(ctx, codeResult, data)
}

func (c *CUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params vo.UpdatePasswordRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	codeResult, err := services.UserLogin().UpdatePasswordRegister(ctx, &params)
	if err != nil {
		fmt.Printf("UpdatePasswordRegister error: %s\n", err.Error())
		global.Logger.Error("UpdatePasswordRegister error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult)
		return
	}

	fmt.Printf("UpdatePasswordRegister success\n: %s", params.Token)
	global.Logger.Info("UpdatePasswordRegister success: ", zap.String("info", params.Token))
	response.SuccessResponse(ctx, codeResult, nil)

}

func (c *CUserLogin) Login(ctx *gin.Context) {
	var params vo.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
		return
	}

	codeResult, data, err := services.UserLogin().Login(ctx, &params)
	if err != nil {
		fmt.Printf("Login error\n: %s", err.Error())
		global.Logger.Error("Login error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult)
		return
	}
	fmt.Printf("Login success\n: %s", data.Token)
	global.Logger.Info("Login success: ", zap.String("info", data.Token))
	response.SuccessResponse(ctx, codeResult, data)
}
