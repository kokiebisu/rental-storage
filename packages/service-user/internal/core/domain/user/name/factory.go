package name

import (
	"errors"
	"strings"
)

type Factory struct{}

func (f *Factory) New(nameType NameType, value string) ValueObject {
	err := validateName(value)
	if err != nil {
		panic(err)
	}

	return ValueObject{
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
