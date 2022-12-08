package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

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
	itemsStringified := []types.AttributeValue{}
	for _, i := range booking.Items {
		itemStringified, err := json.Marshal(i)
		if err != nil {
			return errors.New("failed to marshal item")
		}
		itemsStringified = append(itemsStringified, &types.AttributeValueMemberS{Value: string(itemStringified)})
	}
	_, err := c.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(c.tableName),
		Item: map[string]types.AttributeValue{
			"id":         &types.AttributeValueMemberS{Value: booking.Id},
			"status":     &types.AttributeValueMemberS{Value: string(booking.Status)},
			"owner_id":   &types.AttributeValueMemberS{Value: booking.OwnerId},
			"listing_id": &types.AttributeValueMemberS{Value: booking.ListingId},
			"amount": &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
				"Value":    &types.AttributeValueMemberN{Value: strconv.Itoa(int(booking.Amount.Value))},
				"Currency": &types.AttributeValueMemberS{Value: booking.Amount.Currency},
			}},
			"items":      &types.AttributeValueMemberL{Value: itemsStringified},
			"created_at": &types.AttributeValueMemberS{Value: booking.CreatedAt.Format("YYYY-MM-DD")},
			"updated_at": &types.AttributeValueMemberS{Value: booking.UpdatedAt.Format("YYYY-MM-DD")},
		},
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
	output, err := c.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              &c.tableName,
		KeyConditionExpression: aws.String("userId = :hashKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":hashKey": &types.AttributeValueMemberS{Value: userId},
		},
	})
	if err != nil {
		return []booking.Entity{}, errors.New("error occured when finding item from dynamodb by id")
	}
	targets := []booking.Entity{}
	for _, i := range output.Items {
		target := booking.Entity{}
		err = attributevalue.UnmarshalMap(i, &target)
	}

	if err != nil {
		return []booking.Entity{}, errors.New("error occured when unmarshalling from dynamodb")
	}
	return targets, nil
}
