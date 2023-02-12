package amount

import (
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type Factory struct{}

func isValidAmountValue(value int64) *customerror.CustomError {
	if value < 0 {
		return customerror.ErrorHandler.InvalidValueError("amount", "should be positive")
	}
	return nil
}

func isValidAmountCurrency(value string) *customerror.CustomError {
	if value == "" {
		return customerror.ErrorHandler.InvalidValueError("currency", "cannot be empty")
	}
	return nil
}

func (f *Factory) New(value int64, currency string) (ValueObject, *customerror.CustomError) {
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
