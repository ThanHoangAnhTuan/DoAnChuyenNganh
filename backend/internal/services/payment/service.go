package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	CreatePaymentURL(ctx *gin.Context, in *vo.CreatePaymentURLInput) (codeStatus int, out *vo.CreatePaymentURLOutput, err error)
	VNPayReturn(ctx *gin.Context) (codeStatus int, err error)
	VNPayIPN(ctx *gin.Context)
	PostQueryDR(ctx *gin.Context, in *vo.PostQueryDRInput)
	PostRefund(ctx *gin.Context, in *vo.PostRefundInput)
}
