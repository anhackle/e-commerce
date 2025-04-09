package jwttoken

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID string `json:"userID"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
