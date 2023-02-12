package fee

import "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/amount"

type ValueObject struct {
	Amount amount.ValueObject
	Type   RentalFeeType
}

type DTO struct {
	Amount amount.DTO `json:"amount"`
	Type   string     `json:"type"`
}

type Raw struct {
	Amount amount.Raw
	Type   string
}

const (
	MONTHLY RentalFeeType = "MONTHLY"
)

type RentalFeeType string
