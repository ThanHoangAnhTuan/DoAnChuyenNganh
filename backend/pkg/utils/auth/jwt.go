package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/thanhoanganhtuan/DoAnChuyenNganh/global"
)

type payload struct {
	jwt.RegisteredClaims
	Role string `json:"role,omitempty"`
}

func GenerateTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.Api_secret))
}

func CreateToken(uuidToken string, role string) (string, error) {
	timeEx := global.Config.JWT.JWT_Expiration
	if timeEx == "" {
		timeEx = "1h"
	}
	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}
	expiresAt := jwt.NewNumericDate(time.Now().Add(expiration))
	issuedAt := jwt.NewNumericDate(time.Now())

	return GenerateTokenJWT(&payload{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
			Issuer:    "thantuan",
			Subject:   uuidToken,
		},
		Role: role,
	})
}
