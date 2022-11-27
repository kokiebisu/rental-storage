package domain

import (
	"errors"
)

type Fee struct {
	Amount Amount
	Type   RentalFeeType
}

type FeeDTO struct {
	Amount AmountDTO `json:"amount"`
	Type   string    `json:"type"`
}

type FeeRaw struct {
	Amount AmountRaw
	Type   string
}

const (
	MONTHLY RentalFeeType = "MONTHLY"
)

type RentalFeeType string

func NewFee(feeCurrency string, feeAmount int64, feeType string) (Fee, error) {
	err := isValidFeeType(feeType)
	if err != nil {
		return Fee{}, err
	}
	validatedAmount, err := NewAmount(feeAmount, feeCurrency)
	if err != nil {
		return Fee{}, err
	}
	return Fee{
		Amount: validatedAmount,
		Type:   RentalFeeType(feeType),
	}, nil
}

func isValidFeeType(value string) error {
	if value == "" {
		return errors.New("fee type cannot be empty string")
	}
	return nil
}
