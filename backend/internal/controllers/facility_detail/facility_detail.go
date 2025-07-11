package facility_detail

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

func (c *Controller) CreateFacilityDetail(ctx *gin.Context) {
	var params vo.CreateFacilityDetailInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreateFacilityDetailInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, data, err := services.FacilityDetail().CreateFacilityDetail(ctx, &params)
	if err != nil {
		fmt.Printf("CreateFacilityDetail error: %s\n", err.Error())
		global.Logger.Error("CreateFacilityDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CreateFacilityDetail success: %s\n", data)
	global.Logger.Info("CreateFacilityDetail success: ", zap.String("info", data.ID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetFacilityDetail(ctx *gin.Context) {
	codeStatus, data, err := services.FacilityDetail().GetFacilityDetail(ctx)
	if err != nil {
		fmt.Printf("GetFacilityDetail error: %s\n", err.Error())
		global.Logger.Error("GetFacilityDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetFacilityDetail success\n")
	global.Logger.Info("GetFacilityDetail success")
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) UpdateFacilityDetail(ctx *gin.Context) {
	var params vo.UpdateFacilityDetailInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UpdateFacilityDetailInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, data, err := services.FacilityDetail().UpdateFacilityDetail(ctx, &params)
	if err != nil {
		fmt.Printf("UpdateFacilityDetail error: %s\n", err.Error())
		global.Logger.Error("UpdateFacilityDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UpdateFacilityDetail success: %s\n", data.ID)
	global.Logger.Info("UpdateFacilityDetail success: ", zap.String("info", fmt.Sprintf("update facility detail success: %s", data.ID)))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) DeleteFacilityDetail(ctx *gin.Context) {
	var params vo.DeleteFacilityDetailInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.DeleteFacilityDetailInput) error {
		return ctx.ShouldBindUri(p)
	}); err != nil {
		return
	}

	codeStatus, err := services.FacilityDetail().DeleteFacilityDetail(ctx, &params)
	if err != nil {
		fmt.Printf("DeleteFacilityDetail error: %s\n", err.Error())
		global.Logger.Error("DeleteFacilityDetail error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("DeleteFacilityDetail success: %s\n", params.ID)
	global.Logger.Info("DeleteFacilityDetail success: ", zap.String("info", fmt.Sprintf("delete facility detail success: %s", params.ID)))
	response.SuccessResponse(ctx, codeStatus, nil)
}
