package impl

import (
	"encoding/json"
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
)

type ManagerLoginImpl struct {
	sqlc *database.Queries
}

func (m *ManagerLoginImpl) Login(ctx *gin.Context, in *vo.ManagerLoginInput) (codeStatus int, out *vo.ManagerLoginOutput, err error) {
	out = &vo.ManagerLoginOutput{}

	// TODO: get manager info
	userManager, err := m.sqlc.GetUserManager(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, nil, fmt.Errorf("get user info failed: %s", err)
	}

	// TODO: check password match
	if !crypto.CheckPasswordHash(in.UserPassword, userManager.Password) {
		return response.ErrCodePasswordNotMatch, nil, fmt.Errorf("dose not match password")
	}

	// TODO: check two-factor authentication

	// TODO: update login
	go m.sqlc.UpdateUserManagerLogin(ctx, database.UpdateUserManagerLoginParams{
		LoginTime: utiltime.GetTimeNow(),
		Account:   in.UserAccount,
	})

	// TODO: create uuid user
	subToken := utils.GenerateCliTokenUUID(userManager.ID)

	userManagerInfor := vo.ManagerInfor{
		Account:  userManager.Account,
		UserName: userManager.UserName,
	}

	userManagerInforJson, err := json.Marshal(userManagerInfor)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("convert to json failed: %v", err)
	}

	// TODO: save manager info to redis
	err = global.Redis.SetEx(ctx, subToken, userManagerInforJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeSaveDataFailed, nil, fmt.Errorf("save manager info to redis failed: %s", err)
	}

	out.Token, err = auth.CreateToken(userManager.ID, global.Manager)
	if err != nil {
		return response.ErrCodeCreateJWTTokenFailed, nil, fmt.Errorf("error for create token failed: %s", err)
	}

	out.Account = userManagerInfor.Account
	out.UserName = userManagerInfor.UserName

	return response.ErrCodeLoginSuccess, out, nil
}

func (m *ManagerLoginImpl) Register(ctx *gin.Context, in *vo.ManagerRegisterInput) (codeStatus int, err error) {
	// TODO: check email exists in user manager
	managerFound, err := m.sqlc.CheckUserManagerExistsByEmail(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("error for check manager already exists: %s", err)
	}

	if managerFound {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("manager already exists")
	}

	// TODO: check user spam / rate limiting by ip

	// TODO: create manager
	id := uuid.New().String()
	now := utiltime.GetTimeNow()
	hashPassword, err := crypto.HashPassword(in.UserPassword)
	if err != nil {
		return response.ErrCodeHashPasswordFailed, fmt.Errorf("hash password failed: %s", err)
	}

	err = m.sqlc.CreateUserManage(ctx, database.CreateUserManageParams{
		ID:        id,
		Account:   in.UserAccount,
		Password:  hashPassword,
		CreatedAt: now,
		UpdatedAt: now,
	})

	if err != nil {
		return response.ErrCodeRegisterFailed, fmt.Errorf("error for register manager failed: %s", err)
	}

	return response.ErrCodeRegisterSuccess, nil
}

func NewManagerLoginImpl(sqlc *database.Queries) *ManagerLoginImpl {
	return &ManagerLoginImpl{
		sqlc: sqlc,
	}
}
