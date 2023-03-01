package adapter

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

var adapter *dynamodb.Client

func GetDBAdapter() (*dynamodb.Client, *customerror.CustomError) {
	if adapter != nil {
		return adapter, nil
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "production" {
		// Production mode
		return DynamoDBAdapter()
	} else {
		// Development mode
		return DynamoDBAdapter()
	}
}

func DynamoDBAdapter() (*dynamodb.Client, *customerror.CustomError) {
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
