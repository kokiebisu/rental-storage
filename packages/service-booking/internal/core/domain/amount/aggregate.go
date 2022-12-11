package amount

import "errors"

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

func New(value int32, currency string) (Entity, error) {
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

func validateAmountValue(value int32) error {
	if value < 0 {
		return errors.New("amount value must be greater than 0")
	}
	return nil
}

func validateAmountCurrency(value string) error {
	return nil
}

func (d DTO) ToEntity() Entity {
	return Entity{
		Value:    d.Value,
		Currency: d.Currency,
	}
}

func (e Entity) ToDTO() DTO {
	return DTO{
		Value:    e.Value,
		Currency: e.Currency,
	}
}
