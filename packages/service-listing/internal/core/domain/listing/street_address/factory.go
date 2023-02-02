package streetaddress

import (
	err "errors"

	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type Factory struct{}

func (f *Factory) New(value string) (ValueObject, *errors.CustomError) {
	err := isValidCity(value)
	if err != nil {
		return ValueObject{}, err
	}
	return ValueObject{
		Value: value,
	}, nil
}

func isValidCity(value string) *errors.CustomError {
	// validation here
	if value == "" {
		return errors.ErrorHandler.InternalServerError(err.New("city cannot be empty value"))
	}

	return nil
}