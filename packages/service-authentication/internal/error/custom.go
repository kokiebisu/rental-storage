package customerror

type CustomError struct {
	StatusCode uint16
	ErrorCode  string
	Message    string
	Reason     error
}

type ErrorResponsePayload struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Reason    string `json:"reason"`
}

func (e *CustomError) Error() string {
	return e.Reason.Error()
}

func (e *CustomError) GetPayload() ErrorResponsePayload {
	return ErrorResponsePayload{
		ErrorCode: e.ErrorCode,
		Message:   e.Message,
		Reason:    e.Reason.Error(),
	}
}
