package integration

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/publisher"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/repository"
	"github.com/kokiebisu/rental-storage/service-user/internal/client"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/kokiebisu/rental-storage/service-user/test/data"

	_ "github.com/lib/pq"
)

func TestMain(m *testing.M) {
	setup()
	streamSetup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err *customerror.CustomError
	// Start a PostgreSQL container
	data.PostgresClient, err = client.GetPostgresClient()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	data.UserRepository = repository.NewUserRepository(data.PostgresClient)
	// Set up tables
	err = data.UserRepository.Setup()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	data.KinesisClient, err = client.GetKinesisClient()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	data.UserPublisher = publisher.NewUserPublisher(data.KinesisClient)
}

func streamSetup() {
	ctx := context.Background()
	var err error
	// Create a new Kinesis stream
	streamName := "local-EventStream"
	shardCount := int32(1)
	_, err = data.KinesisClient.CreateStream(ctx, &kinesis.CreateStreamInput{
		StreamName: aws.String(streamName),
		ShardCount: aws.Int32(shardCount),
	})
	if err != nil {
		var aerr *types.ResourceInUseException
		if errors.As(err, &aerr) {
			log.Println("Stream already exists:", aerr)
		} else {
			log.Fatalf(err.Error())
		}
	}

	// Wait for the stream to become active
	streamStatus := ""
	for streamStatus != "ACTIVE" {
		stream, err := data.KinesisClient.DescribeStream(ctx, &kinesis.DescribeStreamInput{
			StreamName: aws.String(streamName),
		})
		if err != nil {
			fmt.Println("Failed to describe Kinesis stream:", err)
			os.Exit(1)
		}
		streamStatus = string(stream.StreamDescription.StreamStatus)
		time.Sleep(1 * time.Second)
	}

	// Put a record into the stream
	partitionKey := "123"
	_, err = data.KinesisClient.PutRecord(ctx, &kinesis.PutRecordInput{
		StreamName:   aws.String(streamName),
		PartitionKey: aws.String(partitionKey),
		Data:         []byte("Hello, world!"),
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// Close the database connection
func teardown() {
	// dropTables()
	data.PostgresClient.Close()
}
