package review

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	CreateReview(ctx *gin.Context, in *vo.CreateReviewInput) (codeStatus int, out *vo.CreateReviewOutput, err error)
	GetReviews(ctx *gin.Context, in *vo.GetReviewsInput) (codeStatus int, out []*vo.GetReviewOutput, pagination *vo.BasePaginationOutput, err error)
}
