package accommodation

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type AccommodationRouter struct {
}

func (ur *AccommodationRouter) InitAccommodationRouter(Router *gin.RouterGroup) {
	accommodationRouterPublic := Router.Group("/accommodations")
	{
		accommodationRouterPublic.GET("", controllers.Accommodation.GetAccommodations)
		accommodationRouterPublic.GET("/:id", controllers.Accommodation.GetAccommodation)
	}

	accommodationRouterPrivate := Router.Group("/accommodations")
	accommodationRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		accommodationRouterPrivate.POST("", controllers.Accommodation.CreateAccommodation)
		accommodationRouterPrivate.PUT("", controllers.Accommodation.UpdateAccommodation)
		accommodationRouterPrivate.DELETE("", controllers.Accommodation.DeleteAccommodation)
	}
}
