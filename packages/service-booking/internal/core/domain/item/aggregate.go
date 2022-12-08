package item

import (
	"errors"
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
	UserId    string   `json:"userId"`
	ListingId string   `json:"listingId"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

type Raw struct {
	Id        int32    `json:"id"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"image_urls"`
	UserId    string   `json:"user_id"`
	ListingId string   `json:"listing_id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func New(name string, imageUrls []string, userId string, listingId string) (Entity, error) {
	return Entity{
		Name:      name,
		ImageUrls: imageUrls,
		UserId:    userId,
		ListingId: listingId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}, nil
}

func (d DTO) ToEntity() (Entity, error) {
	layout := "01/02/2006 3:04:05 PM"
	createdAt, err := time.Parse(layout, d.CreatedAt)
	if err != nil {
		return Entity{}, errors.New("unable to parse CreatedAt")
	}
	updatedAt, err := time.Parse(layout, d.UpdatedAt)
	if err != nil {
		return Entity{}, errors.New("unable to parse UpdatedAt")
	}
	return Entity{
		Id:        d.Id,
		Name:      d.Name,
		ImageUrls: d.ImageUrls,
		UserId:    d.UserId,
		ListingId: d.ListingId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (e Entity) ToDTO() DTO {
	return DTO{
		Id:        e.Id,
		Name:      e.Name,
		ImageUrls: e.ImageUrls,
		UserId:    e.UserId,
		ListingId: e.ListingId,
		CreatedAt: e.CreatedAt.String(),
		UpdatedAt: e.CreatedAt.String(),
	}
}
