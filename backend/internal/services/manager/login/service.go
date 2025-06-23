package login

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	Register(ctx *gin.Context, in *vo.ManagerRegisterInput) (codeStatus int, err error)
	Login(ctx *gin.Context, in *vo.ManagerLoginInput) (codeStatus int, out *vo.ManagerLoginOutput, err error)
}
