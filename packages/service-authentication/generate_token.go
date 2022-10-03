package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// GenerateToken returns a unique token based on the provided email string
func GenerateToken(email string) string {
    hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Hash to store:", string(hash))

    return base64.StdEncoding.EncodeToString(hash)
}