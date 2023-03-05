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

type ApiGatewayAdapter struct {
	service port.BookingService
}

func NewApiGatewayAdapter(service port.BookingService) (port.Controller, *customerror.CustomError) {
	return &ApiGatewayAdapter{
		service,
	}, nil
}

func (a *ApiGatewayAdapter) CreateBooking(event interface{}) (interface{}, *customerror.CustomError) {
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
	bookingId, err := a.service.CreateBooking(uuid.New().String(), body.UserId, body.SpaceId, body.ImageUrls, body.Description, body.StartDate, body.EndDate, "", "")
	return CreateBookingResponsePayload{UId: bookingId}, err.(*customerror.CustomError)
}

func (a *ApiGatewayAdapter) FindBookingById(event interface{}) (interface{}, *customerror.CustomError) {
	bookingId := event.(events.APIGatewayProxyRequest).PathParameters["bookingId"]
	if bookingId == "" {
		return FindBookingByIdResponsePayload{}, customerror.ErrorHandler.InternalServerError("unable to extract bookingId", nil)
	}
	booking, err := a.service.FindById(bookingId)
	return FindBookingByIdResponsePayload{Booking: booking.ToDTO()}, err
}

func (a *ApiGatewayAdapter) FindBookings(event interface{}) (interface{}, *customerror.CustomError) {
	spaceId := event.(events.APIGatewayProxyRequest).QueryStringParameters["spaceId"]
	if spaceId == "" {
		return FindBookingsResponsePayload{}, customerror.ErrorHandler.InternalServerError("unable to extract bookingId", nil)
	}
	b, err := a.service.FindManyBySpaceId(spaceId)
	bs := []booking.DTO{}
	for _, i := range b {
		bs = append(bs, i.ToDTO())
	}
	return FindBookingsResponsePayload{Bookings: bs}, err
}
