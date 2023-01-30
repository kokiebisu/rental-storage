package sender

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/google/uuid"

	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

type KinesisSender struct {
	client *kinesis.Client
}

func New() *KinesisSender {
	defaultConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}
	// Create a Kinesis client with additional configuration
	client := kinesis.NewFromConfig(defaultConfig)
	return &KinesisSender{
		client: client,
	}
}

func (s *KinesisSender) UserCreated(u user.DTO) *errors.CustomError {
	event := map[string]interface{}{
		"sourceEntity": "User",
		"eventName":    "created",
		"data":         u,
	}
	encoded, err := json.Marshal(event)
	if err != nil {
		return errors.ErrorHandler.CustomError("user created event not being created successful", err)
	}
	return s.publish(encoded)
}

func (s *KinesisSender) publish(data []byte) *errors.CustomError {

	streamName := fmt.Sprintf("%s-EventStream", os.Getenv("STAGE"))
	partitionKey := uuid.New().String()

	_, err := s.client.PutRecord(context.TODO(), &kinesis.PutRecordInput{
		StreamName:   &streamName,
		PartitionKey: &partitionKey,
		Data:         data,
	})
	if err != nil {
		return errors.ErrorHandler.CustomError("unable to publish event to kinesis stream", err)
	}
	return nil
}
