package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IFacility interface {
		CreateFacility(ctx *gin.Context, in *vo.CreateFacilityInput) (codeStatus int, out *vo.CreateFacilityOutput, err error)
		GetFacilities(ctx *gin.Context) (codeStatus int, out []*vo.GetFacilitiesOutput, err error)
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
