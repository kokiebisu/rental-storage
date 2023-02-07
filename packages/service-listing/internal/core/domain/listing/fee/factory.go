package fee

import (
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	customerror "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

var (
	amountFactory = &amount.Factory{}
)

type Factory struct{}

func (f *Factory) New(feeCurrency string, feeAmount int64, feeType string) (ValueObject, *customerror.CustomError) {
	err := isValidFeeType(feeType)
	if err != nil {
		return ValueObject{}, err
	}
	validatedAmount, err := amountFactory.New(feeAmount, feeCurrency)
	if err != nil {
		return ValueObject{}, err
	}
	return ValueObject{
		Amount: validatedAmount,
		Type:   RentalFeeType(feeType),
	}, nil
}

func isValidFeeType(value string) *customerror.CustomError {
	if value == "" {
		return customerror.ErrorHandler.InvalidValueError("fee type", "cannot be empty")
	}
	return nil
}
