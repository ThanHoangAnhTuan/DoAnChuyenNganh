package services

import (
	"context"

	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type (
	IManagerLogin interface {
		Register(ctx context.Context, in *vo.ManagerRegisterInput) (codeStatus int, err error)
		Login(ctx context.Context, in *vo.ManagerLoginInput) (codeStatus int, out *vo.ManagerLoginOutput, err error)
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
