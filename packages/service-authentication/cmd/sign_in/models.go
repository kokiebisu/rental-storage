package main

import "github.com/golang-jwt/jwt"

type SignInArgument struct {
	EmailAddress      string `json:"email"`
	Password   string `json:"password"`
}

type Payload struct {
    UId string `json:"uid"`
}

type Claims struct {
    UId string `json:"uid`
    jwt.StandardClaims
}