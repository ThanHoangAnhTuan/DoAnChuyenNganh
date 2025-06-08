package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IManagerLogin interface {
		Register(ctx *gin.Context, in *vo.ManagerRegisterInput) (codeStatus int, err error)
		Login(ctx *gin.Context, in *vo.ManagerLoginInput) (codeStatus int, out *vo.ManagerLoginOutput, err error)
	}
)

var (
	localManagerLogin IManagerLogin
)

func ManagerLogin() IManagerLogin {
	if localManagerLogin == nil {
		panic("Implement localManagerLogin not found for interface IManagerLogin")
	}
	return localManagerLogin
}

func InitManagerLogin(i IManagerLogin) {
	localManagerLogin = i
}
