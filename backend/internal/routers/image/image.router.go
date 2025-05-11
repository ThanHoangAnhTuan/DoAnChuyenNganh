package image

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/controllers"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/middlewares"
)

type ImageRouter struct {
}

func (ir *ImageRouter) InitImageRouter(Router *gin.RouterGroup) {
	imageRouterPrivate := Router.Group("/images")
	imageRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		imageRouterPrivate.POST("/upload-images", controllers.UploadImages.UploadImages)
		imageRouterPrivate.GET("/get-images/:id", controllers.UploadImages.GetImages)
	}
}
