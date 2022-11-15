package service

import (
	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
)

type ItemService struct {
	itemRepository port.ItemRepository
}

func (s *ItemService) AddItems(items []domain.ItemDTO) {
	for _, i := range items {
		item := domain.CreateItem(i.Name, i.OwnerId, i.ListingId)
		s.itemRepository.Save(item)
	}
}
