package info

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type Service interface {
	GetUserInfo(ctx *gin.Context) (codeStatus int, out *vo.GetUserInfoOutput, err error)
	UpdateUserInfo(ctx *gin.Context, in *vo.UpdateUserInfoInput) (codeStatus int, out *vo.UpdateUserInfoOutput, err error)
	UploadUserAvatar(ctx *gin.Context, in *vo.UploadUserAvatarInput) (codeStatus int, avatarPath string, err error)
}
