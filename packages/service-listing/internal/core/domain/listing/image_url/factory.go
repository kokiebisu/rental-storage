package imageurl

import errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"

type Factory struct{}

func (f *Factory) New(url string) (ValueObject, *errors.CustomError) {
	err := isValidImageUrl(url)
	if err != nil {
		return ValueObject{}, err
	}
	return ValueObject{
		Value: url,
	}, nil
}

func isValidImageUrl(value string) *errors.CustomError {
	if value == "" {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}
