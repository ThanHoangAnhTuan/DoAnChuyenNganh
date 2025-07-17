package vo

type CreateOrderInput struct {
	AccommodationID       string   `json:"accommodation_id" validate:"required"`
	AccommodationDetailID []string `json:"accommodation_detail_id" validate:"required"`
	CheckIn               string   `json:"check_in" validate:"required"`
	CheckOut              string   `json:"check_out" validate:"required"`
}

type CreateOrderDetailOutput struct {
	AccommodationDetailName  string `json:"accommodation_detail_name"`
	AccommodationDetailPrice string `json:"accommodation_detail_price"`
}

type CreateOrderOutput struct {
	OrderID       string                    `json:"order_id"`
	TotalPrice    string                    `json:"total_price"`
	OrderStatus   string                    `json:"order_status"`
	CheckIn       string                    `json:"check_in"`
	CheckOut      string                    `json:"check_out"`
	OrderDetails  []CreateOrderDetailOutput `json:"order_details"`
	PaymentMethod string                    `json:"payment_method"`
	OrderDate     string                    `json:"order_date"`
}

type GetOrdersByUserInput struct {
}

type OrderDetailOutput struct {
	AccommodationDetailID   string              `json:"accommodation_detail_id"`
	AccommodationDetailName string              `json:"accommodation_detail_name"`
	Price                   string              `json:"price"`
	Guests                  uint8               `json:"guests"`
	RoomBookings            []RoomBookingOutput `json:"room_bookings,omitempty"`
}

type RoomBookingOutput struct {
	ID                  string `json:"id"`
	AccommodationRoomID string `json:"accommodation_room_id"`
	RoomName            string `json:"room_name"`
	BookingStatus       string `json:"booking_status"`
}

type GetOrdersByUserOutput struct {
	ID                string              `json:"id"`
	FinalTotal        string              `json:"final_total"`
	OrderStatus       string              `json:"order_status"`
	AccommodationID   string              `json:"accommodation_id"`
	AccommodationName string              `json:"accommodation_name"`
	CheckIn           string              `json:"check_in"`
	CheckOut          string              `json:"check_out"`
	OrderDetail       []OrderDetailOutput `json:"order_detail"`
}

type GetOrdersByManagerInput struct {
}
type GetOrdersByManagerOutput struct {
	ID                string              `json:"id"`
	FinalTotal        string              `json:"final_total"`
	OrderStatus       string              `json:"order_status"`
	AccommodationID   string              `json:"accommodation_id"`
	AccommodationName string              `json:"accommodation_name"`
	CheckIn           string              `json:"check_in"`
	CheckOut          string              `json:"check_out"`
	OrderDetail       []OrderDetailOutput `json:"order_detail"`
	Email             string              `json:"email"`
	Username          string              `json:"username"`
	Phone             string              `json:"phone"`
}
type CancelOrderInput struct {
	OrderID string `json:"order_id"`
}

type CheckInInput struct {
	OrderID string `json:"order_id"`
}

type CheckOutInput struct {
	OrderID string `json:"order_id"`
}

type GetOrderInfoAfterPaymentInput struct {
	OrderIDExternal string `json:"order_id"`
	TransactionID   string `json:"transaction_id"`
}

type GetOrderInfoAfterPaymentOutput struct {
	OrderIDExternal string `json:"order_id"`
	OrderStatus     string `json:"order_status"`
	TotalPrice      string `json:"total_price"`
	CheckIn         string `json:"check_in"`
	CheckOut        string `json:"check_out"`
	OrderDate       string `json:"order_date"`
	Username        string `json:"username"`
	TransactionID   string `json:"transaction_id"`
}
