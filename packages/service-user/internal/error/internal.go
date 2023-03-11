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

func (e *Handler) DeleteRowError(target string, err error) *CustomError {
	msg := "unable to delete table " + target
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) FindRowsError(err error) *CustomError {
	msg := "unable to find the space"
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

func (e *Handler) InsertRowError(table string, err error) *CustomError {
	msg := "unable to insert to " + table
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) InvalidParamError(err error) *CustomError {
	msg := "unable to extract any param"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) EventPublishError(err error) *CustomError {
	msg := "unable to publish event to kinesis stream"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) CreateStreamError(err error) *CustomError {
	msg := "unable to create kinesis stream"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) PutRecordError(err error) *CustomError {
	msg := "unable to put record"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}

func (e *Handler) LoggerConfigurationError(err error) *CustomError {
	msg := "unable to configure logger"
	return &CustomError{
		StatusCode: 500,
		Message:    msg,
		Reason:     err,
	}
}
