package imageurl

import (
	"errors"

	customerror "github.com/kokiebisu/rental-storage/service-listing/internal/error"
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
		return customerror.ErrorHandler.InternalServerError(errors.New("image url cannot be empty"))
	}
	return nil
}
