package impl

import (
	"database/sql"
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
)

type UserInfoImpl struct {
	sqlc *database.Queries
}

func (u *UserInfoImpl) UploadUserAvatar(ctx *gin.Context, in *vo.UploadUserAvatarInput) (codeStatus int, avatarPath string, err error) {
	// TODO: get user id in context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, "", fmt.Errorf("userID not found in context")
	}

	// TODO: get user info from db
	user, err := u.sqlc.GetUserInfoByID(ctx, userID)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, "", fmt.Errorf("get user info failed: %s", err)
	}

	err = deleteImageToDisk([]string{user.Image})
	if err != nil {
		return response.ErrCodeDeleteAccommodationDetailImagesFailed, "", fmt.Errorf("delete images in disk of accommodation detail failed: %s", err)
	}

	// TODO: Save image to disk
	codeStatus, imagesFileName, err := saveImageToDisk(ctx, []*multipart.FileHeader{in.Avatar})
	if err != nil {
		return codeStatus, "", err
	}

	// TODO: Save image to db
	err = u.sqlc.UpdateUserAvatar(ctx, database.UpdateUserAvatarParams{
		Image: imagesFileName[0],
		ID:    userID,
	})

	if err != nil {
		return response.ErrCodeUpdateUserInfoFailed, "", fmt.Errorf("update avatar failed: %s", err)
	}

	return response.ErrCodeUpdateUserInfoSuccess, imagesFileName[0], nil
}

func (u *UserInfoImpl) GetUserInfo(ctx *gin.Context) (codeStatus int, out *vo.GetUserInfoOutput, err error) {
	out = &vo.GetUserInfoOutput{}

	// TODO: get user id in context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	// TODO: get user info from db
	user, err := u.sqlc.GetUserInfoByID(ctx, userID)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, nil, fmt.Errorf("get user info failed: %s", err)
	}

	out.Account = user.Account
	out.Birthday = user.Birthday
	out.Gender = map[uint8]string{0: "male", 1: "female"}[user.Gender]
	out.ID = user.ID
	out.Image = user.Image
	out.Phone = user.Phone.String
	out.Username = user.UserName

	return response.ErrCodeGetUserInfoSuccess, out, nil
}

func (u *UserInfoImpl) UpdateUserInfo(ctx *gin.Context, in *vo.UpdateUserInfoInput) (codeStatus int, out *vo.UpdateUserInfoOutput, err error) {
	out = &vo.UpdateUserInfoOutput{}

	// TODO: get user id in context
	userID, ok := utils.GetUserIDFromGin(ctx)
	if !ok {
		return response.ErrCodeUnauthorized, nil, fmt.Errorf("userID not found in context")
	}

	isExists, err := u.sqlc.CheckUserInfoExists(ctx, userID)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, nil, fmt.Errorf("get user info failed: %s", err)
	}

	if !isExists {
		return response.ErrCodeGetUserInfoNotFound, nil, fmt.Errorf("user info not found")
	}

	// TODO: update user info
	now := utiltime.GetTimeNow()
	err = u.sqlc.UpdateUserInfo(ctx, database.UpdateUserInfoParams{
		UserName: in.Username,
		Phone: sql.NullString{
			String: in.Phone,
			Valid:  true,
		},
		Gender:   in.Gender,
		Birthday: in.Birthday,
		// Email: sql.NullString{
		// 	String: in.Email,
		// 	Valid:  true,
		// },
		UpdatedAt: now,
		ID:        userID,
	})

	if err != nil {
		return response.ErrCodeUpdateUserInfoSuccess, nil, fmt.Errorf("update user info failed: %s", err)
	}

	// TODO: get user info
	user, err := u.sqlc.GetUserInfoByID(ctx, userID)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, nil, fmt.Errorf("get user info failed: %s", err)
	}

	out.Account = user.Account
	out.Birthday = user.Birthday
	out.Gender = map[uint8]string{0: "male", 1: "female"}[user.Gender]
	out.ID = user.ID
	out.Phone = user.Phone.String
	out.Username = user.UserName

	return response.ErrCodeUpdateUserInfoSuccess, out, nil
}

func NewUserInfoImpl(sqlc *database.Queries) *UserInfoImpl {
	return &UserInfoImpl{
		sqlc: sqlc,
	}
}
