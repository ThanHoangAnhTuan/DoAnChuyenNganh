package payment

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/services"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"go.uber.org/zap"
)

func (c *Controller) CreatePaymentURL(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreatePaymentURL validation not found\n")
		global.Logger.Error("CreatePaymentURL validation not found")
		response.ErrorResponse(ctx, response.ErrCodeInternalServerError, nil)
		return
	}

	var params vo.CreatePaymentURLInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreatePaymentURL binding error: %s\n", err.Error())
		global.Logger.Error("CreatePaymentURL binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		validationErrors := response.FormatValidationErrorsToStruct(err, params)
		fmt.Printf("CreatePaymentURL validation error: %s\n", validationErrors)
		global.Logger.Error("CreatePaymentURL validation error: ", zap.Any("error", validationErrors))
		response.ErrorResponse(ctx, response.ErrCodeValidator, validationErrors)
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
