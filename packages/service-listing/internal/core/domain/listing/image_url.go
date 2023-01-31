package domain

import errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"

type ImageUrl struct {
	Value string
}

func NewImageUrl(url string) (ImageUrl, *errors.CustomError) {
	err := isValidImageUrl(url)
	if err != nil {
		return ImageUrl{}, err
	}
	return ImageUrl{
		Value: url,
	}, nil
}

func isValidImageUrl(value string) *errors.CustomError {
	if value == "" {
		return errors.ErrorHandler.InvalidValueError("image url")
	}
	return nil
}
