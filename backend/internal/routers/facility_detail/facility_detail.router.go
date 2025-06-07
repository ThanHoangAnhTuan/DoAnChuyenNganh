package facilitydetail

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type FacilityDetailRouter struct {
}

func (ar *FacilityDetailRouter) InitFacilityDetailRouter(Router *gin.RouterGroup) {
	facilityDetailRouterPublic := Router.Group("/facility-detail")
	{
		facilityDetailRouterPublic.GET("/get-facility-detail", controllers.FacilityDetail.GetFacilityDetail)
	}

	facilityDetailRouterPrivate := Router.Group("/facility-detail")
	facilityDetailRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		facilityDetailRouterPrivate.POST("/create-facility-detail", controllers.FacilityDetail.CreateFacilityDetail)
	}
}
