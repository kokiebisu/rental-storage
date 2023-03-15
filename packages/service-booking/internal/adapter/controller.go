package adapter

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/kokiebisu/rental-storage/service-booking/internal/client"
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

func NewApiGatewayAdapter(service port.BookingService) port.Controller {
	return &ApiGatewayAdapter{
		service,
	}
}

func (a *ApiGatewayAdapter) CreateBooking(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
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
	bookingId, err := a.service.CreateBooking(uuid.New().String(), body.UserId, body.SpaceId, body.ImageUrls, "pending", body.Description, body.StartDate, body.EndDate, "", "")
	payload := CreateBookingResponsePayload{UId: bookingId}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, err.(*customerror.CustomError)
}

func (a *ApiGatewayAdapter) FindBookingById(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
	bookingId := event.(events.APIGatewayProxyRequest).PathParameters["bookingId"]
	if bookingId == "" {
		return FindBookingByIdResponsePayload{}, customerror.ErrorHandler.InternalServerError("unable to extract bookingId", nil)
	}
	booking, err := a.service.FindBookingById(bookingId)
	payload := FindBookingByIdResponsePayload{Booking: booking.ToDTO()}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, err
}

func (a *ApiGatewayAdapter) FindBookings(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
	var bs []booking.Entity = []booking.Entity{}
	var err *customerror.CustomError
	bookingStatus := event.(events.APIGatewayProxyRequest).QueryStringParameters["bookingStatus"]

	spaceId := event.(events.APIGatewayProxyRequest).QueryStringParameters["spaceId"]
	if spaceId != "" {
		bs, err = a.service.FindBookingsBySpaceId(spaceId, bookingStatus)
	}

	userId := event.(events.APIGatewayProxyRequest).QueryStringParameters["userId"]
	if userId != "" {
		bs, err = a.service.FindBookingsByUserId(userId, bookingStatus)
	}
	bds := []booking.DTO{}
	for _, i := range bs {
		bds = append(bds, i.ToDTO())
	}
	payload := FindBookingsResponsePayload{Bookings: bds}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, err
}
