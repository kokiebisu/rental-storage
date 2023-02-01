package emailaddress

import (
	"errors"
	"net/mail"
)

func New(value string) ValueObject {
	err := validateEmailAddress(value)
	if err != nil {
		panic(err)
	}
	return ValueObject{
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
