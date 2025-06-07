package controllers

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

type CPayment struct {
}

var Payment = new(CPayment)

func (c *CPayment) CreatePaymentURL(ctx *gin.Context) {
	validation, exists := ctx.Get("validation")
	if !exists {
		fmt.Printf("CreatePaymentURL validation not found")
		global.Logger.Error("CreatePaymentURL validation not found")
		response.ErrorResponse(ctx, response.ErrCodeValidatorNotFound, nil)
		return
	}

	var params vo.CreatePaymentURLInput
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Printf("CreatePaymentURL binding error: %s\n", err.Error())
		global.Logger.Error("CreatePaymentURL binding error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeParamsInvalid, nil)
		return
	}

	err := validation.(*validator.Validate).Struct(params)
	if err != nil {
		fmt.Printf("CreatePaymentURL validation error: %s\n", err.Error())
		global.Logger.Error("CreatePaymentURL validation error: ", zap.String("error", err.Error()))
		response.ErrorResponse(ctx, response.ErrCodeValidator, err.Error())
		return
	}

	services.Payment().CreatePaymentURL(ctx, &params)
}

func (c *CPayment) VNPayReturn(ctx *gin.Context) {
	services.Payment().VNPayReturn(ctx)
}

func (c *CPayment) VNPayIPN(ctx *gin.Context) {
	services.Payment().VNPayIPN(ctx)
}

func (c *CPayment) PostQueryDR(ctx *gin.Context) {

}

func (c *CPayment) PostRefund(ctx *gin.Context) {

}
