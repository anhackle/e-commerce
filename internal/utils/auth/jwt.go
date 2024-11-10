package auth

import (
	"time"

	"github.com/anle/codebase/global"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type PayloadClaims struct {
	jwt.RegisteredClaims
}

func GenTokenJWT(payload jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(global.Config.JWT.Secret))
}

func CreateToken(uuidToken string) (string, error) {
	timeEx := global.Config.JWT.Expire
	if timeEx == "" {
		timeEx = "1h"
	}

	expiration, err := time.ParseDuration(timeEx)
	if err != nil {
		return "", err
	}

	return GenTokenJWT(&PayloadClaims{
		jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(expiration)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			Issuer:    "shopdevgo",
			Subject:   uuidToken,
		},
	})
}
