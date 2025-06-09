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
type GetOrdersByUserOutput struct {
}
type GetOrdersByManagerInput struct {
}
type GetOrdersByManagerOutput struct {
}
type CancelOrderInput struct {
}
type CancelOrderOutput struct {
}
type CheckInInput struct {
}
type CheckInOutput struct {
}
type CheckOutInput struct {
}
type CheckOutOutput struct {
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
