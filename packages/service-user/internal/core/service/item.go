package service

import (
	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ItemService struct {
	itemRepository port.ItemRepository
}

func (s *ItemService) AddItems(items []domain.ItemDTO) *errors.CustomError {
	for _, i := range items {
		item := domain.CreateItem(i.Name, i.OwnerId, i.ListingId)
		err := s.itemRepository.Save(item)
		if err == nil {
			if errors.ErrorHandler != nil {
				return errors.ErrorHandler.InternalServerError()
			}

		}
	}
	return nil
}
