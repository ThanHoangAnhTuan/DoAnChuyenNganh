package admin

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

type CAdminLogin struct{}

func (c *CAdminLogin) Register(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Admin register validation not found\n")
		global.Logger.Error("Admin register validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.AdminRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("Admin register binding error: %s\n", err.Error())
		global.Logger.Error("Admin register binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("Admin register validation error: %s\n", validationErrors)
		global.Logger.Error("Admin register validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeStatus, err := services.AdminLogin().Register(ctx, &params)
	if err != nil {
		fmt.Printf("Admin register error: %s\n", err.Error())
		global.Logger.Error("Admin register error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("Admin register success: %s\n", params.UserAccount)
	global.Logger.Info("Admin register success: ", zap.String("info", params.UserAccount))
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CAdminLogin) Login(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("Admin login validation not found\n")
		global.Logger.Error("Admin login validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.AdminLoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		fmt.Printf("Admin login binding error: %s\n", err.Error())
		global.Logger.Error("Admin login binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err)
		fmt.Printf("Admin login validation error: %s\n", validationErrors)
		global.Logger.Error("Admin login validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
		return
	}

	codeResult, data, err := services.AdminLogin().Login(ctx, &params)
	if err != nil {
		fmt.Printf("Admin login error: %s\n", err.Error())
		global.Logger.Error("Admin login error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeResult, nil)
		return
	}

	fmt.Printf("Admin login success: %s\n", data.Token)
	global.Logger.Info("Admin login success: ", zap.String("info", data.Token))
	response.SuccessResponse(ctx, codeResult, data)
}
