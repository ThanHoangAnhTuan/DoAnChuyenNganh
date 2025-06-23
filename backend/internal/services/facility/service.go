package facility

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	CreateFacility(ctx *gin.Context, in *vo.CreateFacilityInput) (codeStatus int, out *vo.CreateFacilityOutput, err error)
	UpdateFacility(ctx *gin.Context, in *vo.UpdateFacilityInput) (codeStatus int, out *vo.UpdateFacilityOutput, err error)
	DeleteFacility(ctx *gin.Context, in *vo.DeleteFacilityInput) (codeStatus int, err error)
	GetFacilities(ctx *gin.Context) (codeStatus int, out []*vo.GetFacilitiesOutput, err error)
}
