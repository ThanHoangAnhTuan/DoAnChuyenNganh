package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
)

type UserInfoImpl struct {
	sqlc *database.Queries
}

func (u *UserInfoImpl) GetUserInfo(ctx *gin.Context) (codeStatus int, out *vo.GetUserInfoOutput, err error) {
	panic("unimplemented")
}

func (u *UserInfoImpl) UpdateUserInfo(ctx *gin.Context, in *vo.UpdateUserInfoInput) (codeStatus int, out *vo.UpdateUserInfoOutput, err error) {
	panic("unimplemented")
}

func NewUserInfoImpl(sqlc *database.Queries) *UserInfoImpl {
	return &UserInfoImpl{
		sqlc: sqlc,
	}
}
