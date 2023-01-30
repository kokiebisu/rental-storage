package errors

type CustomError struct {
	StatusCode uint16
	Message    string
	Reason     error
}

type ErrorResponsePayload struct {
	Message string `json:"message"`
	Reason  error  `json:"reason"`
}

func (e *CustomError) Error() string {
	return e.Reason.Error()
}

func (e *CustomError) GetPayload() ErrorResponsePayload {
	return ErrorResponsePayload{
		Message: e.Message,
		Reason: e.Reason,
	}
}