package review

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/controllers"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/middlewares"
)

type ReviewRouter struct {
}

func (r ReviewRouter) InitReviewRouter(Router *gin.RouterGroup) {
	reviewRouterPublic := Router.Group("/review")
	{
		reviewRouterPublic.GET("/", controllers.Review.GetReviews)
	}

	reviewRouterPrivate := Router.Group("/review")
	reviewRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		reviewRouterPrivate.POST("/", controllers.Review.CreateReview)
	}
}
