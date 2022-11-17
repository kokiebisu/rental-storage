package domain

type PaymentCustomer struct {
	UserId       string
	CustomerId   string
	ProviderType string
}

type PaymentCustomerDTO struct {
	UserId       string `json:"userId"`
	CustomerId   string `json:"customerId"`
	ProviderType string `json:"providerType"`
}

type PaymentCustomerRaw struct {
	UserId       string `json:"user_id"`
	CustomerId   string `json:"customer_id"`
	ProviderType string `json:"provider_type"`
}

func NewPaymentCustomer(userId string, customerId string, providerType string) PaymentCustomer {
	return PaymentCustomer{
		UserId:       userId,
		CustomerId:   customerId,
		ProviderType: providerType,
	}
}

func (c *PaymentCustomerRaw) ToEntity() PaymentCustomer {
	return PaymentCustomer{
		UserId:       c.UserId,
		CustomerId:   c.CustomerId,
		ProviderType: c.ProviderType,
	}
}
