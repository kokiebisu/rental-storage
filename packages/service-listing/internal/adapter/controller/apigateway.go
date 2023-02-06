package controller

import (
	"encoding/json"
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

func (h *ApiGatewayHandler) FindListingById(event events.APIGatewayProxyRequest) (FindListingByIdResponsePayload, *customerror.CustomError) {
	uid := event.PathParameters["listingId"]
	if uid == "" {
		return FindListingByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError()
	}
	l, err := h.service.FindListingById(uid)
	return FindListingByIdResponsePayload{Listing: l}, err
}

func (h *ApiGatewayHandler) FindListingsWithinLatLng(event events.APIGatewayProxyRequest) (FindListingsWithinLatLngResponsePayload, *customerror.CustomError) {
	latitude, err := strconv.ParseFloat(event.QueryStringParameters["latitude"], 32)
	if err != nil {
		return FindListingsWithinLatLngResponsePayload{}, customerror.ErrorHandler.ConvertError("latitude", "String", err)
	}
	longitude, err := strconv.ParseFloat(event.QueryStringParameters["longitude"], 32)
	if err != nil {
		return FindListingsWithinLatLngResponsePayload{}, customerror.ErrorHandler.ConvertError("longitude", "String", err)
	}
	distance, err := strconv.ParseInt(event.QueryStringParameters["distance"], 10, 32)
	if err != nil {
		return FindListingsWithinLatLngResponsePayload{}, customerror.ErrorHandler.ConvertError("distance", "String", err)
	}
	listings, err := h.service.FindListingsWithinLatLng(latitude, longitude, int32(distance))
	return FindListingsWithinLatLngResponsePayload{
		Listings: listings,
	}, err.(*customerror.CustomError)
}

func (h *ApiGatewayHandler) AddListing(event events.APIGatewayProxyRequest) (AddListingResponsePayload, *customerror.CustomError) {
	body := struct {
		LenderId      string   `json:"lenderId"`
		StreetAddress string   `json:"streetAddress"`
		Latitude      float64  `json:"latitude"`
		Longitude     float64  `json:"longitude"`
		ImageUrls     []string `json:"imageUrls"`
		Title         string   `json:"title"`
		FeeAmount     int32    `json:"feeAmount"`
		FeeCurrency   string   `json:"feeCurrency"`
		FeeType       string   `json:"feeType"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return AddListingResponsePayload{}, customerror.ErrorHandler.UnmarshalError("listing body", err)
	}
	listingId, err := h.service.CreateListing(body.LenderId, body.StreetAddress, body.Latitude, body.Longitude, body.ImageUrls, body.Title, body.FeeAmount, amount.CurrencyType(body.FeeCurrency), fee.RentalFeeType(body.FeeType))
	return AddListingResponsePayload{
		UId: listingId,
	}, err.(*customerror.CustomError)
}

func (h *ApiGatewayHandler) RemoveListingById(event events.APIGatewayProxyRequest) *customerror.CustomError {
	uid := event.PathParameters["listingId"]
	if uid == "" {
		return customerror.ErrorHandler.GetParameterError("listingId")
	}
	err := h.service.RemoveListingById(uid)
	return err
}
