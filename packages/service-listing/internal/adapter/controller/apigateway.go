package controller

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	domain "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/port"
)

type ApiGatewayHandler struct {
	service port.ListingService
}

func NewApiGatewayHandler(service port.ListingService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) FindListingById(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	uid := event.PathParameters["listingId"]
	if uid == "" {
		return sendFailureResponse(errors.New("listing Id not found"))
	}
	listing, err := h.service.FindListingById(uid)
	if err != nil {
		return sendFailureResponse(err)
	}
	return sendResponse(listing)
}

func (h *ApiGatewayHandler) FindListingsWithinLatLng(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	latitude, err := strconv.ParseFloat(event.QueryStringParameters["latitude"], 32)
	if err != nil {
		return sendFailureResponse(err)
	}
	longitude, err := strconv.ParseFloat(event.QueryStringParameters["longitude"], 32)
	if err != nil {
		return sendFailureResponse(err)
	}
	distance, err := strconv.ParseInt(event.QueryStringParameters["distance"], 10, 32)
	if err != nil {
		return sendFailureResponse(err)
	}
	listings, err := h.service.FindListingsWithinLatLng(float32(latitude), float32(longitude), int32(distance))
	if err != nil {
		return sendFailureResponse(err)
	}
	return sendResponse(listings)
}

func (h *ApiGatewayHandler) AddListing(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := struct {
		LenderId      string   `json:"lenderId"`
		StreetAddress string   `json:"streetAddress"`
		Latitude      float32  `json:"latitude"`
		Longitude     float32  `json:"longitude"`
		ImageUrls     []string `json:"imageUrls"`
		Title         string   `json:"title"`
		FeeAmount     int32    `json:"feeAmount"`
		FeeCurrency   string   `json:"feeCurrency"`
		FeeType       string   `json:"feeType"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		sendFailureResponse(err)
	}
	listingId, err := h.service.CreateListing(body.LenderId, body.StreetAddress, body.Latitude, body.Longitude, body.ImageUrls, body.Title, body.FeeAmount, domain.CurrencyType(body.FeeCurrency), domain.RentalFeeType(body.FeeType))
	if err != nil {
		return sendFailureResponse(err)
	}
	return sendCreatedResponse(listingId)
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

func sendDeletedResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 204,
	}, nil
}

func sendCreatedResponse(listingId string) (events.APIGatewayProxyResponse, error) {
	encoded, err := json.Marshal(&struct {
		Uid string `json:"uid"`
	}{
		Uid: listingId,
	})
	if err != nil {
		panic(err)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(encoded),
	}, nil
}

func sendFailureResponse(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       string(err.Error()),
	}, nil
}
