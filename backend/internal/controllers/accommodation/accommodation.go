package accommodation

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
)

func (c *Controller) CreateAccommodation(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.CreateAccommodationInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateAccommodationInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.Accommodation().CreateAccommodation(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodations(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.GetAccommodationsInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationsInput) error {
		return ctx.ShouldBindQuery(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, pagination, err := services.Accommodation().GetAccommodations(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}

func (c *Controller) UpdateAccommodation(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.UpdateAccommodationInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdateAccommodationInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.Accommodation().UpdateAccommodation(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteAccommodation(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.DeleteAccommodationInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.DeleteAccommodationInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, err := services.Accommodation().DeleteAccommodation(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) GetAccommodation(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.GetAccommodationInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationInput) error {
		return ctx.ShouldBindUri(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.Accommodation().GetAccommodation(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodationsByManager(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.GetAccommodationsInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationsInput) error {
		return ctx.ShouldBindQuery(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, pagination, err := services.Accommodation().GetAccommodationsByManager(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponseWithPagination(ctx, codeStatus, data, pagination)
}
