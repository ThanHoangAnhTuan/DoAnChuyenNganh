package impl

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
	"go.uber.org/zap"
)

type OrderImpl struct {
	sqlc *database.Queries
	db   *sql.DB
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

func (o *OrderImpl) CreateOrder(ctx *gin.Context, in *vo.CreateOrderInput) (codeStatus int, out *vo.CreateOrderOutput, err error) {
	// TODO: get userID from context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: check user exists
	isUserBaseExists, err := o.sqlc.CheckUserBaseExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeGetUserBaseFailed, nil, fmt.Errorf("get user base failed: %s", err)
	}

	if !isUserBaseExists {
		return response.ErrCodeUserBaseNotFound, nil, fmt.Errorf("user base not found")
	}

	// TODO: check accommodation id
	isAccommodationExists, err := o.sqlc.CheckAccommodationExists(ctx, in.AccommodationID)
	if err != nil {
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("get accommodation failed: %s", err)
	}

	if !isAccommodationExists {
		return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
	}

	// TODO: check accommodation detail id
	for i := 0; i < len(in.AccommodationDetailID); i++ {
		accommodationDetailID := in.AccommodationDetailID[i]
		isAccommodationDetailExists, err := o.sqlc.CheckAccommodationDetailExists(ctx, accommodationDetailID)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailFailed, nil, fmt.Errorf("get accommodation detail failed: %s", err)
		}

		if !isAccommodationDetailExists {
			return response.ErrCodeAccommodationDetailNotFound, nil, fmt.Errorf("accommodation detail not found: %s", accommodationDetailID)
		}
	}
	// TODO: start transaction
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return response.ErrCodeBeginTransactionFailed, nil, fmt.Errorf("begin transaction failed: %s", err)
	}

	// Ensure rollback if there is an error
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("rollback transaction failed: %s", err)
				global.Logger.Error("rollback transaction failed", zap.Error(rbErr))
			}
		}
	}()

	qtx := database.New(tx)

	// TODO: get accommodation details by ids
	accommodationDetails, err := qtx.GetAccommodationDetailsByIDs(ctx, database.GetAccommodationDetailsByIDsParams{
		Ids:             in.AccommodationDetailID,
		AccommodationID: in.AccommodationID,
	})

	// TODO: create order
	now := utiltime.GetTimeNow()
	orderID := uuid.NewString()

	checkInUnix, err := utiltime.ConvertISOToUnixTimestamp(in.CheckIn)
	if err != nil {
		return response.ErrCodeConvertISOToUnixFailed, nil, fmt.Errorf("convert ISO to Unix failed: %s", err)
	}

	checkOutUnix, err := utiltime.ConvertISOToUnixTimestamp(in.CheckOut)
	if err != nil {
		return response.ErrCodeConvertISOToUnixFailed, nil, fmt.Errorf("convert ISO to Unix failed: %s", err)
	}

	var totalPrice uint32
	totalPrice = 0

	for _, accommodationDetail := range accommodationDetails {
		// TODO: calculate price by:
		// TODO: price = accommodation price - discount price

		totalPrice += accommodationDetail.Price
	}

	err = qtx.CreateOrder(ctx, database.CreateOrderParams{
		ID:              orderID,
		UserID:          userID,
		AccommodationID: in.AccommodationID,
		OrderStatus:     database.EcommerceGoOrderOrderStatusCompleted,
		CheckinDate:     checkInUnix,
		CheckoutDate:    checkOutUnix,
		CreatedAt:       now,
		UpdatedAt:       now,
		FinalTotal:      totalPrice,
		VoucherID: sql.NullString{
			Valid:  false,
			String: "",
		},
	})

	if err != nil {
		return response.ErrCodeCreateOrderFailed, nil, fmt.Errorf("create order failed: %s", err)
	}

	for _, accommodationDetail := range accommodationDetails {
		orderDetailId := uuid.NewString()
		// TODO: calculate price by:
		// TODO: price = accommodation price - discount price

		// TODO: create order detail
		err := qtx.CreateOrderDetail(ctx, database.CreateOrderDetailParams{
			ID:                    orderDetailId,
			OrderID:               orderID,
			Price:                 accommodationDetail.Price,
			AccommodationDetailID: accommodationDetail.ID,
			CreatedAt:             now,
			UpdatedAt:             now,
		})

		if err != nil {
			return response.ErrCodeCreateOrderDetailFailed, nil, fmt.Errorf("create order detail failed: id=%s, err=%s", orderDetailId, err)
		}
	}

	// TODO: create payment
	paymentId := uuid.NewString()
	err = qtx.CreatePayment(ctx, database.CreatePaymentParams{
		ID:            paymentId,
		OrderID:       orderID,
		PaymentStatus: database.EcommerceGoPaymentPaymentStatusSuccess,
		PaymentMethod: database.EcommerceGoPaymentPaymentMethodCard,
		TotalPrice:    totalPrice,
		TransactionID: sql.NullString{
			Valid:  false,
			String: "",
		},
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return response.ErrCodeCreatePaymentFailed, nil, fmt.Errorf("create payment failed: %s", err)
	}

	// TODO: return out
	out = &vo.CreateOrderOutput{}
	out.CheckIn = in.CheckIn
	out.CheckOut = in.CheckOut
	out.OrderID = orderID
	out.OrderStatus = string(database.EcommerceGoOrderOrderStatusCompleted)
	out.PaymentMethod = string(database.EcommerceGoPaymentPaymentMethodCard)
	out.TotalPrice = utils.FormatCurrency(totalPrice)

	orderDate, err := utiltime.ConvertUnixTimestampToISO(ctx, int64(now))
	if err != nil {
		return response.ErrCodeConvertUnixToISOFailed, nil, fmt.Errorf("convert unix to iso failed: %s", err)
	}
	out.OrderDate = orderDate

	for _, accommodationDetail := range accommodationDetails {
		out.OrderDetails = append(out.OrderDetails, vo.CreateOrderDetailOutput{
			AccommodationDetailName:  accommodationDetail.Name,
			AccommodationDetailPrice: utils.FormatCurrency(accommodationDetail.Price),
		})
	}
	if err = tx.Commit(); err != nil {
		return response.ErrCodeCommitTransactionFailed, nil, fmt.Errorf("commit transaction failed: %s", err)
	}
	// TODO: end transaction

	return response.ErrCodeCreateOrderSuccess, out, nil

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
