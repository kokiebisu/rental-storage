package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(hashedPassword string, plainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		log.Println(err)
		return false, nil
	}
	return true, nil
}