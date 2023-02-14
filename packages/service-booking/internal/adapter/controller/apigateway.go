package controller

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type ApiGatewayHandler struct {
	service port.BookingService
}

type CreateBookingResponsePayload struct {
	UId string `json:"uid"`
}

type FindBookingByIdResponsePayload struct {
	Booking booking.DTO `json:"booking"`
}

type FindBookingsResponsePayload struct {
	Bookings []booking.DTO `json:"bookings"`
}

func NewApiGatewayHandler(service port.BookingService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) CreateBooking(event events.APIGatewayProxyRequest) (CreateBookingResponsePayload, *customerror.CustomError) {
	body := struct {
		UserId    string   `json:"userId"`
		SpaceId   string   `json:"spaceId"`
		ImageUrls []string `json:"imageUrls"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return CreateBookingResponsePayload{}, customerror.ErrorHandler.InternalServerError("unable to unmarshal body request", err)
	}
	fmt.Println("ENTERED1: ", body.ImageUrls[0], body.ImageUrls[1])
	bookingId, err := h.service.CreateBooking("", body.UserId, body.SpaceId, body.ImageUrls, "", "")
	return CreateBookingResponsePayload{UId: bookingId}, err.(*customerror.CustomError)
}

func (h *ApiGatewayHandler) FindBookingById(event events.APIGatewayProxyRequest) (FindBookingByIdResponsePayload, *customerror.CustomError) {
	bookingId := event.PathParameters["bookingId"]
	if bookingId == "" {
		return FindBookingByIdResponsePayload{}, customerror.ErrorHandler.InternalServerError("unable to extract bookingId", nil)
	}
	booking, err := h.service.FindById(bookingId)
	return FindBookingByIdResponsePayload{Booking: booking.ToDTO()}, err
}

func (h *ApiGatewayHandler) FindBookings(event events.APIGatewayProxyRequest) (FindBookingsResponsePayload, *customerror.CustomError) {
	spaceId := event.QueryStringParameters["spaceId"]
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
