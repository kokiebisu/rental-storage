package controller

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	domain "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ApiGatewayHandler struct {
	service port.ListingService
}

func NewApiGatewayHandler(service port.ListingService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		service: service,
	}
}

func (h *ApiGatewayHandler) FindListingById(event events.APIGatewayProxyRequest) (domain.ListingDTO, *errors.CustomError) {
	uid := event.PathParameters["listingId"]
	if uid == "" {
		return domain.ListingDTO{}, errors.ErrorHandler.GetParameterError()
	}
	listing, err := h.service.FindListingById(uid)
	return listing, err
}

func (h *ApiGatewayHandler) FindListingsWithinLatLng(event events.APIGatewayProxyRequest) ([]domain.ListingDTO, *errors.CustomError) {
	latitude, err := strconv.ParseFloat(event.QueryStringParameters["latitude"], 32)
	if err != nil {
		return []domain.ListingDTO{}, errors.ErrorHandler.ConvertError("latitude", "String", err)
	}
	longitude, err := strconv.ParseFloat(event.QueryStringParameters["longitude"], 32)
	if err != nil {
		return []domain.ListingDTO{}, errors.ErrorHandler.ConvertError("longitude", "String", err)
	}
	distance, err := strconv.ParseInt(event.QueryStringParameters["distance"], 10, 32)
	if err != nil {
		return []domain.ListingDTO{}, errors.ErrorHandler.ConvertError("distance", "String", err)
	}
	listings, err := h.service.FindListingsWithinLatLng(float32(latitude), float32(longitude), int32(distance))
	return listings, err.(*errors.CustomError)
}

func (h *ApiGatewayHandler) AddListing(event events.APIGatewayProxyRequest) (string, *errors.CustomError) {
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
		return "", errors.ErrorHandler.UnmarshalError("listing body", err)
	}
	listingId, err := h.service.CreateListing(body.LenderId, body.StreetAddress, body.Latitude, body.Longitude, body.ImageUrls, body.Title, body.FeeAmount, domain.CurrencyType(body.FeeCurrency), domain.RentalFeeType(body.FeeType))
	return listingId, err.(*errors.CustomError)
}
