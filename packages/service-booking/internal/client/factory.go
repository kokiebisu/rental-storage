package client

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

var dynamodbClient *dynamodb.Client

func GetDynamoDBClient() (*dynamodb.Client, *customerror.CustomError) {
	if dynamodbClient != nil {
		return dynamodbClient, nil
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "test" {
		// Development mode
		return getDynamoDBDockerClient()
	} else {
		// Production mode
		return getDynamoDBClient()
	}
}
