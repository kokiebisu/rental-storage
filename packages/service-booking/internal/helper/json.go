package helper

import (
	"encoding/json"

	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

func Stringify(data interface{}) (string, *customerror.CustomError) {
	result, err := json.Marshal(data)
	if err != nil {
		return "", customerror.ErrorHandler.InternalServerError("unable to marshal", err)
	}
	return string(result), nil
}
