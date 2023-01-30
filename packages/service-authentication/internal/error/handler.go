package errors



type Handler struct{}

var ErrorHandler *Handler

func init() {
	ErrorHandler = &Handler{}
}

func (e *Handler) CustomError(msg string, err error) *CustomError {
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) UnmarshalError(target string, err error) *CustomError {
	msg := "unable to unmarshal " + target
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) RequestFailError(err error) *CustomError {
	msg := "failed to send request to user service endpoint"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) DecodeError(target string, err error) *CustomError {
	msg := "unable to decode " + target
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) PasswordHashError(err error) *CustomError {
	msg := "cannot hash password"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) UndefinedEndPointError(err error) *CustomError {
	msg := "user service api endpoint not defined"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) ResponseInvalidError(err error) *CustomError {
	msg := "response from user service was invalid"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) SignTokenError(err error) *CustomError {
	msg := "unable to sign token"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) ClaimParseError(err error) *CustomError {
	msg := "unable to parse with claims"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) ClaimCastError(err error) *CustomError {
	msg := "unable to cast to domain claims"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) ClaimUidEmptyError(err error) *CustomError {
	msg := "uid property in claims is empty string"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) ClaimExpiredError(err error) *CustomError {
	msg := "claims should have been expired already"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}

func (e *Handler) CompareHashError(err error) *CustomError {
	msg := "failed when comparing hash and password"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:		err,
	}
}