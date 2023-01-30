package item

import (
	"time"
)

type Entity struct {
	Uid       string
	Name      string
	ImageUrls []string
	CreatedAt time.Time
	UpdatedAt time.Time
	OwnerId   string
	ListingId string
}

type DTO struct {
	Uid       string   `json:"uid"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"imageUrls"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	OwnerId   string   `json:"ownerId"`
	ListingId string   `json:"listingId"`
}

type Raw struct {
	Uid       string   `json:"uid"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"image_urls"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	OwnerId   string   `json:"owner_id"`
	ListingId string   `json:"listing_id"`
}

func (r Raw) ToEntity() Entity {
	createdAt, _ := time.Parse("YYYY-MM-DD", r.CreatedAt)
	updatedAt, _ := time.Parse("YYYY-MM-DD", r.UpdatedAt)
	return Entity{
		Uid:       r.Uid,
		Name:      r.Name,
		ImageUrls: r.ImageUrls,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		OwnerId:   r.OwnerId,
		ListingId: r.ListingId,
	}
}
