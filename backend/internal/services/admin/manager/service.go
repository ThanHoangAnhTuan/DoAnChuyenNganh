package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	GetManagers(ctx *gin.Context, in *vo.GetManagerInput) (codeStatus int, out []*vo.GetManagerOutput, pagination *vo.BasePaginationOutput, err error)
	GetAccommodationsOfManager(ctx *gin.Context, in *vo.GetAccommodationsOfManagerInput) (codeStatus int, out []*vo.GetAccommodationsOfManagerOutput, pagination *vo.BasePaginationOutput, err error)
	VerifyAccommodation(ctx *gin.Context, in *vo.VerifyAccommodationInput) (codeStatus int, err error)
	SetDeletedAccommodation(ctx *gin.Context, in *vo.SetDeletedAccommodationInput) (codeResult int, err error)
}
