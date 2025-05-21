package services

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
)

type (
	IFacility interface {
		CreateFacility(ctx *gin.Context, in *vo.CreateFacilityInput) (codeStatus int, out *vo.CreateFacilityOutput, err error)
		GetFacilities(ctx context.Context) (codeStatus int, out []*vo.GetFacilitiesOutput, err error)
	}
)

var (
	localFacility IFacility
)

func Facility() IFacility {
	if localFacility == nil {
		panic("Implement localFacility not found for interface IFacility")
	}
	return localFacility
}

func InitFacility(i IFacility) {
	localFacility = i
}
