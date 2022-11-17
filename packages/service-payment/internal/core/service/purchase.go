package service

type PurchaseService struct{}

func NewPurchaseService() *PurchaseService {
	return &PurchaseService{}
}

func (s *PurchaseService) MakePayment() error {
	return nil
}
