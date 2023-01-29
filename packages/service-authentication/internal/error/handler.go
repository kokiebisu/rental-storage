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

func (e *Handler) UnmarshalError(target string) *CustomError {
	msg := "unable to unmarshal " + target
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}

func (e *Handler) RequestFailError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("failed to send request to user service endpoint"),
	}
}

func (e *Handler) DecodeError(target string) *CustomError {
	msg := "unable to decode " + target
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}

func (e *Handler) PasswordHashError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("cannot hash password"),
	}
}

func (e *Handler) UndefinedEndPointError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("user service api endpoint not defined"),
	}
}

func (e *Handler) ResponseInvalidError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("response from user service was invalid"),
	}
}

func (e *Handler) SignTokenError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("unable to sign token"),
	}
}

func (e *Handler) ClaimParseError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("unable to parse with claims"),
	}
}

func (e *Handler) ClaimCastError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("unable to cast to domain claims"),
	}
}

func (e *Handler) ClaimUidEmptyError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("uid property in claims is empty string"),
	}
}

func (e *Handler) ClaimExpiredError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("claims should have been expired already"),
	}
}

func (e *Handler) CompareHashError() *CustomError {
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New("failed when comparing hash and password"),
	}
}