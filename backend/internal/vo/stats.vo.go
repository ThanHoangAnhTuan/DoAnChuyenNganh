package vo

type GetMonthlyEarningsOuput struct {
	Month   string `json:"month"`
	Revenue int64  `json:"revenue"`
}
