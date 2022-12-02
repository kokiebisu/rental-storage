package domain

import "errors"

type StreetAddress struct {
	Value string
}

func NewStreetAddress(value string) (StreetAddress, error) {
	err := isValidCity(value)
	if err != nil {
		return StreetAddress{}, err
	}
	return StreetAddress{
		Value: value,
	}, nil
}

func isValidCity(value string) error {
	// validation here
	if value == "" {
		return errors.New("street address cannot be empty")
	}

	return nil
}
