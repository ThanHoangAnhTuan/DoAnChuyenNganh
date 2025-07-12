package initializer

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.ValidatorMiddleware())
	r.Use(middlewares.TimezoneMiddleware())

	// Add tracing middleware
	r.Use(middlewares.TracingMiddleware())
	r.Use(middlewares.AddSpanToContext())

	r.Static("/uploads/", "./storage/uploads")

	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User
	managerRouter := routers.RouterGroupApp.Manager
	accommodationRouter := routers.RouterGroupApp.Accommodation
	accommodationDetailRouter := routers.RouterGroupApp.AccommodationDetail
	uploadRouter := routers.RouterGroupApp.Upload
	orderRouter := routers.RouterGroupApp.Order
	facilityRouter := routers.RouterGroupApp.Facility
	facilityDetailRouter := routers.RouterGroupApp.FacilityDetail
	paymentRouter := routers.RouterGroupApp.Payment
	reviewRouter := routers.RouterGroupApp.Review
	statsRouter := routers.RouterGroupApp.Stats
	roomRouter := routers.RouterGroupApp.Room

	MainGroup := r.Group("api/v1")
	{
		userRouter.InitUserRouter(MainGroup)
	}
	{
		adminRouter.InitAdminRouter(MainGroup)
	}
	{
		managerRouter.InitManagerGroup(MainGroup)
	}
	{
		accommodationRouter.InitAccommodationRouter(MainGroup)
	}
	{
		accommodationDetailRouter.InitAccommodationDetailRouter(MainGroup)
	}
	{
		uploadRouter.InitUploadRouter(MainGroup)
	}
	{
		orderRouter.InitOrderRouter(MainGroup)
	}
	{
		facilityRouter.InitFacilityRouter(MainGroup)
	}
	{
		paymentRouter.InitPaymentRouter(MainGroup)
	}
	{
		facilityDetailRouter.InitFacilityDetailRouter(MainGroup)
	}
	{
		reviewRouter.InitReviewRouter(MainGroup)
	}
	{
		statsRouter.InitStatsRouter(MainGroup)
	}
	{
		roomRouter.InitAccommodationRoomRouter(MainGroup)
	}
	return r
}
