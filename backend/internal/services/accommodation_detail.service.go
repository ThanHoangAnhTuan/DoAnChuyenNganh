package services

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
)

type (
	IAccommodationDetail interface {
		GetAccommodationDetails(ctx context.Context, in *vo.GetAccommodationDetailsInput) (codeStatus int, out []*vo.GetAccommodationDetailsOutput, err error)
		CreateAccommodationDetail(ctx *gin.Context, in *vo.CreateAccommodationDetailInput) (codeStatus int, out *vo.CreateAccommodationDetailOutput, err error)
		UpdateAccommodationDetail(ctx *gin.Context, in *vo.UpdateAccommodationDetailInput) (codeResult int, out *vo.UpdateAccommodationDetailOutput, err error)
		DeleteAccommodationDetail(ctx context.Context, in *vo.DeleteAccommodationDetailInput) (codeResult int, err error)
	}
)

var (
	localAccommodationDetail IAccommodationDetail
)

func AccommodationDetail() IAccommodationDetail {
	if localAccommodationDetail == nil {
		panic("Implement localAccommodationDetail not found for interface IAccommodationDetail")
	}
	return localAccommodationDetail
}

func InitAccommodationDetail(i IAccommodationDetail) {
	localAccommodationDetail = i
}
