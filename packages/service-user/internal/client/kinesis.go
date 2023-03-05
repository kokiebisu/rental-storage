package client

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

func getKinesisClient() (*kinesis.Client, *customerror.CustomError) {
	defaultConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}
	// Create a Kinesis client with additional configuration
	client := kinesis.NewFromConfig(defaultConfig)
	return client, nil
}

func getKinesisDockerClient() (*kinesis.Client, *customerror.CustomError) {
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
		return nil, customerror.ErrorHandler.CreateStreamError(err)
	}
	client := kinesis.NewFromConfig(cfg)
	return client, nil
}
