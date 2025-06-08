package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IAdminLogin interface {
		Register(ctx *gin.Context, in *vo.AdminRegisterInput) (codeStatus int, err error)
		Login(ctx *gin.Context, in *vo.AdminLoginInput) (codeStatus int, out *vo.AdminLoginOutput, err error)
	}
)

var (
	localAdminLogin IAdminLogin
)

func AdminLogin() IAdminLogin {
	if localAdminLogin == nil {
		panic("Implement localAdminLogin not found for interface IAdminLogin")
	}
	return localAdminLogin
}

func InitAdminLogin(i IAdminLogin) {
	localAdminLogin = i
}
