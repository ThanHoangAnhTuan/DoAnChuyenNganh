package stats

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	// Monthly stats cho năm hiện tại
	GetMonthlyEarnings(ctx *gin.Context) (codeStatus int, out []*vo.GetMonthlyEarningsOutput, err error)

	// Daily stats cho tháng hiện tại
	GetDailyEarnings(ctx *gin.Context) (codeStatus int, out []*vo.GetDailyEarningsOutput, err error)

	// Monthly stats cho năm cụ thể
	GetMonthlyEarningsByYear(ctx *gin.Context, in *vo.GetMonthlyEarningsByYearInput) (codeStatus int, out []*vo.GetMonthlyEarningsOutput, err error)

	// Daily stats cho tháng cụ thể
	GetDailyEarningsByMonth(ctx *gin.Context, in *vo.GetDailyEarningsByMonthInput) (codeStatus int, out []*vo.GetDailyEarningsOutput, err error)

	ExportDailyEarningsCSV(ctx *gin.Context, in *vo.GetDailyEarningsByMonthInput) (codeStatus int, csvData []byte, err error)

	ExportMonthlyEarningsCSV(ctx *gin.Context, in *vo.GetMonthlyEarningsByYearInput) (codeStatus int, csvData []byte, err error)
}
