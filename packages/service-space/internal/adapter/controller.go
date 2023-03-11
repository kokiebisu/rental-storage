package adapter

import (
	"encoding/json"

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

type ControllerAdapter struct {
	service port.SpaceService
}

func NewControllerAdapter(service port.SpaceService) (port.Controller, *customerror.CustomError) {
	return NewApiGatewayAdapter(service)
}

func NewApiGatewayAdapter(service port.SpaceService) (port.Controller, *customerror.CustomError) {
	return &ControllerAdapter{
		service,
	}, nil
}

func (h *ControllerAdapter) FindSpaceById(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Error("Error syncing logger", zap.Error(err))
		}
	}()
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

func (h *ControllerAdapter) FindSpaces(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Error("Error syncing logger", zap.Error(err))
		}
	}()
	logger.Info("Event", zap.Any("event", event))
	userId := event.(events.APIGatewayProxyRequest).QueryStringParameters["userId"]
	if userId != "" {
		ls, err := h.service.FindSpacesByUserId(userId)
		logger.Info("Payload", zap.Any("payload", ls))
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
	payload := FindSpacesResponsePayload{}
	logger.Error("Error", zap.Error(customerror.ErrorHandler.InvalidParamError(nil)))
	return payload, customerror.ErrorHandler.InvalidParamError(nil)
}

func (h *ControllerAdapter) AddSpace(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Error("Error syncing logger", zap.Error(err))
		}
	}()
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

func (h *ControllerAdapter) DeleteSpaceById(event interface{}) (interface{}, *customerror.CustomError) {
	logger, _ := client.GetLoggerClient()
	defer func() {
		err := logger.Sync()
		if err != nil {
			logger.Error("Error syncing logger", zap.Error(err))
		}
	}()
	logger.Info("Event", zap.Any("event", event))
	spaceId := event.(events.APIGatewayProxyRequest).PathParameters["spaceId"]
	if spaceId == "" {
		logger.Error("Error", zap.Error(customerror.ErrorHandler.GetParameterError("spaceId")))
		return DeleteSpaceByIdResponsePayload{}, customerror.ErrorHandler.GetParameterError("spaceId")
	}
	uid, err := h.service.DeleteSpaceById(spaceId)
	return DeleteSpaceByIdResponsePayload{UId: uid}, err
}
