package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
)

type CAdminManager struct{}

func (c *CAdminManager) GetManagers(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.GetManagerInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetManagerInput) error {
		return ctx.ShouldBindQuery(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, pagination, err := services.AdminManager().GetManagers(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}

func (c *CAdminManager) GetAccommodationsOfManager(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.GetAccommodationsOfManagerInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationsOfManagerInput) error {
		if err := ctx.ShouldBindUri(p); err != nil {
			return err
		}
		return ctx.ShouldBindQuery(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, pagination, err := services.AdminManager().GetAccommodationsOfManager(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}

func (c *CAdminManager) VerifyAccommodation(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.VerifyAccommodationInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.VerifyAccommodationInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, err := services.AdminManager().VerifyAccommodation(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *CAdminManager) SetDeletedAccommodation(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.SetDeletedAccommodationInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.SetDeletedAccommodationInput) error {
		return ctx.ShouldBindJSON(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, err := services.AdminManager().SetDeletedAccommodation(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	// Create a log-safe version of the response data
	logData := map[string]interface{}{
		"accommodationId": params.AccommodationID,
		"deletedAt":       utils.GetCurrentUTCTimestamp(),
	}

	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}
