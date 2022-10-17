package main

import (
	"github.com/golang-jwt/jwt"
)

type Payload struct {
    UId string `json:"uid"`
}

type Claims struct {
    UId string `json:"uid`
    jwt.StandardClaims
}

// GenerateToken returns a unique token based on the provided email string
func GenerateJWTToken(payload *Payload) (string, error) {
    claims := Claims{
        payload.UId,
        jwt.StandardClaims{
            ExpiresAt: 100000000000000,
            Issuer: "test",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}