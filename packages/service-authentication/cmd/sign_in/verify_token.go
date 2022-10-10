package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func verifyJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, errors.New("couldn't parse")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("Couldn't parse claims")
	}
	if claims.UId == "" {
		return nil, errors.New("UID is empty")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return nil, errors.New("JWT is expired")
	}
	return claims, nil
}