package accommodationdetail

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/controllers"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/middlewares"
)

type AccommodationDetailRouter struct {
}

func (ur *AccommodationDetailRouter) InitAccommodationDetailRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/accommodation-detail")
	{
		userRouterPublic.GET("/get-accommodation-details/:id", controllers.AccommodationDetail.GetAccommodationDetails)
	}

	userRouterPrivate := Router.Group("/accommodation-detail")
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.POST("/create-accommodation-detail", controllers.AccommodationDetail.CreateAccommodationDetail)
		userRouterPrivate.PUT("/update-accommodation-detail", controllers.AccommodationDetail.UpdateAccommodationDetail)
		userRouterPrivate.DELETE("/delete-accommodation-detail/:id", controllers.AccommodationDetail.DeleteAccommodationDetail)
	}
}
