package helper

import (
	"bytes"
	"net/http"

	customerror "github.com/kokiebisu/rental-storage/service-authentication/internal/error"
)

func SendPostRequest(endpoint string, encoded string) (*http.Response, *customerror.CustomError) {
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer([]byte(encoded)))
	if err != nil {
		return nil, customerror.ErrorHandler.RequestFailError(err)
	}
	return resp, nil
}
