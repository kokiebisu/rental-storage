package amount

type ValueObject struct {
	Value    int64
	Currency CurrencyType
}

type DTO struct {
	Value    int64  `json:"value"`
	Currency string `json:"currency"`
}

type Raw struct {
	Value    int64
	Currency string
}

const (
	CAD CurrencyType = "CAD"
	USD CurrencyType = "USD"
)

type CurrencyType string
