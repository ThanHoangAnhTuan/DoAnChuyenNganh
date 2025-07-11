package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"go.uber.org/zap"
)

type CUserLogin struct {
}

func (c *CUserLogin) Register(ctx *gin.Context) {
	var params vo.RegisterInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.RegisterInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
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
	var params vo.VerifyOTPInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.VerifyOTPInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
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
	var params vo.UpdatePasswordRegisterInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdatePasswordRegisterInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
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
	var params vo.LoginInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.LoginInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
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
