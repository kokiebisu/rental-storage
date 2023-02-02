package controller

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ApiGatewayHandler struct {
	service port.ListingService
}

type FindListingByIdResponsePayload struct {
	Listing listing.DTO `json:"listing"`
}

type FindListingsWithinLatLngResponsePayload struct {
	Listings []listing.DTO `json:"listings"`
}

type AddListingResponsePayload struct {
	UId string `json:"uid"`
}

func (h *ApiGatewayHandler) FindListingById(event events.APIGatewayProxyRequest) (listing.DTO, *customerror.CustomError) {
	uid := event.PathParameters["listingId"]
	if uid == "" {
		return listing.DTO{}, customerror.ErrorHandler.InternalServerError(errors.New("listingId not available"))
	}
	listing, err := h.service.FindListingById(uid)
	return listing, err
}

func (h *ApiGatewayHandler) FindListingsWithinLatLng(event events.APIGatewayProxyRequest) (FindListingsWithinLatLngResponsePayload, *customerror.CustomError) {
	latitude, err := strconv.ParseFloat(event.QueryStringParameters["latitude"], 32)
	if err != nil {
		return FindListingsWithinLatLngResponsePayload{}, customerror.ErrorHandler.InternalServerError(errors.New("unable to parse latitude"))
	}
	longitude, err := strconv.ParseFloat(event.QueryStringParameters["longitude"], 32)
	if err != nil {
		return FindListingsWithinLatLngResponsePayload{}, customerror.ErrorHandler.InternalServerError(errors.New("unable to parse longitude"))
	}
	distance, err := strconv.ParseInt(event.QueryStringParameters["distance"], 10, 32)
	if err != nil {
		return FindListingsWithinLatLngResponsePayload{}, customerror.ErrorHandler.InternalServerError(errors.New("unable to parse distance"))
	}
	listings, err := h.service.FindListingsWithinLatLng(float32(latitude), float32(longitude), int32(distance))
	return FindListingsWithinLatLngResponsePayload{
		Listings: listings,
	}, err.(*customerror.CustomError)
}

func (h *ApiGatewayHandler) AddListing(event events.APIGatewayProxyRequest) (AddListingResponsePayload, *customerror.CustomError) {
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
		return AddListingResponsePayload{}, customerror.ErrorHandler.InternalServerError(errors.New("unable to unmarshal"))
	}
	listingId, err := h.service.CreateListing(body.LenderId, body.StreetAddress, body.Latitude, body.Longitude, body.ImageUrls, body.Title, body.FeeAmount, amount.CurrencyType(body.FeeCurrency), fee.RentalFeeType(body.FeeType))
	return AddListingResponsePayload{
		UId: listingId,
	}, err.(*customerror.CustomError)
}
