package services

import (
	"context"

	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/internal/vo"
)

// TODO: Refactor interface
type (
	IUserLogin interface {
		Register(ctx context.Context, in *vo.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *vo.VerifyOTPInput) (codeResult int, out *vo.VerifyOTPOutput, err error)
		UpdatePasswordRegister(ctx context.Context, in *vo.UpdatePasswordRegisterInput) (codeResult int, err error)
		Login(ctx context.Context, in *vo.LoginInput) (codeResult int, out *vo.LoginOutput, err error)
	}

	IUserInfo interface {
		GetUserById(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
	}

	ITest interface {
		// admin
		LoginAdmin(ctx context.Context, in *vo.LoginInput) (codeResult int, out *vo.LoginOutput, err error)

		// accommodation
		GetAccommodations(ctx context.Context)
		CreateAccommodation(ctx context.Context, in *vo.CreateAccommodationInput) (codeResult int, out *vo.CreateAccommodationOutput, err error)
	}
)

var (
	localUserAdmin IUserAdmin
	localUserLogin IUserLogin
	localUserInfo  IUserInfo
	test           ITest
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

func Test() ITest {
	if test == nil {
		panic("Implement test not found for interface ITest")
	}
	return test
}

func InitTest(i ITest) {
	test = i
}
