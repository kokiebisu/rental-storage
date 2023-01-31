package errors

import "errors"

func (e *Handler) GetParameterError() *CustomError {
	msg := "unable to find a parameter"
	err := errors.New(msg)
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}

func (e *Handler) ConvertError(target string, dataType string, err error) *CustomError {
	msg := "unable to convert " +  dataType + " value of " + target
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}

func (e *Handler) UnmarshalError(target string, err error) *CustomError {
	msg := "unable to unmarshal " + target
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}

func (e *Handler) MarshalError(err error) *CustomError {
	msg := "unable to marshal data"
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}

func (e *Handler) DbConfigurationError(err error) *CustomError {
	msg := "configuration error"
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}

func (e *Handler) InvalidValueError(target string, reason string) *CustomError {
	msg := "value of " + target + " is invalid. " + reason
	err := errors.New(msg)
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}

func (e *Handler) CreateTableError(target string, err error) *CustomError {
	msg := "creating table " + target + " failed"
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}

func (e *Handler) DeleteListingRowError(target string, err error) *CustomError {
	msg := "deleting listing from table " + target + " failed"
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}

func (e *Handler) FindListingsRowError(err error) *CustomError {
	msg := "unable to find a listing"
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}

func (e *Handler) ScanRowError(err error) *CustomError {
	msg := "unable to scan a row"
	return &CustomError{
		StatusCode: 500,
		Message:	msg,
		Reason:     err,
	}
}