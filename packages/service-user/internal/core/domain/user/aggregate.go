package user

import (
	"time"

	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	emailaddress "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user/email_address"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user/name"
)

type Entity struct {
	UId          string
	FirstName    name.ValueObject
	LastName     name.ValueObject
	EmailAddress emailaddress.ValueObject
	Password     string
	Items        []item.Entity
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type DTO struct {
	UId          string     `json:"uid"`
	FirstName    string     `json:"firstName"`
	LastName     string     `json:"lastName"`
	EmailAddress string     `json:"emailAddress"`
	Password     string     `json:"password"`
	Items        []item.DTO `json:"items"`
	CreatedAt    string     `json:"createdAt"`
	UpdatedAt    string     `json:"updatedAt"`
}

type Raw struct {
	UId          string     `json:"uid"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	EmailAddress string     `json:"email_address"`
	Password     string     `json:"password"`
	Items        []item.Raw `json:"items"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
}

func (e *Entity) ToDTO() DTO {
	return DTO{
		UId:          e.UId,
		FirstName:    e.FirstName.Value,
		LastName:     e.LastName.Value,
		EmailAddress: e.EmailAddress.Value,
		Password:     e.Password,
		CreatedAt:    e.CreatedAt.Format("YYYY-MM-DD"),
		UpdatedAt:    e.UpdatedAt.Format("YYYY-MM-DD"),
	}
}

func (r *Raw) ToEntity() Entity {
	createdAt, _ := time.Parse("YYYY-MM-DD", r.CreatedAt)
	updatedAt, _ := time.Parse("YYYY-MM-DD", r.UpdatedAt)
	return Entity{
		UId: r.UId,
		FirstName: name.ValueObject{
			Value:    r.FirstName,
			NameType: name.FirstName,
		},
		LastName: name.ValueObject{
			Value:    r.LastName,
			NameType: name.LastName,
		},
		EmailAddress: emailaddress.ValueObject{Value: r.EmailAddress},
		Password:     r.Password,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
