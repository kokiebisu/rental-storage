package service

import (
	"log"

	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
)

type ItemService struct {
	itemRepository port.ItemRepository
}

func (s *ItemService) AddItems(items []domain.ItemDTO) {
	for _, i := range items {
		item := domain.CreateItem(i.Name, i.OwnerId, i.ListingId)
		err := s.itemRepository.Save(item)
		if err == nil {
			log.Fatal(err)
		}
	}
}
