package amount

import errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"

type Entity struct {
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

func New(value int32, currency string) (Entity, *errors.CustomError) {
	err := validateAmountValue(value)
	if err != nil {
		return Entity{}, err
	}
	err = validateAmountCurrency(currency)
	if err != nil {
		return Entity{}, err
	}
	return Entity{
		Value:    value,
		Currency: currency,
	}, nil
}

func validateAmountValue(value int32) *errors.CustomError {
	if value < 0 {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}

func validateAmountCurrency(value string) *errors.CustomError {
	return nil
}

func (d DTO) ToEntity() Entity {
	return Entity(d)
}

func (e Entity) ToDTO() DTO {
	return DTO(e)
}
