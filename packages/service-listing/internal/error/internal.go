package errors

func (e *Handler) InternalServerError(err error) *CustomError {
	msg := "something went wrong"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}
