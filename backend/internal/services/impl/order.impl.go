package impl

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type OrderImpl struct {
	sqlc *database.Queries
	db   *sql.DB
}

func (o *OrderImpl) GetOrderInfoAfterPayment(ctx *gin.Context, in *vo.GetOrderInfoAfterPaymentInput) (codeStatus int, out *vo.GetOrderInfoAfterPaymentOutput, err error) {
	out = &vo.GetOrderInfoAfterPaymentOutput{}

	// TODO: get order info by order id extenal
	order, err := o.sqlc.GetOrderInfoByOrderIDExternal(ctx, in.OrderIDExternal)
	if err != nil {
		return response.ErrCodeGetOrderFailed, nil, err
	}

	// TODO: get payment info by transaction id
	payment, err := o.sqlc.GetPaymentInfo(ctx, database.GetPaymentInfoParams{
		OrderID: order.ID,
		TransactionID: sql.NullString{
			String: in.TransactionID,
			Valid:  true,
		},
	})

	if err != nil {
		return response.ErrCodeGetPaymentFailed, nil, err
	}

	checkInOutput, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CheckinDate))
	if err != nil {
		return response.ErrCodeConvertUnixToISOFailed, nil, err
	}

	checkOutOutput, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CheckoutDate))
	if err != nil {
		return response.ErrCodeConvertUnixToISOFailed, nil, err
	}

	orderDate, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CreatedAt))
	if err != nil {
		return response.ErrCodeConvertUnixToISOFailed, nil, err
	}

	out.CheckIn = checkInOutput
	out.CheckOut = checkOutOutput
	out.OrderDate = orderDate
	out.OrderIDExternal = order.OrderIDExternal
	out.OrderStatus = string(order.OrderStatus)
	out.TotalPrice = payment.TotalPrice.String()
	out.TransactionID = payment.TransactionID.String

	// TODO: get username from user info by user id
	username, err := o.sqlc.GetUsernameByID(ctx, order.UserID)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, nil, err
	}

	out.Username = username

	return response.ErrCodeGetOrderSuccess, out, nil
}

func (o *OrderImpl) CancelOrder(ctx *gin.Context, in *vo.CancelOrderInput) (codeStatus int, out *vo.CancelOrderOutput, err error) {
	panic("unimplemented")
}

func (o *OrderImpl) CheckIn(ctx *gin.Context, in *vo.CheckInInput) (codeStatus int, out *vo.CheckInOutput, err error) {
	panic("unimplemented")
}

func (o *OrderImpl) CheckOut(ctx *gin.Context, in *vo.CheckOutInput) (codeStatus int, out *vo.CheckOutOutput, err error) {
	panic("unimplemented")
}

func (o *OrderImpl) GetOrdersByManager(ctx *gin.Context, in *vo.GetOrdersByManagerInput) (codeStatus int, out *vo.GetOrdersByManagerOutput, err error) {
	panic("unimplemented")
}

func (o *OrderImpl) GetOrdersByUser(ctx *gin.Context, in *vo.GetOrdersByUserInput) (codeStatus int, out *vo.GetOrdersByUserOutput, err error) {
	panic("unimplemented")
}

func NewOrderImpl(sqlc *database.Queries, db *sql.DB) *OrderImpl {
	return &OrderImpl{
		sqlc: sqlc,
		db:   db,
	}
}
