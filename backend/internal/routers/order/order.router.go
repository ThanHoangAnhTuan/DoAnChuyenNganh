package order

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type OrderRouter struct {
}

func (r *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderRouterPrivate := Router.Group("/order")
	orderRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		orderRouterPrivate.POST("/cancel", controllers.Order.CancelOrder)
		orderRouterPrivate.POST("/checkin", controllers.Order.CheckIn)
		orderRouterPrivate.POST("/checkout", controllers.Order.CheckOut)
		orderRouterPrivate.GET("/get-order-info-after-payment", controllers.Order.GetOrderInfoAfterPayment)
		orderRouterPrivate.GET("/get-order-by-manager", controllers.Order.GetOrdersByManager)
		orderRouterPrivate.GET("/get-order-by-user", controllers.Order.GetOrdersByUser)
	}
}
