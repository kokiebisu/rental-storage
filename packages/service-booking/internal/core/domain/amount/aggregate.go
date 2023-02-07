package amount

import customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"

type ValueObject struct {
	Value    int32
	Currency string
}

type DTO struct {
	Value    int32  `json:"value"`
	Currency string `json:"currency"`
}

type Raw struct {
	Value    int32  `json:"value"`
	Currency string `json:"currency"`
}

func New(value int32, currency string) (ValueObject, *customerror.CustomError) {
	err := validateAmountValue(value)
	if err != nil {
		return ValueObject{}, err
	}
	err = validateAmountCurrency(currency)
	if err != nil {
		return ValueObject{}, err
	}
	return ValueObject{
		Value:    value,
		Currency: currency,
	}, nil
}

func validateAmountValue(value int32) *customerror.CustomError {
	if value < 0 {
		return customerror.ErrorHandler.InternalServerError("amount should be greater than 0", nil)
	}
	return nil
}

func validateAmountCurrency(value string) *customerror.CustomError {
	return nil
}

func (d DTO) ToEntity() ValueObject {
	return ValueObject(d)
}

func (e ValueObject) ToDTO() DTO {
	return DTO(e)
}
