package booking

import (
	"time"

	"github.com/google/uuid"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type Entity struct {
	UId       string
	Status    BookingStatus
	UserId    string
	SpaceId   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DTO struct {
	UId       string `json:"uid"`
	Status    string `json:"status"`
	UserId    string `json:"userId"`
	SpaceId   string `json:"spaceId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Raw struct {
	Id        string `json:"Id"`
	Status    string `json:"Status"`
	UserId    string `json:"UserId"`
	SpaceId   string `json:"SpaceId"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
}

const (
	PENDING   BookingStatus = "PENDING"
	COMPLETED BookingStatus = "COMPLETED"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

type BookingStatus string

func New(id string, userId string, spaceId string, createdAtString string, updatedAtString string) (Entity, *customerror.CustomError) {
	if id == "" {
		id = uuid.New().String()
	}
	if createdAtString == "" {
		createdAtString = time.Now().Format(layoutISO)
	}
	createdAt, err := time.Parse(layoutISO, createdAtString)
	if err != nil {
		return Entity{}, customerror.ErrorHandler.InternalServerError("provided createdAt string cannot be parsed", nil)
	}
	var updatedAt time.Time
	if updatedAtString == "" {
		updatedAt = time.Time{}
	} else {
		updatedAt, err = time.Parse(layoutISO, updatedAtString)
		if err != nil {
			return Entity{}, customerror.ErrorHandler.InternalServerError("provided updatedAt string cannot be parsed", nil)
		}
	}
	return Entity{
		UId:       id,
		Status:    PENDING,
		UserId:    userId,
		SpaceId:   spaceId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (d DTO) ToEntity() Entity {
	createdAt, _ := time.Parse(layoutISO, d.CreatedAt)
	updatedAt, _ := time.Parse(layoutISO, d.UpdatedAt)
	return Entity{
		UId:       d.UId,
		Status:    BookingStatus(d.Status),
		UserId:    d.UserId,
		SpaceId:   d.SpaceId,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (e Entity) ToDTO() DTO {
	createdAtString := e.CreatedAt.Format("2016-08-19")
	updatedAtString := e.UpdatedAt.Format("2016-08-19")
	return DTO{
		UId:       e.UId,
		Status:    string(e.Status),
		UserId:    e.UserId,
		SpaceId:   e.SpaceId,
		CreatedAt: createdAtString,
		UpdatedAt: updatedAtString,
	}
}
