package initializer

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/middlewares"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/routers"
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

	r.Static("/uploads/", "./storage/uploads")

	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User
	managerRouter := routers.RouterGroupApp.Manager
	accommodationRouter := routers.RouterGroupApp.Accommodation
	accommodationDetailRouter := routers.RouterGroupApp.AccommodationDetail
	imageRouter := routers.RouterGroupApp.Image

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
		imageRouter.InitImageRouter(MainGroup)
	}
	return r
}
