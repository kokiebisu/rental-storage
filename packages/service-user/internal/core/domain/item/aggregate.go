package item

import (
	"time"
)

type Entity struct {
	UId       string
	Name      string
	ImageUrls []string
	CreatedAt time.Time
	UpdatedAt time.Time
	OwnerId   string
	SpaceId   string
}

type DTO struct {
	UId       string   `json:"uid"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"imageUrls"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	OwnerId   string   `json:"ownerId"`
	SpaceId   string   `json:"spaceId"`
}

type Raw struct {
	UId       string   `json:"uid"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"image_urls"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	OwnerId   string   `json:"owner_id"`
	SpaceId   string   `json:"space_id"`
}

func (r Raw) ToEntity() Entity {
	createdAt, _ := time.Parse("YYYY-MM-DD", r.CreatedAt)
	updatedAt, _ := time.Parse("YYYY-MM-DD", r.UpdatedAt)
	return Entity{
		UId:       r.UId,
		Name:      r.Name,
		ImageUrls: r.ImageUrls,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		OwnerId:   r.OwnerId,
		SpaceId:   r.SpaceId,
	}
}
