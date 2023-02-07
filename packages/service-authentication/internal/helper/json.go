package helper

import (
	"encoding/json"
	"io"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

func Stringify(data interface{}) (string, *customerror.CustomError) {
	result, err := json.Marshal(data)
	if err != nil {
		return "", customerror.ErrorHandler.UnmarshalError("data in stringify", err)
	}
	return string(result), nil
}

func Decode(decodeFrom io.Reader, decodeTo interface{}) (interface{}, *customerror.CustomError) {
	if err := json.NewDecoder(decodeFrom).Decode(&decodeTo); err != nil {
		return "", customerror.ErrorHandler.DecodeError("statusCode and message from status code 500 response", err)
	}
	return decodeTo, nil
}
