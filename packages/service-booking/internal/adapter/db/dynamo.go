package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
)

type NoSQLClient struct {
	client    *dynamodb.Client
	tableName string
}

func NewNoSQLClient() (*NoSQLClient, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	client := dynamodb.NewFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	tableName := fmt.Sprintf("%s-%s", os.Getenv("TABLE_NAME"), os.Getenv("STAGE"))
	if tableName == "" {
		return nil, errors.New("TABLE_NAME not properly retrieved")
	}
	return &NoSQLClient{
		client:    client,
		tableName: tableName,
	}, nil
}

func (c *NoSQLClient) Save(booking booking.Entity) error {
	bookingDTO := booking.ToDTO()
	item, err := attributevalue.MarshalMap(bookingDTO)
	if err != nil {
		return err
	}
	_, err = c.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:              aws.String(c.tableName),
		Item:                   item,
		ReturnConsumedCapacity: "TOTAL",
	})
	if err != nil {
		return err
	}
	return nil
}

func (c NoSQLClient) Delete(id string) error {
	_, err := c.client.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		TableName: &c.tableName,
	})
	if err != nil {
		return errors.New("error occured when deleting from dynamodb")
	}
	return nil
}

func (c NoSQLClient) FindById(id string) (booking.Entity, error) {
	output, err := c.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		TableName: &c.tableName,
	})
	if err != nil {
		return booking.Entity{}, errors.New("error occured when finding item from dynamodb by id")
	}
	target := booking.Entity{}
	err = attributevalue.UnmarshalMap(output.Item, &target)
	if err != nil {
		return booking.Entity{}, errors.New("error occured when unmarshalling from dynamodb")
	}
	return target, nil
}

func (c NoSQLClient) FindManyByUserId(userId string) ([]booking.Entity, error) {
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
		return []booking.Entity{}, errors.New("error occured when finding item from dynamodb by id")
	}
	targets := []booking.Entity{}
	for _, i := range output.Items {
		target := booking.DTO{}
		err = attributevalue.UnmarshalMap(i, &target)
		if err != nil {
			return []booking.Entity{}, err
		}
		entity, err := target.ToEntity()
		if err != nil {
			return []booking.Entity{}, err
		}
		targets = append(targets, entity)
	}

	if err != nil {
		return []booking.Entity{}, errors.New("error occured when unmarshalling from dynamodb")
	}
	return targets, nil
}
