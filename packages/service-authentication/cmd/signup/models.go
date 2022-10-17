package main

import "github.com/golang-jwt/jwt"

type UserCreationPayload struct {
	EmailAddress      string `json:"emailAddress"`
	Password   string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type AuthorizationTokenPayload struct {
	AuthorizationToken string `json:"authorizationToken"`
}

type Payload struct {
    UId string `json:"uid"`
}

type Claims struct {
    UId string `json:"uid"`
    jwt.StandardClaims
}