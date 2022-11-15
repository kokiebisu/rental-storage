package receiver

import "github.com/kokiebisu/rental-storage/service-user/internal/core/port"

func New(service port.UserService) *DynamoDBStreamHandler {
	return NewDynamoDBStreamHandler(service)
}
