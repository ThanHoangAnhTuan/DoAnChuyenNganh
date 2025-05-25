package order

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/controllers"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/middlewares"
)

type OrderRouter struct {
}

func (r *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderRouterPrivate := Router.Group("/order")
	orderRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		orderRouterPrivate.POST("/create-order", controllers.Order.CreateOrder)
		// orderRouterPrivate.GET("/get-images/:id", controllers.UploadImage.GetImages)
	}
}
