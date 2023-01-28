package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type ApiGatewayHandler struct {
	service port.BookingService
}

func NewApiGatewayHandler(service port.BookingService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) CreateBooking(event events.APIGatewayProxyRequest) (string, *errors.CustomError) {
	body := struct {
		Amount    amount.DTO `json:"amount"`
		UserId    string     `json:"userId"`
		ListingId string     `json:"listingId"`
		Items     []item.DTO `json:"items"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return "", errors.ErrorHandler.InternalServerError()
	}
	bookingId, err := h.service.CreateBooking(body.Amount, body.UserId, body.ListingId, body.Items)
	return bookingId, err.(*errors.CustomError)
}

func (h *ApiGatewayHandler) FindUserBookings(event events.APIGatewayProxyRequest) ([]booking.DTO, *errors.CustomError) {
	userId := event.QueryStringParameters["userId"]
	if userId == "" {
		return []booking.DTO{}, errors.ErrorHandler.InternalServerError()
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
