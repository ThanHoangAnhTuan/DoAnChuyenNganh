package order

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type serviceImpl struct {
	sqlc *database.Queries
	db   *sql.DB
}

func New(sqlc *database.Queries, db *sql.DB) Service {
	return &serviceImpl{
		sqlc: sqlc,
		db:   db,
	}
}

func (o *serviceImpl) CancelOrder(ctx *gin.Context, in *vo.CancelOrderInput) (codeStatus int, err error) {
	// TODO: get userId from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("userID not found in context")
	}

	// TODO: check user exists
	exists, err := o.sqlc.CheckUserBaseExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("get user base failed: %s", err)
	}

	if !exists {
		return response.ErrCodeUnauthorized, fmt.Errorf("user base not found")
	}

	// TODO: check order exists
	isExists, err := o.sqlc.CheckOrderExists(ctx, in.OrderID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("check order exists failed: %s", err)
	}

	if !isExists {
		return response.ErrCodeOrderNotFound, fmt.Errorf("order not found")
	}

	// TODO: kiểm tra xem người dùng xoá quyền huỷ order hay không
	// TODO: 1. trước thời gian quy định của khách sạn

	// TODO: update status of order
	now := utiltime.GetTimeNow()
	err = o.sqlc.UpdateOrderStatusByID(ctx, database.UpdateOrderStatusByIDParams{
		OrderStatus: database.EcommerceGoOrderOrderStatusCanceled,
		UpdatedAt:   now,
		ID:          in.OrderID,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("update order status failed: %s", err)
	}

	return response.ErrCodeUpdateOrderSuccess, nil
}

func (o *serviceImpl) CheckIn(ctx *gin.Context, in *vo.CheckInInput) (codeStatus int, err error) {
	// TODO: get userId from context
	managerID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("managerID not found in context")
	}

	// TODO: check user exists
	exists, err := o.sqlc.CheckUserManagerExistsByID(ctx, managerID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("get user base failed: %s", err)
	}

	if !exists {
		return response.ErrCodeUnauthorized, fmt.Errorf("user base not found")
	}

	// TODO: check order exists
	isExists, err := o.sqlc.CheckOrderExists(ctx, in.OrderID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("check order exists failed: %s", err)
	}

	if !isExists {
		return response.ErrCodeOrderNotFound, fmt.Errorf("order not found")
	}

	// TODO: check manager có được đổi status của order hay không
	// TODO: 1. thời gian update chính là thời gian checkin

	// TODO: update status of order
	now := utiltime.GetTimeNow()
	err = o.sqlc.UpdateOrderStatusByID(ctx, database.UpdateOrderStatusByIDParams{
		OrderStatus: database.EcommerceGoOrderOrderStatusCheckedIn,
		UpdatedAt:   now,
		ID:          in.OrderID,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("update order status failed: %s", err)
	}

	return response.ErrCodeUpdateOrderSuccess, nil
}

func (o *serviceImpl) CheckOut(ctx *gin.Context, in *vo.CheckOutInput) (codeStatus int, err error) {
	// TODO: get userId from context
	managerID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, fmt.Errorf("managerID not found in context")
	}

	// TODO: check user exists
	exists, err := o.sqlc.CheckUserManagerExistsByID(ctx, managerID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("get user base failed: %s", err)
	}

	if !exists {
		return response.ErrCodeUnauthorized, fmt.Errorf("user base not found")
	}

	// TODO: check order exists
	isExists, err := o.sqlc.CheckOrderExists(ctx, in.OrderID)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("check order exists failed: %s", err)
	}

	if !isExists {
		return response.ErrCodeOrderNotFound, fmt.Errorf("order not found")
	}

	// TODO: check manager có được đổi status của order hay không
	// TODO: 1. thời gian update chính là thời gian checkout

	// TODO: update status of order
	now := utiltime.GetTimeNow()
	err = o.sqlc.UpdateOrderStatusByID(ctx, database.UpdateOrderStatusByIDParams{
		OrderStatus: database.EcommerceGoOrderOrderStatusCompleted,
		UpdatedAt:   now,
		ID:          in.OrderID,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("update order status failed: %s", err)
	}

	return response.ErrCodeUpdateOrderSuccess, nil
}

func (o *serviceImpl) GetOrderInfoAfterPayment(ctx *gin.Context, in *vo.GetOrderInfoAfterPaymentInput) (codeStatus int, out *vo.GetOrderInfoAfterPaymentOutput, err error) {
	out = &vo.GetOrderInfoAfterPaymentOutput{}

	// TODO: get order info by order id extenal
	order, err := o.sqlc.GetOrderInfoByOrderIDExternal(ctx, in.OrderIDExternal)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
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
		return response.ErrCodeInternalServerError, nil, err
	}

	orderDate, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CreatedAt))
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	checkIn, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CheckinDate))
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	checkOut, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CheckoutDate))
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	out.CheckIn = checkIn
	out.CheckOut = checkOut
	out.OrderDate = orderDate
	out.OrderIDExternal = order.OrderIDExternal
	out.OrderStatus = string(order.OrderStatus)
	out.TotalPrice = payment.TotalPrice.String()
	out.TransactionID = payment.TransactionID.String

	// TODO: get username from user info by user id
	username, err := o.sqlc.GetUsernameByID(ctx, order.UserID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, err
	}

	out.Username = username

	return response.ErrCodeGetOrderSuccess, out, nil
}

