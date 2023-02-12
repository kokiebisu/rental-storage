package imageurl

import (
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type Factory struct{}

func (f *Factory) New(url string) (ValueObject, *customerror.CustomError) {
	err := isValidImageUrl(url)
	if err != nil {
		return ValueObject{}, err
	}
	return ValueObject{
		Value: url,
	}, nil
}

func isValidImageUrl(value string) *customerror.CustomError {
	if value == "" {
		return customerror.ErrorHandler.InvalidValueError("image url", "cannot be empty")
	}
	return nil
}
