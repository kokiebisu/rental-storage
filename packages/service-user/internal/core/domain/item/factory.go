package item

import (
	"time"

	"github.com/google/uuid"
)

func New(name string, ownerId string, spaceId string) Entity {
	return Entity{
		UId:       uuid.New().String(),
		Name:      name,
		ImageUrls: []string{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
		OwnerId:   ownerId,
		SpaceId:   spaceId,
	}
}
