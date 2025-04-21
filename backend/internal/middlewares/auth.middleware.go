package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/global"
	"github.com/thanhoanganhtuan/go-ecommerce-backend-api/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response.ErrorResponse(ctx, response.ErrCodeMissAuthorizationHeader)
			ctx.Abort()
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			response.ErrorResponse(ctx, response.ErrCodeInvalidAuthorizationFormat)
			ctx.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, prefix)

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.Config.JWT.Api_secret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil || !token.Valid {
			response.ErrorResponse(ctx, response.ErrCodeInvalidToken)
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("userId", claims["sub"])
		} else {
			response.ErrorResponse(ctx, response.ErrCodeInvalidToken)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
