package service

import (
	"fmt"
	"service-payment/internal/core/domain"
	"service-payment/internal/core/port"
	"service-payment/internal/repository"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

type CustomerService struct {
	secretsManager port.SecretsManager
	repository     repository.CustomerRepository
}

func NewCustomerService(secretsManager port.SecretsManager, repository repository.CustomerRepository) (*CustomerService, error) {
	err := repository.SetupTables()
	if err != nil {
		return nil, err
	}
	return &CustomerService{
		secretsManager: secretsManager,
		repository:     repository,
	}, nil
}

func (s *CustomerService) CreateCustomer(userId string, firstName string, lastName string, emailAddress string) error {
	param, err := s.secretsManager.GetParam()
	if err != nil {
		return err
	}

	stripe.Key = *param.Parameter.Value
	params := &stripe.CustomerParams{
		// Description: stripe.String("My First Test Customer (created for API docs at https://www.stripe.com/docs/api)"),
		Name:  stripe.String(fmt.Sprintf("%s %s", firstName, lastName)),
		Email: stripe.String(emailAddress),
	}
	c, _ := customer.New(params)

	pc := domain.PaymentCustomer{
		UserId:       userId,
		CustomerId:   c.ID,
		ProviderType: "stripe",
	}

	_, err = s.repository.Save(pc)
	if err != nil {
		return err
	}
	return nil
}
