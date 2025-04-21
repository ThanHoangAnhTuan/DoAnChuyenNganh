package initializer

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
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
	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User
	managerRouter := routers.RouterGroupApp.Manager

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
	return r
}
