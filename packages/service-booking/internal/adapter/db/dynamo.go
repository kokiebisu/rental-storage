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
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type NoSQLClient struct {
	client    *dynamodb.Client
	tableName string
}

func NewNoSQLClient() (*NoSQLClient, *customerror.CustomError) {
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
		return nil, customerror.ErrorHandler.InternalServerError("cannot retrieve table name", nil)
	}
	return &NoSQLClient{
		client:    client,
		tableName: tableName,
	}, nil
}

func (c *NoSQLClient) Save(booking booking.Entity) *customerror.CustomError {
	bookingDTO := booking.ToDTO()
	item, err := attributevalue.MarshalMap(bookingDTO)
	if err != nil {
		return customerror.ErrorHandler.InternalServerError("cannot marshal map", err)
	}
	_, err = c.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:              aws.String(c.tableName),
		Item:                   item,
		ReturnConsumedCapacity: "TOTAL",
	})
	if err != nil {
		return customerror.ErrorHandler.InternalServerError("cannot perform PutItem operation", err)
	}
	return nil
}

func (c NoSQLClient) Delete(uid string) *customerror.CustomError {
	_, err := c.client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"UId": &types.AttributeValueMemberS{Value: uid},
		},
		TableName: &c.tableName,
	})
	if err != nil {
		return customerror.ErrorHandler.InternalServerError("cannot perform DeleteItem operation", err)
	}
	return nil
}

func (c NoSQLClient) FindById(uid string) (booking.Entity, *customerror.CustomError) {
	output, err := c.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"UId": &types.AttributeValueMemberS{Value: uid},
		},
		TableName: &c.tableName,
	})
	if err != nil {
		return booking.Entity{}, customerror.ErrorHandler.InternalServerError("cannot perform FindById operation", err)
	}
	target := booking.DTO{}
	err = attributevalue.UnmarshalMap(output.Item, &target)
	if err != nil {
		return booking.Entity{}, customerror.ErrorHandler.InternalServerError("cannot unmarshal map", err)
	}
	return target.ToEntity(), nil
}

func (c NoSQLClient) FindManyByUserId(userId string) ([]booking.Entity, *customerror.CustomError) {
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
		return []booking.Entity{}, customerror.ErrorHandler.InternalServerError("cannot perform Query operation by userId", err)
	}
	targets := []booking.Entity{}
	for _, i := range output.Items {
		target := booking.DTO{}
		err = attributevalue.UnmarshalMap(i, &target)
		if err != nil {
			return []booking.Entity{}, customerror.ErrorHandler.InternalServerError("cannot unmarshal map", err)
		}
		entity := target.ToEntity()
		targets = append(targets, entity)
	}
	return targets, nil
}

func (c NoSQLClient) FindManyBySpaceId(spaceId string) ([]booking.Entity, *customerror.CustomError) {
	shouldScanIndexForward := false
	indexName := "BookingSpaceIdIndex"
	output, err := c.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              &c.tableName,
		IndexName:              &indexName,
		KeyConditionExpression: aws.String("SpaceId = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: spaceId},
		},
		ScanIndexForward: &shouldScanIndexForward,
	})
	if err != nil {
		log.Fatal(err)
		return []booking.Entity{}, customerror.ErrorHandler.InternalServerError("cannot perform Query operation by spaceId", err)
	}
	targets := []booking.Entity{}
	for _, i := range output.Items {
		target := booking.DTO{}
		err = attributevalue.UnmarshalMap(i, &target)
		if err != nil {
			return []booking.Entity{}, customerror.ErrorHandler.InternalServerError("cannot unmarshal map", err)
		}
		entity := target.ToEntity()
		targets = append(targets, entity)
	}
	return targets, nil
}
