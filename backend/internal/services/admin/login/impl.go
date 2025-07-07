package login

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/consts"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/database"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/internal/vo"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/response"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/auth"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/crypto"
	utiltime "github.com/thanhoanganhtuan/DoAnChuyenNganh/pkg/utils/util_time"
	"go.uber.org/zap"
)

type serviceImpl struct {
	sqlc *database.Queries
}

func New(sqlc *database.Queries) Service {
	return &serviceImpl{sqlc: sqlc}
}

func (m *serviceImpl) Login(ctx *gin.Context, in *vo.AdminLoginInput) (codeStatus int, out *vo.AdminLoginOutput, err error) {
	out = &vo.AdminLoginOutput{}

	// TODO: get admin info
	userAdmin, err := m.sqlc.GetUserAdmin(ctx, in.UserAccount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response.ErrCodeLoginFailed, nil, fmt.Errorf("user admin not found")
		}
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("get user admin failed: %s", err)
	}

	// TODO: check password match
	if !crypto.CheckPasswordHash(in.UserPassword, userAdmin.Password) {
		return response.ErrCodeLoginFailed, nil, fmt.Errorf("dose not match password")
	}

	// TODO: check two-factor authentication

	// TODO: update login
	go func() {
		err := m.sqlc.UpdateUserAdminLogin(ctx, database.UpdateUserAdminLoginParams{
			LoginTime: utiltime.GetTimeNow(),
			Account:   in.UserAccount,
		})
		if err != nil {
			fmt.Printf("failed to update user login time: %v", err.Error())
			global.Logger.Error("failed to update user login time: ", zap.String("error", err.Error()))
		}
	}()

	// TODO: create uuid user
	subToken := utils.GenerateCliTokenUUID(userAdmin.ID)

	userAdminInfor := vo.AdminInfor{
		Account:  userAdmin.Account,
		UserName: userAdmin.UserName,
	}

	userAdminInforJson, err := json.Marshal(userAdminInfor)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("convert admin info to json failed: %v", err)
	}

	// TODO: save admin info to redis
	err = global.Redis.SetEx(ctx, subToken, userAdminInforJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("save admin info to redis failed: %s", err)
	}

	// TODO: create jwt token
	out.Token, err = auth.CreateToken(userAdmin.ID, consts.ADMIN)
	if err != nil {
		return response.ErrCodeInternalServerError, nil, fmt.Errorf("create token failed: %s", err)
	}

	// TODO: response
	out.Account = userAdminInfor.Account
	out.UserName = userAdminInfor.UserName
	return response.ErrCodeLoginSuccess, out, nil
}

func (m *serviceImpl) Register(ctx *gin.Context, in *vo.AdminRegisterInput) (codeStatus int, err error) {
	// TODO: check email exists in user admin
	adminFound, err := m.sqlc.CheckUserAdminExistsByEmail(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("check admin already exists failed: %s", err)
	}

	if adminFound {
		return response.ErrCodeAccountAlreadyExists, fmt.Errorf("admin already exists")
	}

	// TODO: check user spam / rate limiting by ip

	// TODO: create admin
	id := uuid.New().String()
	now := utiltime.GetTimeNow()
	hashPassword, err := crypto.HashPassword(in.UserPassword)
	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("hash password admin failed: %s", err)
	}

	err = m.sqlc.CreateUserAdmin(ctx, database.CreateUserAdminParams{
		ID:        id,
		Account:   in.UserAccount,
		Password:  hashPassword,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return response.ErrCodeInternalServerError, fmt.Errorf("register admin failed: %s", err)
	}

	return response.ErrCodeRegisterSuccess, nil
}
