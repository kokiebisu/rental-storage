package errors

import "errors"

type Handler struct{}

var ErrorHandler *Handler

func init() {
	ErrorHandler = &Handler{}
}

func (e *Handler) PasswordNotMatchedError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("password didn't match"),
	}
}

func (e *Handler) CompareHashAndPasswordError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("compare hash and password error"),
	}
}

func (e *Handler) GenerateJWTTokenError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("generate jwt token error"),
	}
}

func (e *Handler) BodyRequestParseError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("body request parsing failed"),
	}
}

func (e *Handler) JsonDecoderError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("json decoder error"),
	}
}

func (e *Handler) HttpInternalRequestError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("http internal server error"),
	}
}

func (e *Handler) MarshalError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("marshal error"),
	}
}

func (e *Handler) InternalServerError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("something went wrong"),
	}
}

func (e *Handler) CustomError(msg string) *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}
