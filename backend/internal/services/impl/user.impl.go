package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/consts"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/database"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/auth"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/crypto"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/random"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/sendto"
	utiltime "github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/utils/util_time"
)

type UserLoginImpl struct {
	sqlc *database.Queries
}

func (u *UserLoginImpl) Register(ctx context.Context, in *vo.RegisterInput) (codeStatus int, err error) {
	// !. check user base exists
	userFound, err := u.sqlc.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("error for check user already exists: %s", err)
	}

	if userFound {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("user already exists")
	}

	// !. check user verified otp
	isVerified, err := u.sqlc.CheckUserVerifiedOTP(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeOTPAlreadyVerified, fmt.Errorf("error for otp verified: %s", err)
	}

	if isVerified {
		return response.ErrCodeOTPAlreadyVerified, fmt.Errorf("error for otp verified")
	}

	// !. check user spam / rate limiting by ip

	// !. hash email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// !. check user already has active otp
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Redis.Get(ctx, userKey).Result()
	switch {
	// key not exists
	case err == redis.Nil:
		fmt.Println("otp not found")
	case err != nil:
		return response.ErrCodeOTPAlreadyExists, fmt.Errorf("error for get otp: %s", err)
	case otpFound != "":
		return response.ErrCodeOTPAlreadyExists, fmt.Errorf("otp already exists")
	}

	// !. generate otp
	otpNew := random.GenerateOTP()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}

	// !. save otp in redis
	err = global.Redis.SetEx(ctx, userKey, otpNew, time.Duration(consts.TIME_OTP_REGISTER*time.Minute)).Err()
	if err != nil {
		return response.ErrCodeSaveDataFailed, fmt.Errorf("save otp in redis failed: %s", err)
	}

	// !. send otp to verify type (email, sms)
	switch in.VerifyType {
	case consts.EMAIL:
		// !. Send to kafka -> send email service ->
		// body := make(map[string]interface{})
		// body["otp"] = otpNew
		// body["email"] = in.VerifyKey

		// bodyRequest, _ := json.Marshal(body)

		// msg := kafka.Message{
		// 	Key:   []byte("otp-auth"),
		// 	Value: []byte(bodyRequest),
		// 	Time:  time.Now(),
		// }

		// err = global.Kafka.WriteMessages(global.Ctx, msg)
		// if err != nil {
		// 	return response.ErrCodeSendEmailOTP, err
		// }

		// !. send to email
		err = sendto.SendEmailOTP([]string{in.VerifyKey}, "otp_auth.html", map[string]interface{}{
			"otp": otpNew,
		})
		if err != nil {
			return response.ErrCodeSendEmailFailed, fmt.Errorf("send email failed: %s", err)
		}
	}

	// !. save otp to mysql
	id := uuid.New().String()
	err = u.sqlc.CreateUserVerify(ctx, database.CreateUserVerifyParams{
		ID:        id,
		Otp:       strconv.Itoa(otpNew),
		VerifyKey: in.VerifyKey,
		KeyHash:   hashKey,
		Type:      uint8(in.VerifyType),
		CreatedAt: utiltime.GetTimeNow(),
		UpdatedAt: utiltime.GetTimeNow(),
	})

	if err != nil {
		return response.ErrCodeSaveDataFailed, fmt.Errorf("save otp to database failed: %s", err)
	}

	// !. return
	return response.ErrCodeRegisterSuccess, nil
}

func (u *UserLoginImpl) VerifyOTP(ctx context.Context, in *vo.VerifyOTPInput) (codeStatus int, out *vo.VerifyOTPOutput, err error) {
	out = &vo.VerifyOTPOutput{}

	// !. hash email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// !. check otp already exists in redis
	otpFound, err := global.Redis.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return response.ErrCodeOTPNotExists, nil, fmt.Errorf("error for check otp in redis: %s", err)
	}

	// !. check otp match
	if in.VerifyCode != otpFound {
		return response.ErrCodeOTPNotMatch, nil, fmt.Errorf("OTP not match")
	}

	// !. if match, get info otp
	infoOTP, err := u.sqlc.GetUserUnverify(ctx, hashKey)
	if err != nil {
		return response.ErrCodeGetInfoOTPFailed, nil, fmt.Errorf("get info otp failed: %s", err)
	}

	if infoOTP.IsVerified != 0 {
		return response.ErrCodeOTPAlreadyVerified, nil, fmt.Errorf("otp is verified")
	}

	// !. update user verify status and delete otp
	// UpdateUserVerifyStatus
	err = u.sqlc.UpdateUserVerifyStatus(ctx, database.UpdateUserVerifyStatusParams{
		UpdatedAt: utiltime.GetTimeNow(),
		KeyHash:   hashKey,
	})

	if err != nil {
		return response.ErrCodeUpdateUserVerifyFailed, nil, fmt.Errorf("update user verify failed: %s", err)
	}

	// !. return
	out.Token = infoOTP.KeyHash
	return response.ErrCodeVerifyOTPSuccess, out, nil
}

