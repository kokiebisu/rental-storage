package coordinate

import (
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type Factory struct{}

func (f *Factory) New(coordinate float64) (ValueObject, *customerror.CustomError) {
	err := isValidCoordinate(coordinate)
	if err != nil {
		return ValueObject{}, err
	}
	return ValueObject{
		Value: coordinate,
	}, nil
}

func isValidCoordinate(value float64) *customerror.CustomError {
	if value < -180 {
		return customerror.ErrorHandler.InvalidValueError("coordinate", "cannot be less than -180")
	}
	return nil
}
