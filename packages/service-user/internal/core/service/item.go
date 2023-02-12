package service

import (
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type ItemService struct {
	itemRepository port.ItemRepository
	factory        port.ItemFactory
}

func (s *ItemService) AddItems(items []item.DTO) ([]string, *customerror.CustomError) {
	uids := []string{}
	for _, i := range items {
		item := s.factory.New(i.Name, i.OwnerId, i.SpaceId)
		uid, err := s.itemRepository.Save(item)
		if err != nil {
			return []string{}, err
		}
		uids = append(uids, uid)
	}
	return uids, nil
}
