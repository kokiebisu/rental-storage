package port

type Event struct {
	Field string `json:"field"`
	Arguments interface {} `json:"arguments"`
	Identity Identity `json:"identity"`
}

type Identity struct {
	ResolverContext ResolverContext
}

type CreatePaymentCustomerBody struct {
	UserId string `json:"userId"`
	EmailAddress string `json:"emailAddress"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type ResolverContext struct {
	UId string `json:"uid"`
}