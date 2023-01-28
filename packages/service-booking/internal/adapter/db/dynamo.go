package db

import (
	"context"

	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type NoSQLClient struct {
	client    *dynamodb.Client
	tableName string
}

func NewNoSQLClient() (*NoSQLClient, *errors.CustomError) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("TABLE_NAME")
	if tableName == "" {
		return nil, errors.ErrorHandler.InternalServerError()
	}
	return &NoSQLClient{
		client:    client,
		tableName: tableName,
	}, nil
}

func (c *NoSQLClient) Save(booking booking.Entity) *errors.CustomError {
	bookingDTO := booking.ToDTO()
	item, err := attributevalue.MarshalMap(bookingDTO)
	if err != nil {
		return errors.ErrorHandler.InternalServerError()
	}
	_, err = c.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:              aws.String(c.tableName),
		Item:                   item,
		ReturnConsumedCapacity: "TOTAL",
	})
	if err != nil {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}

func (c NoSQLClient) Delete(id string) *errors.CustomError {
	_, err := c.client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		TableName: &c.tableName,
	})
	if err != nil {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}

func (c NoSQLClient) FindById(id string) (booking.Entity, *errors.CustomError) {
	output, err := c.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		TableName: &c.tableName,
	})
	if err != nil {
		return booking.Entity{}, errors.ErrorHandler.InternalServerError()
	}
	target := booking.Entity{}
	err = attributevalue.UnmarshalMap(output.Item, &target)
	if err != nil {
		return booking.Entity{}, errors.ErrorHandler.InternalServerError()
	}
	return target, nil
}

func (c NoSQLClient) FindManyByUserId(userId string) ([]booking.Entity, *errors.CustomError) {
	shouldScanIndexForward := false
	indexName := "BookingUserIdIndex"
	output, err := c.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              &c.tableName,
		IndexName:              &indexName,
		KeyConditionExpression: aws.String("UserId = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: userId},
		},
		ScanIndexForward: &shouldScanIndexForward,
	})
	if err != nil {
		log.Fatal(err)
		return []booking.Entity{}, errors.ErrorHandler.InternalServerError()
	}
	targets := []booking.Entity{}
	for _, i := range output.Items {
		target := booking.DTO{}
		err = attributevalue.UnmarshalMap(i, &target)
		if err != nil {
			return []booking.Entity{}, errors.ErrorHandler.InternalServerError()
		}
		entity, err := target.ToEntity()
		if err != nil {
			return []booking.Entity{}, errors.ErrorHandler.InternalServerError()
		}
		targets = append(targets, entity)
	}

	if err != nil {
		return []booking.Entity{}, errors.ErrorHandler.InternalServerError()
	}
	return targets, nil
}
