package user

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

type CUserLogin struct {
}

func (c *CUserLogin) Register(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("User register validation not found\n")
		global.Logger.Error("User register validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("User register binding error: %s\n", err.Error())
		global.Logger.Error("User register binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("User register validation error: %s\n", validationErrors)
		global.Logger.Error("User register validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.UserLogin().Register(ctx, &params)
	if err != nil {
		fmt.Printf("User register error: %s\n", err.Error())
		global.Logger.Error("User register error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("User register success: %s\n", params.VerifyKey)
	global.Logger.Info("User register success: ", zap.String("info", params.VerifyKey))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CUserLogin) VerifyOTP(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("User verifyOTP validation not found\n")
		global.Logger.Error("User verifyOTP validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.VerifyOTPInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("User verifyOTP binding error: %s\n", err.Error())
		global.Logger.Error("User verifyOTP binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("User verifyOTP validation error: %s\n", err.Error())
		global.Logger.Error("User verifyOTP validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		fmt.Printf("User verifyOTP error: %s\n", err.Error())
		global.Logger.Error("User verifyOTP error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("User verifyOTP success: %s\n", params.VerifyKey)
	global.Logger.Info("User verifyOTP success: ", zap.String("info", params.VerifyKey))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("User updatePasswordRegister validation not found\n")
		global.Logger.Error("User updatePasswordRegister validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.UpdatePasswordRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("User updatePasswordRegister binding error: %s\n", err.Error())
		global.Logger.Error("User updatePasswordRegister binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("User updatePasswordRegister validation error: %s\n", err.Error())
		global.Logger.Error("User updatePasswordRegister validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, err := services.UserLogin().UpdatePasswordRegister(ctx, &params)
	if err != nil {
		fmt.Printf("User updatePasswordRegister error: %s\n", err.Error())
		global.Logger.Error("User updatePasswordRegister error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("User updatePasswordRegister success: %s\n", params.Token)
	global.Logger.Info("User updatePasswordRegister success: ", zap.String("info", params.Token))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CUserLogin) Login(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("User login validation not found\n")
		global.Logger.Error("User login validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("User login binding eror: %s\n", err.Error())
		global.Logger.Error("User login binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("User login validation eror: %s\n", err.Error())
		global.Logger.Error("User login validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	codeStatus, data, err := services.UserLogin().Login(ctx, &params)
	if err != nil {
		fmt.Printf("User login error: %s\n", err.Error())
		global.Logger.Error("User login error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}
	fmt.Printf("User login success: %s\n", data.Token)
	global.Logger.Info("User login success: ", zap.String("info", data.Token))
	response.SuccessResponse(ctx, codeStatus, data)
}
