package booking

import (
	"time"

	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type Entity struct {
	UId       string
	Status    BookingStatus
	UserId    string
	SpaceId   string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DTO struct {
	UId       string `json:"uid"`
	Status    string `json:"status"`
	UserId    string `json:"userId"`
	SpaceId   string `json:"spaceId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Raw struct {
	Id        string `json:"Id"`
	Status    string `json:"Status"`
	UserId    string `json:"UserId"`
	SpaceId   string `json:"SpaceId"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
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

func New(id string, userId string, spaceId string, startDateString string, endDateString string, createdAtString string, updatedAtString string) (Entity, *customerror.CustomError) {
	var startDate time.Time
	var err error
	if startDateString == "" {
		return Entity{}, customerror.ErrorHandler.InternalServerError("provided startDate string cannot be parsed", nil)
	} else {
		startDate, err = time.Parse(layoutISO, startDateString[:10])
		if err != nil {
			return Entity{}, customerror.ErrorHandler.InternalServerError("provided startDate string cannot be parsed", nil)
		}
	}

	var endDate time.Time
	if endDateString == "" {
		return Entity{}, customerror.ErrorHandler.InternalServerError("provided endDate string cannot be parsed", nil)
	} else {
		endDate, err = time.Parse(layoutISO, endDateString[:10])
		if err != nil {
			return Entity{}, customerror.ErrorHandler.InternalServerError("provided endDate string cannot be parsed", nil)
		}
	}

	var createdAt time.Time
	if createdAtString == "" {
		createdAtString = time.Now().Format(layoutISO)
	}
	createdAt, err = time.Parse(layoutISO, createdAtString)
	if err != nil {
		return Entity{}, customerror.ErrorHandler.InternalServerError("provided createdAt string cannot be parsed", nil)
	}

	var updatedAt time.Time
	if updatedAtString == "" {
		updatedAt = time.Time{}
	} else {
		updatedAt, err = time.Parse(layoutISO, updatedAtString)
	}

	if err != nil {
		return Entity{}, customerror.ErrorHandler.InternalServerError("provided updatedAt string cannot be parsed", nil)
	}

	return Entity{
		UId:       id,
		Status:    PENDING,
		UserId:    userId,
		SpaceId:   spaceId,
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (d DTO) ToEntity() Entity {
	startDate, _ := time.Parse(layoutISO, d.StartDate)
	endDate, _ := time.Parse(layoutISO, d.EndDate)
	createdAt, _ := time.Parse(layoutISO, d.CreatedAt)
	updatedAt, _ := time.Parse(layoutISO, d.UpdatedAt)
	return Entity{
		UId:       d.UId,
		Status:    BookingStatus(d.Status),
		UserId:    d.UserId,
		SpaceId:   d.SpaceId,
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (e Entity) ToDTO() DTO {
	startDateString := e.StartDate.Format(layoutISO)
	endDateString := e.EndDate.Format(layoutISO)
	createdAtString := e.CreatedAt.Format(layoutISO)
	updatedAtString := e.UpdatedAt.Format(layoutISO)
	return DTO{
		UId:       e.UId,
		Status:    string(e.Status),
		UserId:    e.UserId,
		SpaceId:   e.SpaceId,
		StartDate: startDateString,
		EndDate:   endDateString,
		CreatedAt: createdAtString,
		UpdatedAt: updatedAtString,
	}
}
