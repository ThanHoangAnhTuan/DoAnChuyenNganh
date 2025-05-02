package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("Register binding error: %s\n", err.Error())
		global.Logger.Error("Register binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("Register validation error: %s\n", err.Error())
		global.Logger.Error("Register validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, err := services.UserLogin().Register(ctx, &params)
	if err != nil {
		fmt.Printf("Register error: %s\n", err.Error())
		global.Logger.Error("Register error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("Register success: %s\n", params.VerifyKey)
	global.Logger.Info("Register success: ", zap.String("info", params.VerifyKey))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CUserLogin) VerifyOTP(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.VerifyOTPInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("VerifyOTP binding error: %s\n", err.Error())
		global.Logger.Error("VerifyOTP binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("VerifyOTP validation error: %s\n", err.Error())
		global.Logger.Error("VerifyOTP validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeResult, data, err := services.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		fmt.Printf("VerifyOTP error: %s\n", err.Error())
		global.Logger.Error("VerifyOTP error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult, nil)
		return
	}

	fmt.Printf("VerifyOTP success: %s\n", params.VerifyKey)
	global.Logger.Info("VerifyOTP success: ", zap.String("info", params.VerifyKey))
	response.SuccessResponse(ctx, codeResult, data)
}

func (c *CUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.UpdatePasswordRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("UpdatePasswordRegister binding error: %s\n", err.Error())
		global.Logger.Error("UpdatePasswordRegister binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("UpdatePasswordRegister validation error: %s\n", err.Error())
		global.Logger.Error("UpdatePasswordRegister validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeResult, err := services.UserLogin().UpdatePasswordRegister(ctx, &params)
	if err != nil {
		fmt.Printf("UpdatePasswordRegister error: %s\n", err.Error())
		global.Logger.Error("UpdatePasswordRegister error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult, nil)
		return
	}

	fmt.Printf("UpdatePasswordRegister success\n: %s", params.Token)
	global.Logger.Info("UpdatePasswordRegister success: ", zap.String("info", params.Token))
	response.SuccessResponse(ctx, codeResult, nil)
}

func (c *CUserLogin) Login(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("Login binding eror: %s\n", err.Error())
		global.Logger.Error("Login binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("Login validation eror: %s\n", err.Error())
		global.Logger.Error("Login validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeResult, data, err := services.UserLogin().Login(ctx, &params)
	if err != nil {
		fmt.Printf("Login error\n: %s", err.Error())
		global.Logger.Error("Login error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult, nil)
		return
	}
	fmt.Printf("Login success\n: %s", data.Token)
	global.Logger.Info("Login success: ", zap.String("info", data.Token))
	response.SuccessResponse(ctx, codeResult, data)
}
