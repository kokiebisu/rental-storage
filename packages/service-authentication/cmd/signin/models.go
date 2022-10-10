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

type User struct {
    Id int `json:"id"`
    Uid string `json:"uid"`
    EmailAddress string `json:"emailAddress"`
    Password string `json:"password"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    CreatedAt string `json:"createdAt"`
}

type Error struct {
    Message string `json:"message"`
}