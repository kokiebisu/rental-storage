package client

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

func getProductionClient() (*dynamodb.Client, *customerror.CustomError) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return client, nil
}

func getDockerClient() (*dynamodb.Client, *customerror.CustomError) {
	ctx := context.TODO()

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           "http://localhost:8000",
			SigningRegion: "us-east-1",
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithSharedConfigProfile("rental-storage"), config.WithEndpointResolverWithOptions(customResolver),
	)

	if err != nil {
		panic(fmt.Sprintf("failed to load config, %v", err))
	}

	client := dynamodb.NewFromConfig(cfg)
	// need to check if the error is ResourceInUseException
	// if err != nil {
	// 	fmt.Println("Error creating table:", err)
	// }
	return client, nil
}
