package controllers

import (
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
	}

	codeStatus, err := services.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Register: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus)
		return
	}
	global.Logger.Info("Register: ", zap.String("info", params.VerifyKey))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CUserLogin) VerifyOTP(ctx *gin.Context) {
	var params vo.VerifyOTPInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
	}

	codeResult, data, err := services.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		global.Logger.Error("VerifyOTP: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult)
		return
	}
	global.Logger.Info("VerifyOTP: ", zap.String("info", params.VerifyKey))
	response.SuccessResponse(ctx, codeResult, data)
}

func (c *CUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params vo.UpdatePasswordRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
	}
	codeResult, err := services.UserLogin().UpdatePasswordRegister(ctx, &params)
	if err != nil {
		global.Logger.Error("UpdatePasswordRegister: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult)
		return
	}
	global.Logger.Info("UpdatePasswordRegister: ", zap.String("info", params.Token))
	response.SuccessResponse(ctx, codeResult, nil)

}

func (c *CUserLogin) Login(ctx *gin.Context) {
	var params vo.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid)
	}
	codeResult, data, err := services.UserLogin().Login(ctx, &params)
	if err != nil {
		global.Logger.Error("Login: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult)
		return
	}
	global.Logger.Info("Login: ", zap.String("info", data.Token))
	response.SuccessResponse(ctx, codeResult, data)
}
