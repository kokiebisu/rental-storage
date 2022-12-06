package booking

import (
	"time"

	"github.com/google/uuid"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
)

type Entity struct {
	Id        string
	Status    BookingStatus
	Amount    amount.Entity
	OwnerId   string
	ListingId string
	Items     []item.Entity
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DTO struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	Amount    string `json:"amount"`
	OwnerId   string `json:"ownerId"`
	ListingId string `json:"listingId"`
	Items     string `json:"items"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Raw struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	Amount    string `json:"amount"`
	OwnerId   string `json:"owner_id"`
	ListingId string `json:"listing_id"`
	Items     string `json:"items"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

const (
	CREATED   BookingStatus = "CREATED"
	COMPLETED BookingStatus = "COMPLETED"
)

type BookingStatus string

func New(amount amount.Entity, ownerId string, listingId string, items []item.Entity) (booking.Entity, error) {
	return Entity{
		Id:        uuid.New().String(),
		Status:    CREATED,
		Amount:    amount,
		OwnerId:   ownerId,
		ListingId: listingId,
		Items:     items,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}, nil
}
