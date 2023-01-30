package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
)

type Factory struct{}

const TimeLayout = "2006-01-02 03:04:05"

func (f *Factory) NewUser(uid string, firstName string, lastName string, emailAddress string, password string, items []item.Entity, createdAt string) user.Entity {
	fmt.Println("INSIDE UID: ", uid)
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
	return user.Entity{
		Uid:          uid,
		FirstName:    domain.CreateName(domain.FirstName, firstName),
		LastName:     domain.CreateName(domain.LastName, lastName),
		EmailAddress: domain.CreateEmailAddress(emailAddress),
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
