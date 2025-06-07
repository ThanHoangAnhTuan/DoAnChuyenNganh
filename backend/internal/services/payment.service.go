package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type IPayment interface {
	CreatePaymentURL(ctx *gin.Context, in *vo.CreatePaymentURLInput) (codeStatus int, err error)
	VNPayReturn(ctx *gin.Context) (codeStatus int, err error)
	VNPayIPN(ctx *gin.Context)
	PostQueryDR(ctx *gin.Context, in *vo.PostQueryDRInput)
	PostRefund(ctx *gin.Context, in *vo.PostRefundInput)
}

var (
	localPayment IPayment
)

func Payment() IPayment {
	if localPayment == nil {
		panic("Implement localPayment not found for interface IPayment")
	}
	return localPayment
}

func InitPayment(i IPayment) {
	localPayment = i
}
