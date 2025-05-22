package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
	"go.uber.org/zap"
)

type OrderImpl struct {
	sqlc *database.Queries
	db   *sql.DB
}

func (o *OrderImpl) CancelOrder(ctx context.Context, in *vo.CancelOrderInput) (codeStatus int, out *vo.CancelOrderOutput, err error) {
	panic("unimplemented")
}

func (o *OrderImpl) CheckIn(ctx context.Context, in *vo.CheckInInput) (codeStatus int, out *vo.CheckInOutput, err error) {
	panic("unimplemented")
}

func (o *OrderImpl) CheckOut(ctx context.Context, in *vo.CheckOutInput) (codeStatus int, out *vo.CheckOutOutput, err error) {
	panic("unimplemented")
}

func (o *OrderImpl) CreateOrder(ctx context.Context, in *vo.CreateOrderInput) (codeStatus int, out *vo.CreateOrderOutput, err error) {
	// TODO: get user id in context.Context
	val := ctx.Value("userId")
	if val == nil {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("unauthorized")
	}
	userID, ok := val.(string)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("invalid user id format")
	}

	// TODO: check user id
	isUserBaseExists, err := o.sqlc.CheckUserBaseExistsById(ctx, userID)
	if err != nil {
		return response.ErrCodeGetUserBaseFailed, nil, fmt.Errorf("get user base failed: %s", err)
	}

	fmt.Printf("isUserBaseExists: %v", isUserBaseExists)

	if !isUserBaseExists {
		return response.ErrCodeUserBaseNotFound, nil, fmt.Errorf("user base not found")
	}

	// TODO: check accommodation id
	isAccommodationExists, err := o.sqlc.CheckAccommodationExists(ctx, in.AccommodationId)
	if err != nil {
		return response.ErrCodeGetAccommodationFailed, nil, fmt.Errorf("get accommodation failed: %s", err)
	}

	if !isAccommodationExists {
		return response.ErrCodeAccommodationNotFound, nil, fmt.Errorf("accommodation not found")
	}

	// TODO: check accommodation detail id
	for i := 0; i < len(in.AccommodationDetailId); i++ {
		accommodationDetailId := in.AccommodationDetailId[i]
		isAccommodationDetailExists, err := o.sqlc.CheckAccommodationDetailExists(ctx, accommodationDetailId)
		if err != nil {
			return response.ErrCodeGetAccommodationDetailFailed, nil, fmt.Errorf("get accommodation detail failed: %s", err)
		}

		if !isAccommodationDetailExists {
			return response.ErrCodeAccommodationDetailNotFound, nil, fmt.Errorf("accommodation detail not found: %s", accommodationDetailId)
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
		Ids:             in.AccommodationDetailId,
		AccommodationID: in.AccommodationId,
	})

	// TODO: create order
	now := utiltime.GetTimeNow()
	orderId := uuid.NewString()

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
		ID:              orderId,
		UserID:          userID,
		AccommodationID: in.AccommodationId,
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
			OrderID:               orderId,
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
		OrderID:       orderId,
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
	out.OrderId = orderId
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

func (o *OrderImpl) GetOrdersByManager(ctx context.Context, in *vo.GetOrdersByManagerInput) (codeStatus int, out *vo.GetOrdersByManagerOutput, err error) {
	panic("unimplemented")
}

func (o *OrderImpl) GetOrdersByUser(ctx context.Context, in *vo.GetOrdersByUserInput) (codeStatus int, out *vo.GetOrdersByUserOutput, err error) {
	panic("unimplemented")
}

func NewOrderImpl(sqlc *database.Queries, db *sql.DB) *OrderImpl {
	return &OrderImpl{
		sqlc: sqlc,
		db:   db,
	}
}
