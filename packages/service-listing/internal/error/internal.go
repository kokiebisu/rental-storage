package errors

import "errors"

func (e *Handler) GetParameterError() *CustomError {
	msg := "unable to find a parameter"
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}

func (e *Handler) StringConvertError(target string) *CustomError {
	msg := "unable to convert string value of " + target + " to a number"
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

func (e *Handler) MarshalError() *CustomError {
	msg := "unable to marshal data"
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}

func (e *Handler) DbConfigurationError() *CustomError {
	msg := "configuration error"
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}

func (e *Handler) InvalidValueError(target string) *CustomError {
	msg := "value of " + target + " is invalid"
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}

func (e *Handler) CreateTableError(target string) *CustomError {
	msg := "creating table " + target + " failed"
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}

func (e *Handler) DeleteListingError(target string) *CustomError {
	msg := "deleting listing from table " + target + " failed"
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}

func (e *Handler) FindListingError() *CustomError {
	msg := "unable to find a listing"
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}

func (e *Handler) ScanRowError() *CustomError {
	msg := "unable to scan a row"
	return &CustomError{
		StatusCode: 500,
		Err:        errors.New(msg),
	}
}