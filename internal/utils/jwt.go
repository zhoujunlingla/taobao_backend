package utils

import (
	"fmt"
	"taobao_backend/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type myCustomClaims struct {
	Id int
	jwt.StandardClaims
}

func Jwt_generate(id int) string {
	expireTime := time.Now().Add(12 * time.Hour)
	claims := &myCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.Cfg.Secret
	token_string, _ := token.SignedString([]byte(secret))

	return token_string
}

func Jwt_verify(tokenString string) (int, string, int) {
	secret := config.Cfg.Secret
	var claims myCustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*myCustomClaims); ok && token.Valid {
		return 0, "登录成功", claims.Id
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return -1, "That's not even a token", 0
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return -1, "Timing is everything", 0
		} else {
			return -1, fmt.Sprintln("Couldn't handle this token:", err), 0
		}
	} else {
		return -1, fmt.Sprintln("Couldn't handle this token:", err), 0
	}
}
