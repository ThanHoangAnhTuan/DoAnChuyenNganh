package facility

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/controllerutil"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
)

func (c *Controller) CreateFacility(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.CreateFacilityInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateFacilityInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.Facility().CreateFacility(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	logData := map[string]interface{}{
		"facilityId": data.ID,
		"name":       data.Name,
		"createdAt":  utils.GetCurrentUTCTimestamp(),
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) UpdateFacility(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.UpdateFacilityInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdateFacilityInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, data, err := services.Facility().UpdateFacility(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	logData := map[string]interface{}{
		"facilityId": data.ID,
		"name":       data.Name,
		"updatedAt":  utils.GetCurrentUTCTimestamp(),
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteFacility(ctx *gin.Context) {
	start := utils.NowMs()
	var params vo.DeleteFacilityInput

	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.DeleteFacilityInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", response.ErrCodeValidator, utils.NowMs()-start, err)
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	codeStatus, err := services.Facility().DeleteFacility(ctx, &params)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, params, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	logData := map[string]interface{}{
		"facilityId": params.ID,
		"deletedAt":  utils.GetCurrentUTCTimestamp(),
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, nil)
}

func (c *Controller) GetFacilities(ctx *gin.Context) {
	start := utils.NowMs()

	codeStatus, data, err := services.Facility().GetFacilities(ctx)
	duration := utils.NowMs() - start

	if err != nil {
		controllerutil.HandleStructuredLog(ctx, nil, "error", codeStatus, duration, err)
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	logData := map[string]interface{}{
		"count": len(data),
	}
	controllerutil.HandleStructuredLog(ctx, logData, "success", codeStatus, duration, nil)
	response.SuccessResponse(ctx, codeStatus, data)
}
