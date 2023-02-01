package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	emailaddress "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user/email_address"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user/name"
)

var (
	nameFactory = &name.Factory{}
)

type Factory struct{}

const TimeLayout = "2006-01-02 03:04:05"

func (f *Factory) New(uid string, firstName string, lastName string, emailAddress string, password string, items []item.Entity, createdAt string) Entity {
	var CreatedAt time.Time
	if uid == "" {
		uid = uuid.New().String()
	}
	if createdAt == "" {
		CreatedAt = time.Now()
	} else {
		c, _ := time.Parse(TimeLayout, createdAt)
		CreatedAt = c
	}
	return Entity{
		Uid:          uid,
		FirstName:    nameFactory.New(name.FirstName, firstName),
		LastName:     nameFactory.New(name.LastName, lastName),
		EmailAddress: emailaddress.New(emailAddress),
		Password:     password,
		Items:        items,
		CreatedAt:    CreatedAt,
	}
}

func (f *Factory) NewItem(name string, ownerId string, listingId string) item.Entity {
	return item.Entity{
		Uid:       uuid.New().String(),
		Name:      name,
		ImageUrls: []string{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		OwnerId:   ownerId,
		ListingId: listingId,
	}
}
