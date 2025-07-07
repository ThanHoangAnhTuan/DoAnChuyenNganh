package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/register", controllers.AdminLogin.Register)
		adminRouterPublic.POST("/login", controllers.AdminLogin.Login)
	}

	adminRouterPrivate := Router.Group("/admin")
	adminRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		adminRouterPrivate.GET("/managers", controllers.AdminManager.GetManagers)
		adminRouterPrivate.GET("/manager/:id/accommodations", controllers.AdminManager.GetAccommodationsOfManager)
		adminRouterPrivate.PUT("/verify-accommodation", controllers.AdminManager.VerifyAccommodation)
		adminRouterPrivate.PUT("/set-deleted-accommodation", controllers.AdminManager.SetDeletedAccommodation)
	}
}
