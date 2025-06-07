package services

import (
	"context"

	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IOrder interface {
		GetOrdersByUser(ctx context.Context, in *vo.GetOrdersByUserInput) (codeStatus int, out *vo.GetOrdersByUserOutput, err error)
		GetOrdersByManager(ctx context.Context, in *vo.GetOrdersByManagerInput) (codeStatus int, out *vo.GetOrdersByManagerOutput, err error)
		CancelOrder(ctx context.Context, in *vo.CancelOrderInput) (codeStatus int, out *vo.CancelOrderOutput, err error)
		CheckIn(ctx context.Context, in *vo.CheckInInput) (codeStatus int, out *vo.CheckInOutput, err error)
		CheckOut(ctx context.Context, in *vo.CheckOutInput) (codeStatus int, out *vo.CheckOutOutput, err error)
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
