package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type IReview interface {
	CreateReview(ctx *gin.Context, in *vo.CreateReviewInput) (codeStatus int, out *vo.CreateReviewOutput, err error)
	GetReviews(ctx *gin.Context, in *vo.GetReviewsInput) (codeStatus int, out *vo.GetReviewOutput, err error)
}

var localReview IReview

func Review() IReview {
	if localReview == nil {
		panic("Implement localReview not found for interface IReview")
	}
	return localReview
}

func InitReview(i IReview) {
	localReview = i
}
