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
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/validate"
)

type UserLoginImpl struct {
	sqlc *database.Queries
}

// Register implements services.IUserLogin.
func (u *UserLoginImpl) Register(ctx context.Context, in *vo.RegisterInput) (codeResult int, err error) {
	// !input: email, type, purpose
	// !1. validate email
	if !validate.IsValidEmail(in.VerifyKey) {
		return response.ErrCodeInvalidEmailFormat, fmt.Errorf("invalid email format")
	}

	// !2. check email exists in db (user has successfully registered)
	userFound, err := u.sqlc.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("error for check user already exists: %s", err)
	}

	if userFound > 0 {
		return response.ErrCodeUserAlreadyExists, fmt.Errorf("user already exists")
	}

	// !3. check user spam / rate limiting by ip

	// !4. hash email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// !5. check user already has active otp
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

	// !6. generate otp
	otpNew := random.GenerateOTP()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}

	// !7. save otp in redis
	err = global.Redis.SetEx(ctx, userKey, otpNew, time.Duration(consts.TIME_OTP_REGISTER*time.Minute)).Err()
	if err != nil {
		return response.ErrCodeSaveDataFailed, fmt.Errorf("save otp in redis failed: %s", err)
	}

	// !8. send otp to verify type (email, sms)
	switch in.VerifyType {
	case consts.EMAIL:
		// !8.1 Send to kafka -> send email service ->
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
		// 	fmt.Printf("Sent OTP to kafka error: %s", err)
		// 	return response.ErrCodeSendEmailOTP, err
		// }
		// fmt.Printf("Sent OTP to kafka success")

		// !8.2 send to email
		err = sendto.SendEmailOTP([]string{in.VerifyKey}, "otp_auth.html", map[string]interface{}{
			"otp": otpNew,
		})
		if err != nil {
			return response.ErrCodeSendEmailFailed, fmt.Errorf("send email failed: %s", err)
		}
	}

	// !9. save otp to mysql
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

	// !10. return
	return response.ErrCodeRegisterSuccess, nil
}

// VerifyOTP implements services.IUserLogin.
func (u *UserLoginImpl) VerifyOTP(ctx context.Context, in *vo.VerifyOTPInput) (codeResult int, out *vo.VerifyOTPOutput, err error) {
	out = &vo.VerifyOTPOutput{}

	// !1. hash email
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// !2. check otp already exists in redis
	otpFound, err := global.Redis.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return response.ErrCodeOTPNotExists, out, fmt.Errorf("error for check otp in redis: %s", err)
	}

	// !3. check otp match
	if in.VerifyCode != otpFound {
		return response.ErrCodeOTPNotMatch, out, fmt.Errorf("OTP not match")
	}

	// !4. if match, get info otp
	infoOTP, err := u.sqlc.GetUserUnverify(ctx, hashKey)
	if err != nil {
		return response.ErrCodeGetInfoOTPFailed, out, fmt.Errorf("get info otp failed: %s", err)
	}

	// !5. update user verify status and delete otp
	// UpdateUserVerifyStatus
	err = u.sqlc.UpdateUserVerifyStatus(ctx, database.UpdateUserVerifyStatusParams{
		UpdatedAt: utiltime.GetTimeNow(),
		KeyHash:   hashKey,
	})

	if err != nil {
		return response.ErrCodeUpdateUserVerifyFailed, out, fmt.Errorf("update user verify failed: %s", err)
	}

	// !6. return
	out.Token = infoOTP.KeyHash
	return response.ErrCodeVerifyOTPSuccess, out, nil
}

// UpdatePasswordRegister implements services.IUserLogin.
func (u *UserLoginImpl) UpdatePasswordRegister(ctx context.Context, in *vo.UpdatePasswordRegisterInput) (codeResult int, err error) {
	// !1. get info otp by key hash
	infoOTP, err := u.sqlc.GetUserVerified(ctx, in.Token)
	if err != nil {
		return response.ErrCodeGetInfoOTPFailed, fmt.Errorf("get info otp failed: %s", err)
	}
	// !2. check otp is verified
	if infoOTP.IsVerified == 0 {
		return response.ErrCodeOTPNotVerified, fmt.Errorf("user OTP not verified")
	}

	// !3. update user base
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
	userBase.UpdatedAt = now
	userBase.LoginTime = now
	userBase.LogoutTime = 0
	userBase.LoginIp = ""

	err = u.sqlc.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeSaveDataFailed, fmt.Errorf("save user base failed: %s", err)
	}

	// !4. create user info
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

// Login implements services.IUserLogin.
func (u *UserLoginImpl) Login(ctx context.Context, in *vo.LoginInput) (codeResult int, out *vo.LoginOutput, err error) {
	out = &vo.LoginOutput{}

	// !1. get user info
	userBase, err := u.sqlc.GetUserBaseByAccount(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, out, fmt.Errorf("get user info failed: %s", err)
	}

	// !2. check password match
	if !crypto.CheckPasswordHash(in.UserPassword, userBase.Password) {
		return response.ErrCodePasswordNotMatch, out, fmt.Errorf("dose not match password")
	}

	// !3. check two-factor authentication

	// !4. update login
	go u.sqlc.LoginUserBase(ctx, database.LoginUserBaseParams{
		LoginIp:   "",
		Account:   in.UserAccount,
		LoginTime: utiltime.GetTimeNow(),
	})

	// !5. create uuid user
	subToken := utils.GenerateCliTokenUUID(userBase.ID)

	// !6. get user info
	infoUser, err := u.sqlc.GetUserInfo(ctx, userBase.Account)
	if err != nil {
		return response.ErrCodeGetUserInfoFailed, out, fmt.Errorf("get user info failed: %s", err)
	}

	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeMarshalFailed, out, fmt.Errorf("convert to json failed: %v", err)
	}

	// !7. give info user json to redis
	err = global.Redis.SetEx(ctx, subToken, infoUserJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeSaveDataFailed, out, fmt.Errorf("save user info to redis failed: %s", err)
	}

	out.Token, err = auth.CreateToken(userBase.ID)
	if err != nil {
		return response.ErrCodeCreateJWTTokenFailed, out, err
	}

	return response.ErrCodeLoginSuccess, out, nil
}

func NewUserLoginImpl(sqlc *database.Queries) *UserLoginImpl {
	return &UserLoginImpl{
		sqlc: sqlc,
	}
}
