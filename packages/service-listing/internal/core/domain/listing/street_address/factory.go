package streetaddress

import (
	customerror "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type Factory struct{}

func (f *Factory) New(value string) (ValueObject, *customerror.CustomError) {
	err := isValidCity(value)
	if err != nil {
		return ValueObject{}, err
	}
	return ValueObject{
		Value: value,
	}, nil
}

func isValidCity(value string) *customerror.CustomError {
	// validation here
	if value == "" {
		return customerror.ErrorHandler.InvalidValueError("street address", "cannot be empty")
	}

	return nil
}
