package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/controllers"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/middlewares"
)

type UploadRouter struct {
}

func (ir *UploadRouter) InitUploadRouter(Router *gin.RouterGroup) {
	uploadRouterPrivate := Router.Group("/image")
	uploadRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		uploadRouterPrivate.POST("/upload-images", controllers.UploadImage.UploadImages)
		uploadRouterPrivate.GET("/get-images/:id", controllers.UploadImage.GetImages)
	}
}
