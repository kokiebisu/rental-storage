package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Uid          string
	FirstName    Name
	LastName     Name
	EmailAddress EmailAddress
	Password     string
	Items        []Item
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserDTO struct {
	Uid          string    `json:"uid"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	EmailAddress string    `json:"emailAddress"`
	Password     string    `json:"password"`
	Items        []ItemDTO `json:"items"`
	CreatedAt    string    `json:"createdAt"`
	UpdatedAt    string    `json:"updatedAt"`
}

type UserRaw struct {
	Uid          string    `json:"uid"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	EmailAddress string    `json:"email_address"`
	Password     string    `json:"password"`
	Items        []ItemRaw `json:"items"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
}

func CreateUser(firstName string, lastName string, emailAddress string, password string) User {
	return User{
		Uid:          uuid.New().String(),
		FirstName:    CreateName(FirstName, firstName),
		LastName:     CreateName(LastName, lastName),
		EmailAddress: CreateEmailAddress(emailAddress),
		Password:     password,
		Items:        []Item{},
		CreatedAt:    time.Now(),
	}
}

func (u *User) ToDTO() UserDTO {
	return UserDTO{
		Uid:          u.Uid,
		FirstName:    u.FirstName.Value,
		LastName:     u.LastName.Value,
		EmailAddress: u.EmailAddress.Value,
		Password:     u.Password,
		CreatedAt:    u.CreatedAt.Format("YYYY-MM-DD"),
		UpdatedAt:    u.UpdatedAt.Format("YYYY-MM-DD"),
	}
}

func (r *UserRaw) ToEntity() User {
	createdAt, _ := time.Parse("YYYY-MM-DD", r.CreatedAt)
	updatedAt, _ := time.Parse("YYYY-MM-DD", r.UpdatedAt)
	return User{
		Uid:          r.Uid,
		FirstName:    CreateName(FirstName, r.FirstName),
		LastName:     CreateName(LastName, r.LastName),
		EmailAddress: CreateEmailAddress(r.EmailAddress),
		Password:     r.Password,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
