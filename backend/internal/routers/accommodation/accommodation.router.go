package accommodation

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type AccommodationRouter struct {
}

func (ur *AccommodationRouter) InitAccommodationRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/accommodation")
	{
		userRouterPublic.GET("/get-accommodations", controllers.Accommodation.GetAccommodations)
		userRouterPublic.GET("/get-accommodations-by-city/:city", controllers.Accommodation.GetAccommodationByCity)
		userRouterPublic.GET("/get-accommodation-by-id/:id", controllers.Accommodation.GetAccommodationById)
	}

	userRouterPrivate := Router.Group("/accommodation")
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("/get-accommodations-by-manager", controllers.Accommodation.GetAccommodationsByManager)
		userRouterPrivate.POST("/create-accommodation", controllers.Accommodation.CreateAccommodation)
		userRouterPrivate.PUT("/update-accommodation", controllers.Accommodation.UpdateAccommodation)
		userRouterPrivate.DELETE("/delete-accommodation", controllers.Accommodation.DeleteAccommodation)
	}
}
