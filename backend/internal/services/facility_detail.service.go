package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IFacilityDetail interface {
		CreateFacilityDetail(ctx *gin.Context, in *vo.CreateFacilityDetailInput) (codeStatus int, out *vo.CreateFacilityDetailOutput, err error)
		GetFacilityDetail(ctx *gin.Context) (codeStatus int, out []*vo.GetFacilityDetailOutput, err error)
	}
)

var (
	localFacilityDetail IFacilityDetail
)

func FacilityDetail() IFacilityDetail {
	if localFacilityDetail == nil {
		panic("Implement localFacilityDetail not found for interface IFacilityDetail")
	}
	return localFacilityDetail
}

func InitFacilityDetail(i IFacilityDetail) {
	localFacilityDetail = i
}
