package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IAccommodation interface {
		GetAccommodations(ctx *gin.Context, in *vo.GetAccommodationsInput) (codeStatus int, out []*vo.GetAccommodationsOutput, pagination *vo.BasePaginationOutput, err error)
		GetAccommodationsByManager(ctx *gin.Context, in *vo.GetAccommodationsInput) (codeStatus int, out []*vo.GetAccommodationsOutput, pagination *vo.BasePaginationOutput, err error)
		GetAccommodation(ctx *gin.Context, in *vo.GetAccommodationInput) (codeStatus int, out *vo.GetAccommodationOutput, err error)
		CreateAccommodation(ctx *gin.Context, in *vo.CreateAccommodationInput) (codeStatus int, out *vo.CreateAccommodationOutput, err error)
		UpdateAccommodation(ctx *gin.Context, in *vo.UpdateAccommodationInput) (codeResult int, out *vo.UpdateAccommodationOutput, err error)
		DeleteAccommodation(ctx *gin.Context, in *vo.DeleteAccommodationInput) (codeResult int, err error)
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
