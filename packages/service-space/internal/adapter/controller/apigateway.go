package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/location"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type ApiGatewayHandler struct {
	service port.SpaceService
}

type FindSpaceByIdResponsePayload struct {
	Space space.DTO `json:"space"`
}

type FindSpacesResponsePayload struct {
	Spaces []space.DTO `json:"spaces"`
}

type AddSpaceResponsePayload struct {
	UId string `json:"uid"`
}

type RemoveSpaceByIdResponsePayload struct {
	UId string `json:"uid"`
}

func (h *ApiGatewayHandler) FindSpaceById(event events.APIGatewayProxyRequest) (FindSpaceByIdResponsePayload, *customerror.CustomError) {
	spaceId := event.PathParameters["spaceId"]
	if spaceId == "" {
		return FindSpaceByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("spaceId")
	}
	l, err := h.service.FindSpaceById(spaceId)
	return FindSpaceByIdResponsePayload{Space: l}, err
}

func (h *ApiGatewayHandler) FindSpaces(event events.APIGatewayProxyRequest) (FindSpacesResponsePayload, *customerror.CustomError) {
	userId := event.QueryStringParameters["userId"]
	if userId != "" {
		ls, err := h.service.FindSpacesByUserId(userId)
		return FindSpacesResponsePayload{Spaces: ls}, err
	}
	// @deprecated
	// latitudeString := event.QueryStringParameters["lat"]
	// longitudeString := event.QueryStringParameters["lng"]
	// distanceString := event.QueryStringParameters["range"]

	// if latitudeString != "" && longitudeString != "" && distanceString != "" {

	// 	latitude, err := strconv.ParseFloat(latitudeString, 32)
	// 	if err != nil {
	// 		return FindSpacesResponsePayload{}, customerror.ErrorHandler.ConvertError("latitude", "String", err)
	// 	}
	// 	longitude, err := strconv.ParseFloat(longitudeString, 32)
	// 	if err != nil {
	// 		return FindSpacesResponsePayload{}, customerror.ErrorHandler.ConvertError("longitude", "String", err)
	// 	}
	// 	distance, err := strconv.ParseInt(distanceString, 10, 32)
	// 	if err != nil {
	// 		return FindSpacesResponsePayload{}, customerror.ErrorHandler.ConvertError("distance", "String", err)
	// 	}
	// 	spaces, err := h.service.FindSpacesWithinLatLng(latitude, longitude, int32(distance))
	// 	return FindSpacesResponsePayload{
	// 		Spaces: spaces,
	// 	}, err.(*customerror.CustomError)
	// }
	return FindSpacesResponsePayload{}, customerror.ErrorHandler.InvalidParamError(nil)
}

func (h *ApiGatewayHandler) AddSpace(event events.APIGatewayProxyRequest) (AddSpaceResponsePayload, *customerror.CustomError) {
	body := struct {
		LenderId    string       `json:"lenderId"`
		Location    location.DTO `json:"location"`
		ImageUrls   []string     `json:"imageUrls"`
		Title       string       `json:"title"`
		Description string       `json:"description"`
	}{}
	err := json.Unmarshal([]byte(event.Body), &body)
	if err != nil {
		return AddSpaceResponsePayload{}, customerror.ErrorHandler.UnmarshalError("space body", err)
	}
	spaceId, err := h.service.CreateSpace(uuid.New().String(), body.LenderId, body.Location, body.ImageUrls, body.Title, body.Description, "", "")
	return AddSpaceResponsePayload{
		UId: spaceId,
	}, err.(*customerror.CustomError)
}

func (h *ApiGatewayHandler) RemoveSpaceById(event events.APIGatewayProxyRequest) (RemoveSpaceByIdResponsePayload, *customerror.CustomError) {
	spaceId := event.PathParameters["spaceId"]
	if spaceId == "" {
		return RemoveSpaceByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("spaceId")
	}
	uid, err := h.service.RemoveSpaceById(spaceId)
	return RemoveSpaceByIdResponsePayload{UId: uid}, err
}
