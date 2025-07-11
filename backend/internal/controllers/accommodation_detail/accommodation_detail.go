package accommodation_detail

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
)

func (c *Controller) CreateAccommodationDetail(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.CreateAccommodationDetailInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateAccommodationDetailInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().CreateAccommodationDetail(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodationDetails(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.GetAccommodationDetailsInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationDetailsInput) error {
		if err := ctx.ShouldBindUri(p); err != nil {
			return err
		}
		return ctx.ShouldBindQuery(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().GetAccommodationDetails(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetAccommodationDetailsByManager(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.GetAccommodationDetailsByManagerInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetAccommodationDetailsByManagerInput) error {
		if err := ctx.ShouldBindUri(p); err != nil {
			return err
		}
		return ctx.ShouldBindQuery(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().GetAccommodationDetailsByManager(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) UpdateAccommodationDetail(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.UpdateAccommodationDetailInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdateAccommodationDetailInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.AccommodationDetail().UpdateAccommodationDetail(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteAccommodationDetail(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.DeleteAccommodationDetailInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.DeleteAccommodationDetailInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, err := services.AccommodationDetail().DeleteAccommodationDetail(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	controllerutil.HandleStructuredLog(ctx, params, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}
