package stats

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type StatsRouter struct {
}

func (r StatsRouter) InitStatsRouter(Router *gin.RouterGroup) {
	// statsRouterPublic := Router.Group("/stats")
	// {
	// 	statsRouterPublic.GET("/", controllers.Stats)
	// }

	statsRouterPublic := Router.Group("/stats")
	statsRouterPublic.Use(middlewares.AuthMiddleware())
	{
		// GET /stats => thống kê các tháng trong năm hiện tại
		statsRouterPublic.GET("", controllers.Stats.GetMonthlyEarnings)

		// GET /stats/daily => thống kê các ngày trong tháng hiện tại
		statsRouterPublic.GET("/daily", controllers.Stats.GetDailyEarnings)

		// GET /stats/daily/:year/:month => thống kê các ngày trong tháng và năm cụ thể
		statsRouterPublic.GET("/daily/:year/:month", controllers.Stats.GetDailyEarningsByMonth)

		// GET /stats/monthly/:year => thống kê các tháng trong năm cụ thể
		statsRouterPublic.GET("/monthly/:year", controllers.Stats.GetMonthlyEarningsByYear)

		statsRouterPublic.GET("/export/daily-earnings/csv/:year/:month", controllers.Stats.ExportDailyEarningsCSV)

		statsRouterPublic.GET("/export/monthly-earnings/csv/:year", controllers.Stats.ExportMonthlyEarningsCSV)
	}
}
