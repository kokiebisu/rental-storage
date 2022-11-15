package sender

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"

	"github.com/google/uuid"

	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
)

type KinesisSender struct {
	client *kinesis.Kinesis
}

type Event struct {
	SourceEntity string
	EventName    string
	Data         domain.UserDTO
}

func New() *KinesisSender {
	region := "us-east-1"
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: &region,
	}))

	// Create a Kinesis client with additional configuration
	client := kinesis.New(mySession)
	return &KinesisSender{
		client: client,
	}
}

func (s *KinesisSender) UserCreated(user domain.UserDTO) error {
	event := &Event{
		SourceEntity: "User",
		EventName:    "created",
		Data:         user,
	}

	encoded, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return s.publish(encoded)
}

func (s *KinesisSender) publish(data []byte) error {
	streamName := fmt.Sprintf("%s-EventStream", os.Getenv("STAGE"))
	partitionKey := uuid.New().String()
	_, err := s.client.PutRecord(&kinesis.PutRecordInput{
		StreamName:   &streamName,
		PartitionKey: &partitionKey,
		Data:         data,
	})
	return err
}
