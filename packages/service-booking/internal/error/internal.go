package errors

import "errors"

func (e *Handler) InternalServerError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("something went wrong"),
	}
}
