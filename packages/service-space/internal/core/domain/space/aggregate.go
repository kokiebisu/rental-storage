package space

import (
	"time"

	location "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/location"
)

type ImageUrl string

type Entity struct {
	UId         string
	Title       string
	LenderId    string
	Location    location.ValueObject
	ImageUrls   []ImageUrl
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type DTO struct {
	UId         string       `json:"uid"`
	Title       string       `json:"title"`
	LenderId    string       `json:"lenderId"`
	Location    location.DTO `json:"location"`
	ImageUrls   []string     `json:"imageUrls"`
	Description string       `json:"description"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   string       `json:"updatedAt"`
}

type Raw struct {
	UId         string       `json:"uid"`
	Title       string       `json:"title"`
	LenderId    string       `json:"lender_id"`
	Location    location.Raw `json:"space_location"`
	ImageUrls   []string     `json:"image_urls"`
	Description string       `json:"description"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func (r Raw) ToEntity() Entity {
	createdAt, _ := time.Parse(layoutISO, r.CreatedAt)
	updatedAt, _ := time.Parse(layoutISO, r.UpdatedAt)
	urls := []ImageUrl{}
	for _, url := range r.ImageUrls {
		urls = append(urls, ImageUrl(url))
	}
	return Entity{
		UId:         r.UId,
		Title:       r.Title,
		LenderId:    r.LenderId,
		Location:    r.Location.ToValueObject(),
		ImageUrls:   urls,
		Description: r.Description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func (e Entity) ToDTO() DTO {
	urls := []string{}
	for _, url := range e.ImageUrls {
		urls = append(urls, string(url))
	}
	return DTO{
		UId:         e.UId,
		Title:       e.Title,
		LenderId:    e.LenderId,
		Location:    e.Location.ToDTO(),
		ImageUrls:   urls,
		Description: e.Description,
		CreatedAt:   e.CreatedAt.Format(layoutISO),
		UpdatedAt:   e.UpdatedAt.Format(layoutISO),
	}
}

func (d DTO) ToEntity() Entity {
	createdAt, _ := time.Parse(layoutISO, d.CreatedAt)
	updatedAt, _ := time.Parse(layoutISO, d.UpdatedAt)

	urls := []ImageUrl{}
	for _, url := range d.ImageUrls {
		urls = append(urls, ImageUrl(url))
	}
	return Entity{
		UId:         d.UId,
		Title:       d.Title,
		LenderId:    d.LenderId,
		Location:    d.Location.ToValueObject(),
		ImageUrls:   urls,
		Description: d.Description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
