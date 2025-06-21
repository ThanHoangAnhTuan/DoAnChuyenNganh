package login

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	Register(ctx *gin.Context, in *vo.AdminRegisterInput) (codeStatus int, err error)
	Login(ctx *gin.Context, in *vo.AdminLoginInput) (codeStatus int, out *vo.AdminLoginOutput, err error)
}
