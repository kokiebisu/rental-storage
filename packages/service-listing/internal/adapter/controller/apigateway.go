package controller

import (
	"encoding/json"
	"fmt"
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

type FindListingsResponsePayload struct {
	Listings []listing.DTO `json:"listings"`
}

type AddListingResponsePayload struct {
	UId string `json:"uid"`
}

type RemoveListingByIdResponsePayload struct {
	UId string `json:"uid"`
}

func (h *ApiGatewayHandler) FindListingById(event events.APIGatewayProxyRequest) (FindListingByIdResponsePayload, *customerror.CustomError) {
	listingId := event.PathParameters["listingId"]
	if listingId == "" {
		return FindListingByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("listingId")
	}
	l, err := h.service.FindListingById(listingId)
	return FindListingByIdResponsePayload{Listing: l}, err
}

func (h *ApiGatewayHandler) FindListings(event events.APIGatewayProxyRequest) (FindListingsResponsePayload, *customerror.CustomError) {
	userId := event.QueryStringParameters["userId"]
	if userId != "" {
		fmt.Println("FIND LISTINGS BY USER ID")
		ls, err := h.service.FindListingsByUserId(userId)
		return FindListingsResponsePayload{Listings: ls}, err
	}
	latitudeString := event.QueryStringParameters["lat"]
	longitudeString := event.QueryStringParameters["lng"]
	distanceString := event.QueryStringParameters["range"]

	if latitudeString != "" && longitudeString != "" && distanceString != "" {

		latitude, err := strconv.ParseFloat(latitudeString, 32)
		if err != nil {
			return FindListingsResponsePayload{}, customerror.ErrorHandler.ConvertError("latitude", "String", err)
		}
		longitude, err := strconv.ParseFloat(longitudeString, 32)
		if err != nil {
			return FindListingsResponsePayload{}, customerror.ErrorHandler.ConvertError("longitude", "String", err)
		}
		distance, err := strconv.ParseInt(distanceString, 10, 32)
		if err != nil {
			return FindListingsResponsePayload{}, customerror.ErrorHandler.ConvertError("distance", "String", err)
		}
		fmt.Println("FindListingsWithinLatLng ", latitude, longitude, distance)
		listings, err := h.service.FindListingsWithinLatLng(latitude, longitude, int32(distance))
		return FindListingsResponsePayload{
			Listings: listings,
		}, err.(*customerror.CustomError)
	}
	return FindListingsResponsePayload{}, customerror.ErrorHandler.InvalidParamError(nil)
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

func (h *ApiGatewayHandler) RemoveListingById(event events.APIGatewayProxyRequest) (RemoveListingByIdResponsePayload, *customerror.CustomError) {
	listingId := event.PathParameters["listingId"]
	if listingId == "" {
		return RemoveListingByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("listingId")
	}
	uid, err := h.service.RemoveListingById(listingId)
	return RemoveListingByIdResponsePayload{UId: uid}, err
}
