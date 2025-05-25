package services

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IAccommodation interface {
		GetAccommodations(ctx context.Context) (codeStatus int, out []*vo.GetAccommodations, err error)
		GetAccommodationByCity(ctx context.Context, in *vo.GetAccommodationByCityInput) (codeStatus int, out []*vo.GetAccommodationsByCity, err error)
		GetAccommodationById(ctx context.Context, in *vo.GetAccommodationByIdInput) (codeStatus int, out *vo.GetAccommodationByIdOutput, err error)

		CreateAccommodation(ctx *gin.Context, in *vo.CreateAccommodationInput) (codeStatus int, out *vo.CreateAccommodationOutput, err error)
		UpdateAccommodation(ctx *gin.Context, in *vo.UpdateAccommodationInput) (codeResult int, out *vo.UpdateAccommodationOutput, err error)
		DeleteAccommodation(ctx context.Context, in *vo.DeleteAccommodationInput) (codeResult int, err error)
		GetAccommodationsByManager(ctx context.Context) (codeStatus int, out []*vo.GetAccommodations, err error)
	}
)

var (
	localAccommodation IAccommodation
)

func Accommodation() IAccommodation {
	if localAccommodation == nil {
		panic("Implement localAccommodation not found for interface IAccommodation")
	}
	return localAccommodation
}

func InitAccommodation(i IAccommodation) {
	localAccommodation = i
}
