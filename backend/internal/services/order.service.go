package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IOrder interface {
		GetOrdersByUser(ctx *gin.Context) (codeStatus int, out []*vo.GetOrdersByUserOutput, err error)
		GetOrdersByManager(ctx *gin.Context) (codeStatus int, out []*vo.GetOrdersByManagerOutput, err error)
		CancelOrder(ctx *gin.Context, in *vo.CancelOrderInput) (codeStatus int, err error)
		CheckIn(ctx *gin.Context, in *vo.CheckInInput) (codeStatus int, err error)
		CheckOut(ctx *gin.Context, in *vo.CheckOutInput) (codeStatus int, err error)
		GetOrderInfoAfterPayment(ctx *gin.Context, in *vo.GetOrderInfoAfterPaymentInput) (codeStatus int, out *vo.GetOrderInfoAfterPaymentOutput, err error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("Implement localOrder not found for interface IOrder")
	}
	return localOrder
}

func InitOrder(i IOrder) {
	localOrder = i
}
