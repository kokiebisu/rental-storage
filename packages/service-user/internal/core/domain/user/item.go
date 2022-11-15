package domain

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	Uid       string
	Name      string
	ImageUrls []string
	CreatedAt time.Time
	UpdatedAt time.Time
	OwnerId   string
	ListingId string
}

type ItemDTO struct {
	Uid       string   `json:"uid"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"imageUrls"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	OwnerId   string   `json:"ownerId"`
	ListingId string   `json:"listingId"`
}

type ItemRaw struct {
	Uid       string   `json:"uid"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"image_urls"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	OwnerId   string   `json:"owner_id"`
	ListingId string   `json:"listing_id"`
}

func CreateItem(name string, ownerId string, listingId string) Item {
	return Item{
		Uid:       uuid.New().String(),
		Name:      name,
		ImageUrls: []string{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		OwnerId:   ownerId,
		ListingId: listingId,
	}
}
