package facilitydetail

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	CreateFacilityDetail(ctx *gin.Context, in *vo.CreateFacilityDetailInput) (codeStatus int, out *vo.CreateFacilityDetailOutput, err error)
	UpdateFacilityDetail(ctx *gin.Context, in *vo.UpdateFacilityDetailInput) (codeStatus int, out *vo.UpdateFacilityDetailOutput, err error)
	DeleteFacilityDetail(ctx *gin.Context, in *vo.DeleteFacilityDetailInput) (codeStatus int, err error)
	GetFacilityDetail(ctx *gin.Context) (codeStatus int, out []*vo.GetFacilityDetailOutput, err error)
}
