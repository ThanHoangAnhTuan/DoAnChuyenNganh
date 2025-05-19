package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/controllers"
)

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/register", controllers.AdminLogin.Register)
		adminRouterPublic.POST("/login", controllers.AdminLogin.Login)
	}
}
