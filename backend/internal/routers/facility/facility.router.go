package facility

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type FacilityRouter struct {
}

func (ar *FacilityRouter) InitFacilityRouter(Router *gin.RouterGroup) {
	facilityRouterPublic := Router.Group("/facility")
	{
		facilityRouterPublic.GET("/get-facilities", controllers.Facility.GetFacilities)
	}

	facilityRouterPrivate := Router.Group("/facility")
	facilityRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		facilityRouterPrivate.POST("/create-facility", controllers.Facility.CreateFacility)
	}
}
