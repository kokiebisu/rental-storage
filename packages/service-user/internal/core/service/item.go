package service

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ItemService struct {
	itemRepository port.ItemRepository
	factory        port.ItemFactory
}

func (s *ItemService) AddItems(items []item.DTO) *errors.CustomError {
	for _, i := range items {
		item := s.factory.New(i.Name, i.OwnerId, i.ListingId)
		err := s.itemRepository.Save(item)
		if err != nil {
			return err
		}
	}
	return nil
}
