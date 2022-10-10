package main

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(providedPassword string, userPassword string) (bool, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(providedPassword), 10)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(userPassword))
	if err != nil {
		return false, nil
	}
	return true, nil
}