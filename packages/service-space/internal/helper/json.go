package helper

import (
	"encoding/json"

	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

func Stringify(data interface{}) (string, *customerror.CustomError) {
	result, err := json.Marshal(data)
	if err != nil {
		return "", customerror.ErrorHandler.MarshalError(err)
	}
	return string(result), nil
}
