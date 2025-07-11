package manager

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

func (c *Controller) Register(ctx *gin.Context) {
	var params vo.ManagerRegisterInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.ManagerRegisterInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
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

func (c *Controller) Login(ctx *gin.Context) {
	var params vo.ManagerLoginInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.ManagerLoginInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
		return
	}

	codeStatus, data, err := services.ManagerLogin().Login(ctx, &params)
	if err != nil {
		fmt.Printf("Manager login error: %s\n", err.Error())
		global.Logger.Error("Manager login error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("Manager login success: %s\n", data.Token)
	global.Logger.Info("Manager login success: ", zap.String("info", data.Token))
	response.SuccessResponse(ctx, codeStatus, data)
}
