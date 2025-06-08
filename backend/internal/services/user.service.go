package services

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

// TODO: Refactor interface
type (
	IUserLogin interface {
		Register(ctx *gin.Context, in *vo.RegisterInput) (codeStatus int, err error)
		VerifyOTP(ctx *gin.Context, in *vo.VerifyOTPInput) (codeStatus int, out *vo.VerifyOTPOutput, err error)
		UpdatePasswordRegister(ctx *gin.Context, in *vo.UpdatePasswordRegisterInput) (codeStatus int, err error)
		Login(ctx *gin.Context, in *vo.LoginInput) (codeStatus int, out *vo.LoginOutput, err error)
	}

	IUserInfo interface {
		GetUserById(ctx *gin.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx *gin.Context) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserLogin IUserLogin
	localUserInfo  IUserInfo
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("Implement localUserAdmin not found for interface IUserAdmin")
	}
	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("Implement localUserLogin not found for interface IUserLogin")
	}
	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("Implement localUserInfo not found for interface IUserInfo")
	}
	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}
