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

const TimeLayout = "2006-01-02 03:04:05"

func New(uid string, firstName string, lastName string, emailAddress string, password string, items []item.Entity, createdAt string) Entity {
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
		UId:          uid,
		FirstName:    nameFactory.New(name.FirstName, firstName),
		LastName:     nameFactory.New(name.LastName, lastName),
		EmailAddress: emailaddress.New(emailAddress),
		Password:     password,
		Items:        items,
		CreatedAt:    CreatedAt,
	}
}
