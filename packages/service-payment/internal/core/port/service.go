package port

type CustomerService interface {
	CreateCustomer(userId string, firstName string, lastName string, emailAddress string) error
}

type PurchaseService interface {
	MakePayment() error
}
