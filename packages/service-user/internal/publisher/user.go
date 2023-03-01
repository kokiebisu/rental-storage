package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/google/uuid"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type UserPublisher struct {
	client *kinesis.Client
}

func NewUserPublisher(client *kinesis.Client) *UserPublisher {
	return &UserPublisher{
		client,
	}
}

func (s *UserPublisher) UserCreated(u user.DTO) *customerror.CustomError {
	event := map[string]interface{}{
		"sourceEntity": "User",
		"eventName":    "created",
		"data":         u,
	}
	encoded, err := json.Marshal(event)
	if err != nil {
		return customerror.ErrorHandler.MarshalError(err)
	}
	return s.publish(encoded)
}

func (s *UserPublisher) publish(data []byte) *customerror.CustomError {
	var environment string
	if os.Getenv("STAGE") == "" {
		environment = "local"
	} else {
		environment = os.Getenv("STAGE")
	}

	streamName := fmt.Sprintf("%s-EventStream", environment)
	partitionKey := uuid.New().String()

	_, err := s.client.PutRecord(context.TODO(), &kinesis.PutRecordInput{
		StreamName:   aws.String(streamName),
		PartitionKey: aws.String(partitionKey),
		Data:         data,
	})
	if err != nil {
		return customerror.ErrorHandler.EventPublishError(err)
	}
	return nil
}
