package domain

import "github.com/golang-jwt/jwt"

type Claims struct {
    UId string `json:"uid"`
    jwt.StandardClaims
}