package adapter

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type CreateBookingResponsePayload struct {
	UId string `json:"uid"`
}

type FindBookingByIdResponsePayload struct {
	Booking booking.DTO `json:"booking"`
}

type FindBookingsResponsePayload struct {
	Bookings []booking.DTO `json:"bookings"`
}

type ControllerAdapter struct {
	service port.BookingService
}

func NewControllerAdapter(service port.BookingService) (port.Controller, *customerror.CustomError) {
	return NewApiGatewayAdapter(service)
}

func NewApiGatewayAdapter(service port.BookingService) (port.Controller, *customerror.CustomError) {
	// db, err := GetDBAdapter()
	// if err != nil {
	// 	return nil, err
	// }
	// repo := repository.NewUserRepository(db)
	// pa, err := GetPublisherAdapter()
	// if err != nil {
	// 	return nil, err
	// }
	// publisher := publisher.NewUserPublisher(pa)
	// err = repo.Setup()

	// service := service.NewUserService(repo, publisher)
	// return &ApiGatewayAdapter{
	// 	service,
	// }, err
	return &ControllerAdapter{
		service,
	}, nil
}

func (h *ControllerAdapter) CreateBooking(event interface{}) (interface{}, *customerror.CustomError) {
	body := struct {
		UserId      string   `json:"userId"`
		SpaceId     string   `json:"spaceId"`
		ImageUrls   []string `json:"imageUrls"`
		Description string   `json:"description"`
		StartDate   string   `json:"startDate"`
		EndDate     string   `json:"endDate"`
	}{}
	err := json.Unmarshal([]byte(event.(events.APIGatewayProxyRequest).Body), &body)
	if err != nil {
		return CreateBookingResponsePayload{}, customerror.ErrorHandler.InternalServerError("unable to unmarshal body request", err)
	}
	bookingId, err := h.service.CreateBooking(uuid.New().String(), body.UserId, body.SpaceId, body.ImageUrls, body.Description, body.StartDate, body.EndDate, "", "")
	return CreateBookingResponsePayload{UId: bookingId}, err.(*customerror.CustomError)
}

func (h *ControllerAdapter) FindBookingById(event interface{}) (interface{}, *customerror.CustomError) {
	bookingId := event.(events.APIGatewayProxyRequest).PathParameters["bookingId"]
	if bookingId == "" {
		return FindBookingByIdResponsePayload{}, customerror.ErrorHandler.InternalServerError("unable to extract bookingId", nil)
	}
	booking, err := h.service.FindById(bookingId)
	return FindBookingByIdResponsePayload{Booking: booking.ToDTO()}, err
}

func (h *ControllerAdapter) FindBookings(event interface{}) (interface{}, *customerror.CustomError) {
	spaceId := event.(events.APIGatewayProxyRequest).QueryStringParameters["spaceId"]
	if spaceId == "" {
		return FindBookingsResponsePayload{}, customerror.ErrorHandler.InternalServerError("unable to extract bookingId", nil)
	}
	b, err := h.service.FindManyBySpaceId(spaceId)
	bs := []booking.DTO{}
	for _, i := range b {
		bs = append(bs, i.ToDTO())
	}
	return FindBookingsResponsePayload{Bookings: bs}, err
}
