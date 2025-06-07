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
		statsRouterPublic.POST("/", controllers.Stats.GetMonthlyEarnings)
	}
}
