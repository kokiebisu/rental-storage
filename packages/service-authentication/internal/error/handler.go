package customerror

type Handler struct{}

var ErrorHandler *Handler

func init() {
	ErrorHandler = &Handler{}
}

func (e *Handler) UnmarshalError(target string, err error) *CustomError {
	msg := "unable to unmarshal " + target
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "UNMARSHAL_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) RequestFailError(err error) *CustomError {
	msg := "failed to send request to user service endpoint"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "REQUEST_FAIL_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) RequestInternalError(err error) *CustomError {
	msg := "request returned an internal server error"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "REQUEST_INTERNAL_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) DecodeError(target string, err error) *CustomError {
	msg := "unable to decode " + target
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "DECODE_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) PasswordHashError(err error) *CustomError {
	msg := "cannot hash password"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "PASSWORD_HASH_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) EnvironmentVariableError(err error, variableName string) *CustomError {
	msg := variableName + " not defined"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "ENVIRONMENT_VARIABLE_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) ResponseInvalidError(err error) *CustomError {
	msg := "response from user service was invalid"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "RESPONSE_INVALID_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) SignTokenError(err error) *CustomError {
	msg := "unable to sign token"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "SIGN_TOKEN_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) ClaimParseError(err error) *CustomError {
	msg := "unable to parse with claims"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "CLAIM_PARSE_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) ClaimCastError(err error) *CustomError {
	msg := "unable to cast to domain claims"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "CLAIM_CAST_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) ClaimUIdEmptyError(err error) *CustomError {
	msg := "uid property in claims is empty string"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "CLAIM_UID_EMPTY_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) ClaimExpiredError(err error) *CustomError {
	msg := "claims should have been expired already"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "CLAIM_EXPIRED_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) CompareHashError(err error) *CustomError {
	msg := "failed when comparing hash and password"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "COMPARE_HASH_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) PasswordGenerationError(err error) *CustomError {
	msg := "unable to generate password"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "PASSWORD_GENERATION_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) LoggerConfigurationError(err error) *CustomError {
	msg := "Logger configuration error"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "LOGGER_CONFIGURATION_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) RedisConnectionError(err error) *CustomError {
	msg := "Redis connection error"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "REDIS_CONNECTION_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) RedisTokenStoreError(err error) *CustomError {
	msg := "Redis token store error"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "REDIS_TOKEN_STORE_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) PasswordNotMatchedError(err error) *CustomError {
	msg := "password not matched"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "PASSWORD_NOT_MATCHED_ERROR",
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) TokenGenerationError(err error) *CustomError {
	msg := "unable to generate token"
	return &CustomError{
		StatusCode: 500,
		ErrorCode:  "TOKEN_GENERATION_ERROR",
		Message:    msg,
		Reason:     err,
	}
}
