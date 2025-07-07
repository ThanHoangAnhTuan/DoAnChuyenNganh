package accommodationroom

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type AccommodationRoomRouter struct {
}

func (ur *AccommodationRoomRouter) InitAccommodationRoomRouter(Router *gin.RouterGroup) {
	userRouterPrivate := Router.Group("/accommodation-room")
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("/:accommodation_type_id", controllers.AccommodationRoom.GetAccommodationRooms)
		userRouterPrivate.POST("", controllers.AccommodationRoom.CreateAccommodationRoom)
		userRouterPrivate.PUT("", controllers.AccommodationRoom.UpdateAccommodationRoom)
		userRouterPrivate.DELETE("/:id", controllers.AccommodationRoom.DeleteAccommodationRoom)
	}
}
