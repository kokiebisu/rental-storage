package item

import (
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type Entity struct {
	Id        uint32
	Name      string
	ImageUrls []string
}

type DTO struct {
	Id        uint32   `json:"id"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"imageUrls"`
}

type Raw struct {
	Id        uint32   `json:"id"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"image_urls"`
}

func New(id uint32, name string, imageUrls []string) (Entity, *customerror.CustomError) {
	return Entity{
		Id:        id,
		Name:      name,
		ImageUrls: imageUrls,
	}, nil
}

func (d DTO) ToEntity() Entity {
	return Entity(d)
}

func (e Entity) ToDTO() DTO {
	return DTO(e)
}
