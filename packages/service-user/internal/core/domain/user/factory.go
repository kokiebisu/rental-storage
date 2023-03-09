package user

import (
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

const TimeLayout = "2006-01-02 15:04:05"

func New(uid string, firstName string, lastName string, emailAddress string, password string, items []item.Entity, createdAtString string, updatedAtString string) (Entity, *customerror.CustomError) {
	var createdAt time.Time
	var updatedAt time.Time
	var err error
	if uid == "" {
		uid = uuid.New().String()
	}
	if createdAtString == "" {
		createdAt = time.Now()
	} else {
		createdAt, err = time.Parse(TimeLayout, createdAtString)
		if err != nil {
			return Entity{}, customerror.ErrorHandler.InvalidValueError("createdAt", "cannot be parsed")
		}
	}
	if updatedAtString == "" {
		updatedAt = time.Time{}
	} else {
		updatedAt, err = time.Parse(TimeLayout, updatedAtString)
		if err != nil {
			return Entity{}, customerror.ErrorHandler.InvalidValueError("updatedAt", "cannot be parsed")
		}
	}
	if err := Name(firstName).validate(); err != nil {
		return Entity{}, err
	}
	if err := Name(lastName).validate(); err != nil {
		return Entity{}, err
	}
	if err := EmailAddress(emailAddress).validate(); err != nil {
		return Entity{}, err
	}
	return Entity{
		UId:          uid,
		FirstName:    Name(firstName),
		LastName:     Name(lastName),
		EmailAddress: EmailAddress(emailAddress),
		Password:     password,
		Items:        items,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}

func (email EmailAddress) validate() *customerror.CustomError {
	_, err := mail.ParseAddress(string(email))
	if err != nil {
		return customerror.ErrorHandler.InvalidValueError("emailAddress", "not following correct format")
	}
	return nil
}

func (name Name) validate() *customerror.CustomError {
	if name == "" {
		return customerror.ErrorHandler.InvalidValueError("name", "should not be empty")
	}
	// edge case: check if it is only consisted of one word
	return nil
}
