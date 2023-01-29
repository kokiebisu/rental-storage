package helper

import (
	"encoding/json"

	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

func Stringify(data interface{}) (string, *errors.CustomError) {
	result, err := json.Marshal(data)
	if err != nil {
		return "", errors.ErrorHandler.CustomError("unable to marshal", err)
	}
	return string(result), nil
}
