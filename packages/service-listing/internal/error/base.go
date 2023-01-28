package errors

type CustomError struct {
	StatusCode uint16
	Err        error
}

type Handler struct{}

var ErrorHandler *Handler

func (e *CustomError) Error() string {
	return e.Err.Error()
}

func init() {
	ErrorHandler = &Handler{}
}
