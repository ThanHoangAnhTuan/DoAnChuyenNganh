package upload

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

func (c *Controller) UploadImages(ctx *gin.Context) {
	var params vo.UploadImages
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.UploadImages) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}
	codeStatus, data, err := services.Upload().UploadImages(ctx, &params)
	if err != nil {
		fmt.Printf("UploadImages error: %s\n", err.Error())
		global.Logger.Error("UploadImages error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("UploadImages success: %s", params.ID)
	global.Logger.Info("UploadImages success", zap.String("info", params.ID))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) GetImages(ctx *gin.Context) {
	var params vo.GetImagesInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.GetImagesInput) error {
		if err := ctx.ShouldBindUri(p); err != nil {
			return err
		}
		return ctx.ShouldBindQuery(p)
	}); err != nil {
		return
	}

	codeStatus, data, err := services.Upload().GetImages(ctx, &params)
	if err != nil {
		fmt.Printf("GetImages error: %s\n", err.Error())
		global.Logger.Error("GetImages error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("GetImages success")
	global.Logger.Info("GetImages success")
	response.SuccessResponse(ctx, codeStatus, data)
}
