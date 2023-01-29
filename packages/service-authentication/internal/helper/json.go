package helper

import (
	"encoding/json"

	errors "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

func Stringify(data interface{}) (string, *errors.CustomError) {
	result, err := json.Marshal(data)
	if err != nil {
		return "", errors.ErrorHandler.CustomError("unable to marshal data in stringify")
	}
	return string(result), nil
}
