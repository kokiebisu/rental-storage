package errors

func (e *Handler) LoggerConfigurationError(err error) *CustomError {
	return &CustomError{
		StatusCode: 500,
		Message:    "Logger configuration error",
		Reason:     err,
	}
}
