package user

import (
	"time"

	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
)

type EmailAddress string
type Name string

type Entity struct {
	UId          string
	FirstName    Name
	LastName     Name
	EmailAddress EmailAddress
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

var layoutISO = "2006-01-02 15:04:05"

func (e *Entity) ToDTO() DTO {
	return DTO{
		UId:          e.UId,
		FirstName:    string(e.FirstName),
		LastName:     string(e.LastName),
		EmailAddress: string(e.EmailAddress),
		Password:     e.Password,
		Items:        []item.DTO{},
		CreatedAt:    e.CreatedAt.Format(layoutISO),
		UpdatedAt:    e.UpdatedAt.Format(layoutISO),
	}
}

func (r *DTO) ToEntity() Entity {
	createdAt, _ := time.Parse(layoutISO, r.CreatedAt)
	updatedAt, _ := time.Parse(layoutISO, r.UpdatedAt)

	return Entity{
		UId:          r.UId,
		FirstName:    Name(r.FirstName),
		LastName:     Name(r.LastName),
		EmailAddress: EmailAddress(r.EmailAddress),
		Password:     r.Password,
		Items:        []item.Entity{},
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func (r *Raw) ToEntity() Entity {
	createdAt, _ := time.Parse(layoutISO, r.CreatedAt)
	updatedAt, _ := time.Parse(layoutISO, r.UpdatedAt)

	return Entity{
		UId:          r.UId,
		FirstName:    Name(r.FirstName),
		LastName:     Name(r.LastName),
		EmailAddress: EmailAddress(r.EmailAddress),
		Password:     r.Password,
		Items:        []item.Entity{},
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
