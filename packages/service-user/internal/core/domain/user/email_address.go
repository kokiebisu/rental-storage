package user

import (
	"errors"
	"net/mail"
)

type EmailAddress struct {
	Value string
}

func CreateEmailAddress(value string) EmailAddress {
	err := validateEmailAddress(value)
	if err != nil {
		panic(err)
	}
	return EmailAddress{
		Value: value,
	}
}

func validateEmailAddress(value string) error {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return errors.New("not a valid email address")
	}
	return nil
}