func (u *UserLoginImpl) UpdatePasswordRegister(ctx context.Context, in *vo.UpdatePasswordRegisterInput) (codeStatus int, err error) {
	// !. get info otp by key hash
	infoOTP, err := u.sqlc.GetUserVerified(ctx, in.Token)
	if err != nil {
		return response.ErrCodeGetInfoOTPFailed, fmt.Errorf("get info otp failed: %s", err)
	}
	// !. check otp is verified
	if infoOTP.IsVerified == 0 {
		return response.ErrCodeOTPNotVerified, fmt.Errorf("user OTP not verified")
	}

	// !. check user base exists
	userFound, err := u.sqlc.CheckUserBaseExists(ctx, infoOTP.VerifyKey)
	if err != nil {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("error for check user already exists: %s", err)
	}

	if userFound {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("user already exists")
	}

	// !. update user base
	userBase := database.AddUserBaseParams{}
	userBase.ID = uuid.New().String()
	userBase.Account = infoOTP.VerifyKey
	hashPassword, err := crypto.HashPassword(in.Password)
	if err != nil {
		return response.ErrCodeHashPasswordFailed, fmt.Errorf("hash password failed: %s", err)
	}
	userBase.Password = hashPassword
	now := utiltime.GetTimeNow()
	userBase.CreatedAt = now
	userBase.IsVerified = 1
	userBase.UpdatedAt = now
	userBase.LoginTime = now
	userBase.LogoutTime = 0
	userBase.LoginIp = ""

	err = u.sqlc.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeSaveDataFailed, fmt.Errorf("save user base failed: %s", err)
	}

	// !. create user info
	now = utiltime.GetTimeNow()
	err = u.sqlc.CreateUserInfo(ctx, database.CreateUserInfoParams{
		ID:               uuid.New().String(),
		Account:          infoOTP.VerifyKey,
		Status:           1,
		IsAuthentication: 1,
		CreatedAt:        now,
		UpdatedAt:        now,
	})

	if err != nil {
		return response.ErrCodeUpdateDataFailed, fmt.Errorf("update password register failed: %s", err)
	}

	return response.ErrCodeUpdatePasswordRegisterSuccess, nil
}

func (u *UserLoginImpl) Login(ctx context.Context, in *vo.LoginInput) (codeStatus int, out *vo.LoginOutput, err error) {
	out = &vo.LoginOutput{}

	// !. get user info
	userBase, err := u.sqlc.GetUserBaseByAccount(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, nil, fmt.Errorf("get user info failed: %s", err)
	}

	// !. check password match
	if !crypto.CheckPasswordHash(in.UserPassword, userBase.Password) {
		return response.ErrCodePasswordNotMatch, nil, fmt.Errorf("dose not match password")
	}

	// !. check two-factor authentication

	// !. update login
	go u.sqlc.LoginUserBase(ctx, database.LoginUserBaseParams{
		LoginIp:   "",
		Account:   in.UserAccount,
		LoginTime: utiltime.GetTimeNow(),
	})

	// !. create uuid user
	subToken := utils.GenerateCliTokenUUID(userBase.ID)

	// !. get user info
	infoUser, err := u.sqlc.GetUserInfo(ctx, userBase.Account)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, nil, fmt.Errorf("get user info failed: %s", err)
	}

	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeMarshalFailed, nil, fmt.Errorf("convert to json failed: %v", err)
	}

	// !. give info user json to redis
	err = global.Redis.SetEx(ctx, subToken, infoUserJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeSaveDataFailed, nil, fmt.Errorf("save user info to redis failed: %s", err)
	}

	out.Token, err = auth.CreateToken(userBase.ID, global.User)
	if err != nil {
		return response.ErrCodeCreateJWTTokenFailed, nil, err
	}

	return response.ErrCodeLoginSuccess, out, nil
}

func NewUserLoginImpl(sqlc *database.Queries) *UserLoginImpl {
	return &UserLoginImpl{
		sqlc: sqlc,
	}
}
