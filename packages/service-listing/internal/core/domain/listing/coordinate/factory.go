package coordinate

import (
	err "errors"

	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type Factory struct{}

func (f *Factory) New(coordinate float32) (ValueObject, *errors.CustomError) {
	err := isValidCoordinate(coordinate)
	if err != nil {
		return ValueObject{}, err
	}
	return ValueObject{
		Value: coordinate,
	}, nil
}

func isValidCoordinate(value float32) *errors.CustomError {
	if value < -180 {
		return errors.ErrorHandler.InternalServerError(err.New("coordinate cannot be smaller than 180"))
	}
	return nil
}
