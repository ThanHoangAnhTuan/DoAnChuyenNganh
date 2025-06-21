package login

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	Register(ctx *gin.Context, in *vo.RegisterInput) (codeStatus int, err error)
	VerifyOTP(ctx *gin.Context, in *vo.VerifyOTPInput) (codeStatus int, out *vo.VerifyOTPOutput, err error)
	UpdatePasswordRegister(ctx *gin.Context, in *vo.UpdatePasswordRegisterInput) (codeStatus int, err error)
	Login(ctx *gin.Context, in *vo.LoginInput) (codeStatus int, out *vo.LoginOutput, err error)
}
