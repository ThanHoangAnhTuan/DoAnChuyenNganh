package accommodation

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/controllers"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/middlewares"
)

type AccommodationRouter struct {
}

func (ur *AccommodationRouter) InitAccommodationRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/accommodation")
	{
		userRouterPublic.GET("/get-accommodations", controllers.Accommodation.GetAccommodations)
	}

	userRouterPrivate := Router.Group("/accommodation")
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("/get-accommodations-by-manager-id", controllers.Accommodation.GetAccommodationsByManager)
		userRouterPrivate.POST("create-accommodation", controllers.Accommodation.CreateAccommodation)
		userRouterPrivate.PUT("update-accommodation", controllers.Accommodation.UpdateAccommodation)
		userRouterPrivate.DELETE("delete-accommodation/:id", controllers.Accommodation.DeleteAccommodation)
	}
}
