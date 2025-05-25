package payment

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type PaymentRouter struct {
}

func (r PaymentRouter) InitPaymentRouter(Router *gin.RouterGroup) {
	paymentRouterPublic := Router.Group("/payment")
	{
		paymentRouterPublic.GET("/vnpay_return")
		paymentRouterPublic.GET("/vnpay_ipn")
	}

	paymentRouterPrivate := Router.Group("/payment")
	paymentRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		paymentRouterPrivate.GET("/")
		paymentRouterPrivate.GET("/create_payment_url")
		paymentRouterPrivate.POST("/create_payment_url")
		paymentRouterPrivate.GET("/querydr")
		paymentRouterPrivate.POST("/querydr")
		paymentRouterPrivate.GET("/refund")
		paymentRouterPrivate.POST("/refund")
	}

}
