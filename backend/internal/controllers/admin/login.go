package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
)

type CAdminLogin struct{}

func (c *CAdminLogin) Register(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.AdminRegisterInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.AdminRegisterInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, err := services.AdminLogin().Register(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CAdminLogin) Login(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.AdminLoginInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.AdminLoginInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.AdminLogin().Login(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	// Create a safe version of response data for logging (exclude sensitive info)
	logData := map[string]interface{}{
		"userAccount": params.UserAccount,
		"loginAt":     utils.GetCurrentUTCTimestamp(),
	}

	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}
