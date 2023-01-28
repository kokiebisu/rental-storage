package domain

import errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"

type Coordinate struct {
	Value float32
}

func NewCoordinate(coordinate float32) (Coordinate, *errors.CustomError) {
	err := isValidCoordinate(coordinate)
	if err != nil {
		return Coordinate{}, err
	}
	return Coordinate{
		Value: coordinate,
	}, nil
}

func isValidCoordinate(value float32) *errors.CustomError {
	if value < -180 {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}
