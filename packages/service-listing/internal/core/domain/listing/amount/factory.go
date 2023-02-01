package amount

import (
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type Factory struct{}

func isValidAmountValue(value int64) *errors.CustomError {
	if value < 0 {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}

func isValidAmountCurrency(value string) *errors.CustomError {
	if value == "" {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}

func (f *Factory) New(value int64, currency string) (ValueObject, *errors.CustomError) {
	err := isValidAmountValue(value)
	if err != nil {
		return ValueObject{}, errors.ErrorHandler.InternalServerError()
	}
	err = isValidAmountCurrency(currency)
	if err != nil {
		return ValueObject{}, errors.ErrorHandler.InternalServerError()
	}
	return ValueObject{
		Value:    value,
		Currency: CurrencyType(currency),
	}, nil
}
