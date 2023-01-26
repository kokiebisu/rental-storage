package controller

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/amount"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
)

type ApiGatewayHandler struct {
	service port.BookingService
}

func NewApiGatewayHandler(service port.BookingService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) CreateBooking(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := struct {
		Amount    amount.DTO `json:"amount"`
		UserId    string     `json:"userId"`
		ListingId string     `json:"listingId"`
		Items     []item.DTO `json:"items"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return sendFailureResponse(err)
	}
	bookingId, err := h.service.CreateBooking(body.Amount, body.UserId, body.ListingId, body.Items)
	if err != nil {
		return sendFailureResponse(err)
	}
	return sendCreatedResponse(bookingId)
}

func (h *ApiGatewayHandler) FindUserBookings(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	userId := event.QueryStringParameters["userId"]
	if userId == "" {
		return sendFailureResponse(errors.New("userId not provided"))
	}
	bookings, err := h.service.FindUserBookings(userId)
	if err != nil {
		return sendFailureResponse(err)
	}
	bookingDTOs := []booking.DTO{}
	for _, b := range bookings {
		bookingDTOs = append(bookingDTOs, b.ToDTO())
	}
	return sendResponse(bookingDTOs)
}

func sendFailureResponse(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       string(err.Error()),
	}, nil
}

func sendResponse(data interface{}) (events.APIGatewayProxyResponse, error) {
	encoded, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{
		Body:       string(encoded),
		StatusCode: 200,
	}, nil
}

// func sendDeletedResponse() (events.APIGatewayProxyResponse, error) {
// 	return events.APIGatewayProxyResponse{
// 		StatusCode: 204,
// 	}, nil
// }

func sendCreatedResponse(bookingId string) (events.APIGatewayProxyResponse, error) {
	encoded, err := json.Marshal(&struct {
		Uid string `json:"uid"`
	}{
		Uid: bookingId,
	})
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(encoded),
	}, nil
}
