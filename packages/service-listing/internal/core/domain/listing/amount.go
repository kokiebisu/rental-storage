package domain

import "errors"

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

func NewAmount(value int64, currency string) (Amount, error) {
	err := isValidValue(value)
	if err != nil {
		return Amount{}, err
	}
	err = isValidCurrency(currency)
	if err != nil {
		return Amount{}, err
	}
	return Amount{
		Value:    value,
		Currency: CurrencyType(currency),
	}, nil
}

func isValidValue(value int64) error {
	if value < 0 {
		return errors.New("amount value cannot be negative")
	}
	return nil
}

func isValidCurrency(value string) error {
	if value == "" {
		return errors.New("amount currency cannot be empty string")
	}
	return nil
}
