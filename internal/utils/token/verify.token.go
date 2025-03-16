package token

import (
	"fmt"

	"github.com/anle/codebase/global"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWTToken(tokenString string) (userID int, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.Key), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok {
		return claims.UserID, nil
	}

	return 0, fmt.Errorf("token invalid")

}
