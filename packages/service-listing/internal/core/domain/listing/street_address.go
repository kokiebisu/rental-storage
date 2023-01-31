package domain

import errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"

type StreetAddress struct {
	Value string
}

func NewStreetAddress(value string) (StreetAddress, *errors.CustomError) {
	err := isValidCity(value)
	if err != nil {
		return StreetAddress{}, err
	}
	return StreetAddress{
		Value: value,
	}, nil
}

func isValidCity(value string) *errors.CustomError {
	// validation here
	if value == "" {
		// city name
		return errors.ErrorHandler.InvalidValueError("street address", "street address cannot be empty")
	}

	return nil
}