func (o *serviceImpl) GetOrdersByManager(ctx *gin.Context) (codeStatus int, out []*vo.GetOrdersByManagerOutput, err error) {
	out = []*vo.GetOrdersByManagerOutput{}

	// TODO: get userId from context
	managerID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("managerID not found in context")
	}

	// TODO: check user exists
	exists, err := o.sqlc.CheckUserManagerExistsByID(ctx, managerID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get user base failed: %s", err)
	}

	if !exists {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user base not found")
	}
	orders, err := o.sqlc.GetOrdersByManager(ctx, managerID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get order failed: %s", err)
	}

	for _, order := range orders {
		// TODO: get order detail:
		orderDetails, err := o.sqlc.GetOrderDetailsByManager(ctx, order.OrderID)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("get order detail failed: %s", err)
		}
		detail := []vo.OrderDetailOutput{}
		for _, orderDetail := range orderDetails {
			// TODO: get room bookings for this order detail
			roomBookings, err := o.sqlc.GetOrderRoomBookingsByOrderDetailIDWithRoomInfo(ctx, orderDetail.OrderDetailID)
			if err != nil {
				return response.ErrCodeInternalServerError, nil, fmt.Errorf("get room bookings failed: %s", err)
			}

			roomBookingsList := []vo.RoomBookingOutput{}
			for _, roomBooking := range roomBookings {
				roomBookingsList = append(roomBookingsList, vo.RoomBookingOutput{
					ID:                  roomBooking.ID,
					AccommodationRoomID: roomBooking.AccommodationRoomID,
					RoomName:            roomBooking.RoomName,
					BookingStatus:       string(roomBooking.BookingStatus),
				})
			}

			detail = append(detail, vo.OrderDetailOutput{
				AccommodationDetailID:   orderDetail.AccommodationDetailID,
				AccommodationDetailName: orderDetail.AccommodationDetailName,
				Price:                   orderDetail.Price.String(),
				RoomBookings:            roomBookingsList,
			})
		}

		checkIn, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CheckinDate))
		if err != nil {
			return response.ErrCodeInternalServerError, nil, err
		}

		checkOut, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CheckoutDate))
		if err != nil {
			return response.ErrCodeInternalServerError, nil, err
		}

		out = append(out, &vo.GetOrdersByManagerOutput{
			ID:                order.OrderID,
			FinalTotal:        order.FinalTotal.String(),
			OrderStatus:       string(order.OrderStatus),
			AccommodationID:   order.AccommodationID,
			AccommodationName: order.AccommodationName,
			CheckIn:           checkIn,
			CheckOut:          checkOut,
			OrderDetail:       detail,
			Email:             order.Email,
			Username:          order.Username,
			Phone:             order.Phone.String,
		})
	}
	return response.ErrCodeGetOrderSuccess, out, nil
}

func (o *serviceImpl) GetOrdersByUser(ctx *gin.Context) (codeStatus int, out []*vo.GetOrdersByUserOutput, err error) {
	out = []*vo.GetOrdersByUserOutput{}

	// TODO: get userId from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user exists
	exists, err := o.sqlc.CheckUserBaseExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get user base failed: %s", err)
	}

	if !exists {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("user base not found")
	}
	orders, err := o.sqlc.GetOrdersByUser(ctx, userID)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get order failed: %s", err)
	}

	for _, order := range orders {
		// TODO: get order detail:
		orderDetails, err := o.sqlc.GetOrderDetailsByUser(ctx, order.OrderID)
		if err != nil {
			return response.ErrCodeInternalServerError, nil, fmt.Errorf("get order detail failed: %s", err)
		}
		detail := []vo.OrderDetailOutput{}
		for _, orderDetail := range orderDetails {
			detail = append(detail, vo.OrderDetailOutput{
				AccommodationDetailID:   orderDetail.AccommodationDetailID,
				AccommodationDetailName: orderDetail.AccommodationDetailName,
				Price:                   orderDetail.Price.String(),
				Guests:                  orderDetail.Guests,
			})
		}

		checkIn, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CheckinDate))
		if err != nil {
			return response.ErrCodeInternalServerError, nil, err
		}
		checkOut, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(order.CheckoutDate))
		if err != nil {
			return response.ErrCodeInternalServerError, nil, err
		}

		out = append(out, &vo.GetOrdersByUserOutput{
			ID:                order.OrderID,
			FinalTotal:        order.FinalTotal.String(),
			OrderStatus:       string(order.OrderStatus),
			AccommodationID:   order.AccommodationID,
			AccommodationName: order.AccommodationName,
			CheckIn:           checkIn,
			CheckOut:          checkOut,
			OrderDetail:       detail,
		})
	}
	return response.ErrCodeGetOrderSuccess, out, nil
}
