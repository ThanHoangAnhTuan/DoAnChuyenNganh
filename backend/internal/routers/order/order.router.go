package order

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type OrderRouter struct {
}

func (r *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	orderRouterPrivate := Router.Group("/order")
	orderRouterPrivate.Use(middlewares.AuthMiddleware())
}
