package amount

import (
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type Factory struct{}

func isValidAmountValue(value int64) *errors.CustomError {
	if value < 0 {
		return errors.ErrorHandler.InvalidValueError("amount", "the value should be positive")
	}
	return nil
}

func isValidAmountCurrency(value string) *errors.CustomError {
	if value == "" {
		return errors.ErrorHandler.InvalidValueError("currency", "the currency cannot be empty")
	}
	return nil
}

func (f *Factory) New(value int64, currency string) (ValueObject, *errors.CustomError) {
	err := isValidAmountValue(value)
	if err != nil {
		return ValueObject{}, err
	}
	err = isValidAmountCurrency(currency)
	if err != nil {
		return ValueObject{}, err
	}
	return ValueObject{
		Value:    value,
		Currency: CurrencyType(currency),
	}, nil
}
