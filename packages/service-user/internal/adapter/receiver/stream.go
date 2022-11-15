package receiver

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
)

type DynamoDBStreamHandler struct {
	service port.UserService
}

func NewDynamoDBStreamHandler(service port.UserService) *DynamoDBStreamHandler {
	return &DynamoDBStreamHandler{
		service: service,
	}
}

func (h *DynamoDBStreamHandler) AddItem(event events.DynamoDBEvent) {
	for _, r := range event.Records {
		fmt.Println("Record: ", r)
	}
}
