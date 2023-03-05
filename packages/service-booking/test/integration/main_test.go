package integration

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/repository"
	"github.com/kokiebisu/rental-storage/service-booking/internal/client"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/service"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"

	"github.com/kokiebisu/rental-storage/service-booking/test/data"
)

func TestMain(m *testing.M) {
	setup()
	tableSetup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	var err *customerror.CustomError
	// Start a PostgreSQL container
	data.DBClient, err = client.GetDBClient()
	if err != nil {
		log.Fatal(err)
	}
	br, err := repository.NewBookingRepository(data.DBClient)
	if err != nil {
		log.Fatal(err)
	}
	data.BookingService = service.NewBookingService(br)
}

func tableSetup() {
	var apiErr smithy.APIError

	tableName := "booking"
	// Define the input parameters for the CreateTable operation
	input := &dynamodb.CreateTableInput{
		TableName: aws.String(tableName), // Replace with your table name
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("UId"),
				KeyType:       types.KeyTypeHash,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("UId"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("UserId"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("SpaceId"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("CreatedAt"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		BillingMode: types.BillingModePayPerRequest,
		StreamSpecification: &types.StreamSpecification{
			StreamEnabled:  aws.Bool(true),
			StreamViewType: types.StreamViewTypeNewAndOldImages,
		},
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("BookingSpaceIdIndex"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("SpaceId"),
						KeyType:       types.KeyTypeHash,
					},
					{
						AttributeName: aws.String("CreatedAt"),
						KeyType:       types.KeyTypeRange,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
			},
			{
				IndexName: aws.String("BookingCreatedAtIndex"),
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("UserId"),
						KeyType:       types.KeyTypeHash,
					},
					{
						AttributeName: aws.String("CreatedAt"),
						KeyType:       types.KeyTypeRange,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
			},
		},
	}

	_, err := data.DBClient.CreateTable(context.TODO(), input)
	if err != nil {
		if errors.As(err, &apiErr) {
			switch apiErr.(type) {
			case *types.ResourceInUseException:
				fmt.Println("Reusing existing table...")
			default:
				log.Fatalf("unknown error: %v\n", err)
			}
		}
	}
}
