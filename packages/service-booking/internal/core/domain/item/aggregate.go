package item

import (
	"time"
)

type Entity struct {
	Id        int32
	Name      string
	ImageUrls []string
	UserId    string
	ListingId string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DTO struct {
	Id        int32    `json:"id"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"imageUrls"`
}

type Raw struct {
	Id        int32    `json:"id"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"image_urls"`
}

func New(name string, imageUrls []string) (Entity, error) {
	return Entity{
		Name:      name,
		ImageUrls: imageUrls,
	}, nil
}

func (d DTO) ToEntity() (Entity, error) {
	return Entity{
		Id:        d.Id,
		Name:      d.Name,
		ImageUrls: d.ImageUrls,
	}, nil
}

func (e Entity) ToDTO() DTO {
	return DTO{
		Id:        e.Id,
		Name:      e.Name,
		ImageUrls: e.ImageUrls,
	}
}
