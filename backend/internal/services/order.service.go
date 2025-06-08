package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IOrder interface {
		GetOrdersByUser(ctx *gin.Context, in *vo.GetOrdersByUserInput) (codeStatus int, out *vo.GetOrdersByUserOutput, err error)
		GetOrdersByManager(ctx *gin.Context, in *vo.GetOrdersByManagerInput) (codeStatus int, out *vo.GetOrdersByManagerOutput, err error)
		CancelOrder(ctx *gin.Context, in *vo.CancelOrderInput) (codeStatus int, out *vo.CancelOrderOutput, err error)
		CheckIn(ctx *gin.Context, in *vo.CheckInInput) (codeStatus int, out *vo.CheckInOutput, err error)
		CheckOut(ctx *gin.Context, in *vo.CheckOutInput) (codeStatus int, out *vo.CheckOutOutput, err error)
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
