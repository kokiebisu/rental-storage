package client

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

var client *dynamodb.Client

func GetDBClient() (*dynamodb.Client, *customerror.CustomError) {
	if client != nil {
		return client, nil
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "production" {
		// Production mode
		return getProductionClient()
	} else {
		// Development mode
		return getDockerClient()
	}
}
