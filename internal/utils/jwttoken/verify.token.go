package jwttoken

import (
	"fmt"

	"github.com/anle/codebase/global"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWTToken(tokenString string) (userID string, role string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.JWT.Key), nil
	})
	if err != nil {
		return userID, role, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok {
		return claims.UserID, claims.Role, nil
	}

	return userID, role, fmt.Errorf("token invalid")

}
