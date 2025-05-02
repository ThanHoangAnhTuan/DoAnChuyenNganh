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

var ManagerLogin = new(CManagerLogin)

type CManagerLogin struct {
}

func (c *CManagerLogin) Register(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.ManagerRegisterInput
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

	codeStatus, err := services.ManagerLogin().Register(ctx, &params)
	if err != nil {
		fmt.Printf("Register error: %s\n", err.Error())
		global.Logger.Error("Register error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("Register success: %s\n", params.UserAccount)
	global.Logger.Info("Register success: ", zap.String("info", params.UserAccount))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CManagerLogin) Login(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Validation not found")
		global.Logger.Error("Validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.ManagerLoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("Login binding error\n: %s", err.Error())
		global.Logger.Error("Login binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("Login validation error\n: %s", err.Error())
		global.Logger.Error("Login validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeResult, data, err := services.ManagerLogin().Login(ctx, &params)
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
