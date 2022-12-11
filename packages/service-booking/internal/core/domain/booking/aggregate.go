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
	Id        string     `json:"id"`
	Status    string     `json:"status"`
	Amount    amount.DTO `json:"amount"`
	OwnerId   string     `json:"ownerId"`
	ListingId string     `json:"listingId"`
	Items     []item.DTO `json:"items"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
}

type Raw struct {
	Id        string     `json:"id"`
	Status    string     `json:"status"`
	Amount    amount.Raw `json:"amount"`
	OwnerId   string     `json:"owner_id"`
	ListingId string     `json:"listing_id"`
	Items     []item.Raw `json:"items"`
	CreatedAt string     `json:"created_at"`
	UpdatedAt string     `json:"updated_at"`
}

const (
	CREATED   BookingStatus = "CREATED"
	COMPLETED BookingStatus = "COMPLETED"
)

type BookingStatus string

func New(amount amount.Entity, ownerId string, listingId string, items []item.Entity) (Entity, error) {
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

func (d *DTO) ToEntity() (Entity, error) {
	amountEntity, err := amount.New(d.Amount.Value, d.Amount.Currency)
	if err != nil {
		return Entity{}, err
	}
	items := []item.Entity{}
	for _, i := range d.Items {
		itemEntity, err := i.ToEntity()
		if err != nil {
			return Entity{}, err
		}
		items = append(items, itemEntity)
	}
	return Entity{
		Id:        d.Id,
		Status:    BookingStatus(d.Status),
		Amount:    amountEntity,
		OwnerId:   d.OwnerId,
		ListingId: d.ListingId,
		Items:     items,
	}, nil
}

func (e *Entity) ToDTO() DTO {
	createdAtString := e.CreatedAt.Format("2016-08-19")
	updatedAtString := e.UpdatedAt.Format("2016-08-19")
	itemsDTO := []item.DTO{}
	for _, i := range e.Items {
		itemsDTO = append(itemsDTO, i.ToDTO())
	}
	return DTO{
		Id:        e.Id,
		Status:    string(e.Status),
		Amount:    e.Amount.ToDTO(),
		OwnerId:   e.OwnerId,
		ListingId: e.ListingId,
		Items:     itemsDTO,
		CreatedAt: createdAtString,
		UpdatedAt: updatedAtString,
	}
}
