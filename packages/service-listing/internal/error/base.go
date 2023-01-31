package errors

type CustomError struct {
	StatusCode uint16
	Message    string
	Reason     error
}

type Handler struct{}

var ErrorHandler *Handler

func (e *CustomError) Error() string {
	return e.Reason.Error()
}

func init() {
	ErrorHandler = &Handler{}
}
