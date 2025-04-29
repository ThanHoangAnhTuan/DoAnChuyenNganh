package impl

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
)

type AccommodationDetailImpl struct {
	sqlc *database.Queries
}

func (a *AccommodationDetailImpl) CreateAccommodationDetail(ctx *gin.Context, in *vo.CreateAccommodationDetailInput) (codeStatus int, out *vo.CreateAccommodationDetailOutput, err error) {
	panic("unimplemented")
}

func (a *AccommodationDetailImpl) DeleteAccommodationDetail(ctx context.Context, in *vo.DeleteAccommodationDetailInput) (codeResult int, err error) {
	panic("unimplemented")
}

func (a *AccommodationDetailImpl) GetAccommodationDetails(ctx context.Context) (codeStatus int, out []*vo.GetAccommodationDetails, err error) {
	panic("unimplemented")
}

func (a *AccommodationDetailImpl) GetAccommodationDetailsByManager(ctx context.Context) (codeStatus int, out []*vo.GetAccommodationDetails, err error) {
	panic("unimplemented")
}

func (a *AccommodationDetailImpl) UpdateAccommodationDetail(ctx *gin.Context, in *vo.UpdateAccommodationDetailInput) (codeResult int, out *vo.UpdateAccommodationDetailOutput, err error) {
	panic("unimplemented")
}

func NewAccommodationDetailImpl(sqlc *database.Queries) *AccommodationDetailImpl {
	return &AccommodationDetailImpl{
		sqlc: sqlc,
	}
}
