package adapter

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

var instance *kinesis.Client

func GetPublisherAdapter() (*kinesis.Client, *customerror.CustomError) {
	if instance != nil {
		return instance, nil
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "production" {
		// Production mode
		return NewKinesisAdapter()
	} else {
		// Development mode
		return NewDockerAdapter()
	}
}

func NewKinesisAdapter() (*kinesis.Client, *customerror.CustomError) {
	defaultConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}
	// Create a Kinesis client with additional configuration
	client := kinesis.NewFromConfig(defaultConfig)
	return client, nil
}

func NewDockerAdapter() (*kinesis.Client, *customerror.CustomError) {
	ctx := context.Background()

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           "http://localhost:4567",
			SigningRegion: "us-east-1",
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("rental-storage"), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		log.Fatal(err)
	}
	client := kinesis.NewFromConfig(cfg)

	// Create a new Kinesis stream
	streamName := "local-EventStream"
	shardCount := int32(1)
	_, err = client.CreateStream(ctx, &kinesis.CreateStreamInput{
		StreamName: aws.String(streamName),
		ShardCount: aws.Int32(shardCount),
	})
	if err != nil {
		var aerr *types.ResourceInUseException
		if errors.As(err, &aerr) {
			log.Println("Stream already exists:", aerr)
		} else {
			return nil, customerror.ErrorHandler.CreateStreamError(err)
		}
	}

	// Wait for the stream to become active
	streamStatus := ""
	for streamStatus != "ACTIVE" {
		stream, err := client.DescribeStream(ctx, &kinesis.DescribeStreamInput{
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
	data := []byte("Hello, world!")
	_, err = client.PutRecord(ctx, &kinesis.PutRecordInput{
		StreamName:   aws.String(streamName),
		PartitionKey: aws.String(partitionKey),
		Data:         data,
	})
	if err != nil {
		return nil, customerror.ErrorHandler.PutRecordError(err)
	}
	return client, nil
}
