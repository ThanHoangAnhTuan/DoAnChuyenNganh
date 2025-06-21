package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", controllers.UserLogin.Register)
		userRouterPublic.POST("/verify-otp", controllers.UserLogin.VerifyOTP)
		userRouterPublic.POST("/login", controllers.UserLogin.Login)
		userRouterPublic.POST("/update-password-register", controllers.UserLogin.UpdatePasswordRegister)
	}

	userRouterPrivate := Router.Group("/user")
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("/get-user-info", controllers.UserInfo.GetUserInfo)
		userRouterPrivate.POST("/update-user-info", controllers.UserInfo.UpdateUserInfo)
		userRouterPrivate.POST("/upload-user-avatar", controllers.UserInfo.UploadUserAvatar)
	}
}
