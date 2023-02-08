package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
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

func NewApiGatewayHandler(service port.BookingService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) CreateBooking(event events.APIGatewayProxyRequest) (CreateBookingResponsePayload, *customerror.CustomError) {
	body := struct {
		Amount    amount.DTO `json:"amount"`
		UserId    string     `json:"userId"`
		ListingId string     `json:"listingId"`
		Items     []item.DTO `json:"items"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return CreateBookingResponsePayload{}, customerror.ErrorHandler.InternalServerError("unable to unmarshal body request", err)
	}
	bookingId, err := h.service.CreateBooking("", body.Amount, body.UserId, body.ListingId, body.Items, "", "")
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

func (h *ApiGatewayHandler) FindUserBookings(event events.APIGatewayProxyRequest) ([]booking.DTO, *customerror.CustomError) {
	userId := event.QueryStringParameters["userId"]
	if userId == "" {
		return []booking.DTO{}, customerror.ErrorHandler.InternalServerError("unable to extract userId", nil)
	}
	bookings, err := h.service.FindUserBookings(userId)
	if err != nil {
		return []booking.DTO{}, err
	}
	bookingDTOs := []booking.DTO{}
	for _, b := range bookings {
		bookingDTOs = append(bookingDTOs, b.ToDTO())
	}
	return bookingDTOs, nil
}
