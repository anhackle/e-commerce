package jwttoken

import (
	"time"

	"github.com/anle/codebase/global"
	"github.com/golang-jwt/jwt/v5"
)

func GenJWTToken(userID int, role string) (token string, err error) {
	mySigningKey := []byte(global.Config.JWT.Key)

	claims := JWTClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "ecommerce-backend",
		},
	}

	tokenGenerator := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenGenerator.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
