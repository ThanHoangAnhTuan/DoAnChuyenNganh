package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/consts"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/auth"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/crypto"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
)

type ManagerLoginImpl struct {
	sqlc *database.Queries
}

func (m *ManagerLoginImpl) Login(ctx context.Context, in *vo.ManagerLoginInput) (codeStatus int, out *vo.ManagerLoginOutput, err error) {
	out = &vo.ManagerLoginOutput{}

	// !. get manager info
	userManager, err := m.sqlc.GetUserManager(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, nil, fmt.Errorf("get user info failed: %s", err)
	}

	// !. check password match
	if !crypto.CheckPasswordHash(in.UserPassword, userManager.Password) {
		return response.ErrCodePasswordNotMatch, nil, fmt.Errorf("dose not match password")
	}

	// !. check two-factor authentication

	// !. update login
	go m.sqlc.UpdateUserManagerLogin(ctx, database.UpdateUserManagerLoginParams{
		LoginTime: utiltime.GetTimeNow(),
		Account:   in.UserAccount,
	})

	// !. create uuid user
	subToken := utils.GenerateCliTokenUUID(userManager.ID)

	userManagerInfor := vo.ManagerInfor{
		Account:  userManager.Account,
		UserName: userManager.UserName,
	}

	userManagerInforJson, err := json.Marshal(userManagerInfor)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("convert to json failed: %v", err)
	}

	// !. save manager info to redis
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

func (m *ManagerLoginImpl) Register(ctx context.Context, in *vo.ManagerRegisterInput) (codeStatus int, err error) {
	// !. check email exists in user manager
	managerFound, err := m.sqlc.CheckUserManagerExistsByEmail(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("error for check manager already exists: %s", err)
	}

	if managerFound {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("manager already exists")
	}

	// !. check user spam / rate limiting by ip

	// !. create manager
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
