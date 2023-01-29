package errors

type CustomError struct {
	StatusCode uint16
	Err        error
}

func (e *CustomError) Error() string {
	return e.Err.Error()
}
