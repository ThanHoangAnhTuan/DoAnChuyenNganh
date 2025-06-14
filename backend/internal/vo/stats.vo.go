package vo

type GetMonthlyEarningsOutput struct {
	Month        int32  `json:"month"`
	TotalOrders  int64  `json:"total_orders"`
	TotalRevenue string `json:"total_revenue"`
}

type GetDailyEarningsOutput struct {
	Day          string `json:"day"`
	TotalOrders  int64  `json:"total_orders"`
	TotalRevenue string `json:"total_revenue"`
}

type GetYearlyEarningsOutput struct {
	Year         int32  `json:"year"`
	TotalOrders  int64  `json:"total_orders"`
	TotalRevenue string `json:"total_revenue"`
}

type GetDailyEarningsByMonthInput struct {
	Year  int `uri:"year"`
	Month int `uri:"month"`
}

type GetMonthlyEarningsByYearInput struct {
	Year int `uri:"year"`
}
