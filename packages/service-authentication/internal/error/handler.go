package errors

import "errors"

type Handler struct{}

var ErrorHandler *Handler

func init() {
	ErrorHandler = &Handler{}
}

func (e *Handler) CustomError(msg string) *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}
