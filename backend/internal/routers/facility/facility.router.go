package facility

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/controllers"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/middlewares"
)

type FacilityRouter struct {
}

func (ar *FacilityRouter) InitFacilityRouter(Router *gin.RouterGroup) {
	facilityRouterPublic := Router.Group("/facility")
	facilityRouterPublic.Use(middlewares.AuthMiddleware())
	{
		facilityRouterPublic.POST("/create-facility", controllers.Facility.CreateFacility)
		facilityRouterPublic.GET("/get-facilities", controllers.Facility.GetFacilities)
	}
}
