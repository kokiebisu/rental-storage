package domain

import errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"

type Amount struct {
	Value    int64
	Currency CurrencyType
}

type AmountDTO struct {
	Value    int64  `json:"value"`
	Currency string `json:"currency"`
}

type AmountRaw struct {
	Value    int64
	Currency string
}

const (
	CAD CurrencyType = "CAD"
	USD CurrencyType = "USD"
)

type CurrencyType string

func NewAmount(value int64, currency string) (Amount, *errors.CustomError) {
	err := isValidValue(value)
	if err != nil {
		return Amount{}, errors.ErrorHandler.InternalServerError()
	}
	err = isValidCurrency(currency)
	if err != nil {
		return Amount{}, errors.ErrorHandler.InternalServerError()
	}
	return Amount{
		Value:    value,
		Currency: CurrencyType(currency),
	}, nil
}

func isValidValue(value int64) *errors.CustomError {
	if value < 0 {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}

func isValidCurrency(value string) *errors.CustomError {
	if value == "" {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}
