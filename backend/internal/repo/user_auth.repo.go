package repo

import (
	"fmt"
	"time"

	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
)

type IUserAuthRepo interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepo struct {
}

// AddOTP implements IUserAuthRepo.
func (ur *userAuthRepo) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("user:%s:otp", email)
	return global.Redis.SetEx(global.Ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepo() IUserAuthRepo {
	return &userAuthRepo{}
}
