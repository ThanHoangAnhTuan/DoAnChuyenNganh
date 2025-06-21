package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type ManagerGroup struct {
}

func (g *ManagerGroup) InitManagerGroup(Router *gin.RouterGroup) {
	managerRouterPublic := Router.Group("/manager")
	{

		managerRouterPublic.POST("/login", controllers.ManagerLogin.Login)
	}

	managerRouterPrivate := Router.Group("/manager")
	managerRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		managerRouterPublic.POST("/register", controllers.ManagerLogin.Register)
		managerRouterPrivate.GET("/accommodations", controllers.Accommodation.GetAccommodationsByManager)
	}
}
