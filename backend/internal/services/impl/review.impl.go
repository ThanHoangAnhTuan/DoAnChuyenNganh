package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type ReviewImpl struct {
	sqlc *database.Queries
}

func (r *ReviewImpl) CreateReview(ctx *gin.Context, in *vo.CreateReviewInput) (codeStatus int, out *vo.CreateReviewOutput, err error) {
	// TODO: check user exists

	// TODO: check accommodation exists

	// TODO: Check if the user has booked a room before

	// TODO: Create review
	panic("unimplemented")
}

func (r *ReviewImpl) GetReviews(ctx *gin.Context, in *vo.GetReviewsInput) (codeStatus int, out *vo.GetReviewOutput, err error) {
	// TODO: check accommodation exists

	// TODO: get reviews
	panic("unimplemented")
}

func NewReviewImpl(sqlc *database.Queries) *ReviewImpl {
	return &ReviewImpl{
		sqlc: sqlc,
	}
}
