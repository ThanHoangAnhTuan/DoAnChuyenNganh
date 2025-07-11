package payment

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

func (c *Controller) CreatePaymentURL(ctx *gin.Context) {
	var params vo.CreatePaymentURLInput
	if err := controllerutil.BindAndValidate(ctx, &params, func(p *vo.CreatePaymentURLInput) error {
		return ctx.ShouldBind(p)
	}); err != nil {
		return
	}

	codeStatus, data, err := services.Payment().CreatePaymentURL(ctx, &params)
	if err != nil {
		fmt.Printf("CreatePaymentURL error: %s\n", err.Error())
		global.Logger.Error("CreatePaymentURL error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, codeStatus, nil)
		return
	}

	fmt.Printf("CreatePaymentURL success: %s\n", data.Url)
	global.Logger.Info("CreatePaymentURL success", zap.String("info", data.Url))
	response.SuccessResponse(ctx, codeStatus, data)
}

func (c *Controller) VNPayReturn(ctx *gin.Context) {
	services.Payment().VNPayReturn(ctx)
}

func (c *Controller) VNPayIPN(ctx *gin.Context) {
	services.Payment().VNPayIPN(ctx)
}

func (c *Controller) PostQueryDR(ctx *gin.Context) {

}

func (c *Controller) PostRefund(ctx *gin.Context) {

}
