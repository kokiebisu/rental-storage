package user

import (
	"errors"
	"strings"
)

type NameType string

const (
	FirstName NameType = "firstName"
	LastName  NameType = "lastName"
)

type Name struct {
	Value    string
	NameType NameType
}

func CreateName(nameType NameType, value string) Name {
	err := validateName(value)
	if err != nil {
		panic(err)
	}

	return Name{
		Value:    capitalize(value),
		NameType: nameType,
	}
}

func validateName(value string) error {
	if value == "" {
		return errors.New("value is empty")
	}
	// edge case: check if it is only consisted of one word
	return nil
}

func capitalize(value string) string {
	return strings.ToUpper(value[:1]) + strings.ToLower(value[1:])
}
