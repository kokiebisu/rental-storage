package domain

import "errors"

type Coordinate struct {
	Value float32
}

func NewCoordinate(coordinate float32) (Coordinate, error) {
	err := isValidCoordinate(coordinate)
	if err != nil {
		return Coordinate{}, err
	}
	return Coordinate{
		Value: coordinate,
	}, nil
}

func isValidCoordinate(value float32) error {
	if value < -180 {
		return errors.New("coordinate cannot be lower than 180")
	}
	return nil
}
