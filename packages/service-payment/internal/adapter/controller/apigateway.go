package controller

import (
	"encoding/json"
	"service-payment/internal/core/port"

	"github.com/aws/aws-lambda-go/events"
)

type ApiGatewayHandler struct {
	customerService port.CustomerService
	purchaseService port.PurchaseService
}

type Event struct {
	Field     string      `json:"field"`
	Arguments interface{} `json:"arguments"`
	Identity  Identity    `json:"identity"`
}

type Identity struct {
	ResolverContext ResolverContext
}

type ResolverContext struct {
	UId string `json:"uid"`
}

type ResponsePayload struct {
	ProviderId   string `json:"providerId"`
	ProviderType string `json:"providerType"`
}

type Error struct {
	Message string `json:"message"`
}

func NewApiGatewayHandler(customerService port.CustomerService, purchaseService port.PurchaseService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		customerService: customerService,
		purchaseService: purchaseService,
	}
}

func (h *ApiGatewayHandler) MakePayment(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// fmt.Println("EVENT: ", event)
	// fmt.Println("EVENT Identity: ", event.Identity)
	// fmt.Println("EVENT Resolver conext: ", event.Identity.ResolverContext)
	// fmt.Println("EVENT Resolver conext uid: ", event.Identity.ResolverContext.UId)

	// endpoint := os.Getenv("SERVICE_API_ENDPOINT")
	// fmt.Println("ENDPOINT: ", endpoint)
	return sendSuccessResponse()
}

func (h *ApiGatewayHandler) CreateCustomer(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := struct {
		UserId       string `json:"userId"`
		EmailAddress string `json:"emailAddress"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
	}{}
	json.Unmarshal([]byte(event.Body), &body)
	err := h.customerService.CreateCustomer(body.UserId, body.FirstName, body.LastName, body.EmailAddress)
	if err != nil {
		return sendFailureResponse(err)
	}
	return sendCreatedResponse()
}

func sendSuccessResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func sendCreatedResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
	}, nil
}

func sendFailureResponse(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       string(err.Error()),
	}, nil
}
