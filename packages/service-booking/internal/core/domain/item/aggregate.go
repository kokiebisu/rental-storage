package item

import "time"

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
