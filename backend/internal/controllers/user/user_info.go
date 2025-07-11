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
	var params vo.UpdateUserInfoInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdateUserInfoInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
		return
	}

	codeStatus, data, err := services.UserInfo().UpdateUserInfo(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateUserInfo error: %s\n", err.Error())
		global.Logger.Error("UpdateUserInfo error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UpdateUserInfo success: %s\n", params.Username)
	global.Logger.Info("UpdateUserInfo success: ", zap.String("info", params.Username))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *CUserInfo) UploadUserAvatar(ctx *gin.Context) {
	var params vo.UploadUserAvatarInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UploadUserAvatarInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, data, err := services.UserInfo().UploadUserAvatar(ctx, &params)
	if err != nil {
		fmt.Printf("UploadUserAvatar error: %s\n", err.Error())
		global.Logger.Error("UploadUserAvatar error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UploadUserAvatar success: %s\n", data)
	global.Logger.Info("UploadUserAvatar success: ", zap.String("info", data))
	response.SuccessResponse(ctx, codeStatus, data)
}
