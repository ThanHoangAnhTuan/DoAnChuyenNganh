package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	GetManagers(ctx *gin.Context, in *vo.GetManagerInput) (codeStatus int, out []*vo.GetManagerOutput, pagination *vo.BasePaginationOutput, err error)
}
