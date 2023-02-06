package errors

import "errors"

func (e *Handler) GetParameterError(parameter string) *CustomError {
	msg := "unable to find a parameter " + parameter
	err := errors.New(msg)
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) ConvertError(target string, dataType string, err error) *CustomError {
	msg := "unable to convert " + dataType + " value of " + target
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) UnmarshalError(target string, err error) *CustomError {
	msg := "unable to unmarshal " + target
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) MarshalError(err error) *CustomError {
	msg := "unable to marshal data"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) DbConfigurationError(err error) *CustomError {
	msg := "unable to configure database"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) InvalidValueError(target string, reason string) *CustomError {
	msg := "value of " + target + " is invalid. " + target + " " + reason
	err := errors.New(msg)
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) CreateTableError(target string, err error) *CustomError {
	msg := "unable to create table " + target
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) DeleteListingRowError(target string, err error) *CustomError {
	msg := "unable to delete table " + target
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) FindListingsRowError(err error) *CustomError {
	msg := "unable to find the listing"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) ScanRowError(err error) *CustomError {
	msg := "unable to scan the row"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}
