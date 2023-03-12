package errors

func (e *Handler) InternalServerError(msg string, err error) *CustomError {
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) LoggerConfigurationError(err error) *CustomError {
	return &CustomError{
		StatusCode: 500,
		Message:    "Logger configuration error",
		Reason:     err,
	}
}
