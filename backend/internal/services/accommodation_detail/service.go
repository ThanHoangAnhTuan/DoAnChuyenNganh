package accommodationdetail

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	GetAccommodationDetails(ctx *gin.Context, in *vo.GetAccommodationDetailsInput) (codeStatus int, out []*vo.GetAccommodationDetailsOutput, err error)
	GetAccommodationDetailsByManager(ctx *gin.Context, in *vo.GetAccommodationDetailsByManagerInput) (codeStatus int, out []*vo.GetAccommodationDetailsByManagerOutput, err error)
	CreateAccommodationDetail(ctx *gin.Context, in *vo.CreateAccommodationDetailInput) (codeStatus int, out *vo.CreateAccommodationDetailOutput, err error)
	UpdateAccommodationDetail(ctx *gin.Context, in *vo.UpdateAccommodationDetailInput) (codeResult int, out *vo.UpdateAccommodationDetailOutput, err error)
	DeleteAccommodationDetail(ctx *gin.Context, in *vo.DeleteAccommodationDetailInput) (codeResult int, err error)
}
