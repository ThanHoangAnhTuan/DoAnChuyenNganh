package accommodationdetail

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/controllers"
)

type AccommodationDetailRouter struct {
}

func (ur *AccommodationDetailRouter) InitAccommodationDetailRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/accommodation-detail")
	{
		userRouterPublic.GET("/get-accommodation-details", controllers.AccommodationDetail.GetAccommodationDetails)
	}

	userRouterPrivate := Router.Group("/accommodation-detail")
	// userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPublic.GET("/get-accommodation-details-by-manager-id/", controllers.AccommodationDetail.GetAccommodationDetailsByManager)
		userRouterPrivate.POST("create-accommodation-detail", controllers.AccommodationDetail.CreateAccommodationDetail)
		userRouterPrivate.PUT("update-accommodation-detail", controllers.AccommodationDetail.UpdateAccommodationDetail)
		userRouterPrivate.DELETE("delete-accommodation-detail/:id", controllers.AccommodationDetail.DeleteAccommodationDetail)
	}
}
