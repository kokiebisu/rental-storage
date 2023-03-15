package adapter

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/kokiebisu/rental-storage/service-space/internal/client"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/location"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type FindSpaceByIdResponsePayload struct {
	Space space.DTO `json:"space"`
}

type FindSpacesResponsePayload struct {
	Spaces []space.DTO `json:"spaces"`
}

type AddSpaceResponsePayload struct {
	UId string `json:"uid"`
}

type DeleteSpaceByIdResponsePayload struct {
	UId string `json:"uid"`
}

type ApiGatewayAdapter struct {
	service port.SpaceService
}

func NewApiGatewayAdapter(service port.SpaceService) port.Controller {
	return &ApiGatewayAdapter{
		service,
	}
}

func (h *ApiGatewayAdapter) FindSpaceById(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
	spaceId := event.(events.APIGatewayProxyRequest).PathParameters["spaceId"]
	if spaceId == "" {
		return FindSpaceByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("spaceId")
	}
	l, err := h.service.FindSpaceById(spaceId)
	payload := FindSpaceByIdResponsePayload{Space: l}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, err
}

func (h *ApiGatewayAdapter) FindSpaces(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
	userId := event.(events.APIGatewayProxyRequest).QueryStringParameters["userId"]
	if userId != "" {
		ls, err := h.service.FindSpacesByUserId(userId)
		payload := FindSpacesResponsePayload{Spaces: ls}
		logger.Info("Payload", zap.Any("payload", payload))
		return payload, err
	}
	offset, err := strconv.Atoi(event.(events.APIGatewayProxyRequest).QueryStringParameters["offset"])
	if err != nil {
		return FindSpacesResponsePayload{}, customerror.ErrorHandler.ConvertError("offset", "String", err)
	}
	limit, err := strconv.Atoi(event.(events.APIGatewayProxyRequest).QueryStringParameters["limit"])
	if err != nil {
		return FindSpacesResponsePayload{}, customerror.ErrorHandler.ConvertError("limit", "String", err)
	}
	if offset >= 0 && limit >= 0 {
		ls, err := h.service.FindSpaces(offset, limit)
		if err != nil {
			return FindSpacesResponsePayload{}, err
		}
		payload := FindSpacesResponsePayload{Spaces: ls}
		logger.Info("Payload", zap.Any("payload", payload))
		return payload, nil
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
	return FindSpacesResponsePayload{}, customerror.ErrorHandler.OffsetLimitNotProvidedError(nil)
}

func (h *ApiGatewayAdapter) AddSpace(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
	body := struct {
		LenderId    string       `json:"lenderId"`
		Location    location.DTO `json:"location"`
		ImageUrls   []string     `json:"imageUrls"`
		Title       string       `json:"title"`
		Description string       `json:"description"`
	}{}
	err := json.Unmarshal([]byte(event.(events.APIGatewayProxyRequest).Body), &body)
	if err != nil {
		logger.Error("Error", zap.Error(err))
		return AddSpaceResponsePayload{}, customerror.ErrorHandler.UnmarshalError("space body", err)
	}
	spaceId, err := h.service.CreateSpace(uuid.New().String(), body.LenderId, body.Location, body.ImageUrls, body.Title, body.Description, "", "")
	payload := AddSpaceResponsePayload{
		UId: spaceId,
	}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, err.(*customerror.CustomError)
}

func (h *ApiGatewayAdapter) DeleteSpaceById(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	logger.Info("Event", zap.Any("event", event))
	spaceId := event.(events.APIGatewayProxyRequest).PathParameters["spaceId"]
	if spaceId == "" {
		logger.Error("Error", zap.Error(customerror.ErrorHandler.GetParameterError("spaceId")))
		return DeleteSpaceByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("spaceId")
	}
	uid, err := h.service.DeleteSpaceById(spaceId)
	payload := DeleteSpaceByIdResponsePayload{UId: uid}
	logger.Info("Payload", zap.Any("payload", payload))
	return payload, err
}
