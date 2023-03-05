package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingRepository struct {
	client    *dynamodb.Client
	tableName string
}

func NewBookingRepository(client *dynamodb.Client) (*BookingRepository, *customerror.CustomError) {
	var tableName string
	if os.Getenv("TABLE_NAME") == "" {
		tableName = "booking"
	} else {
		tableName = os.Getenv("TABLE_NAME")
	}
	if tableName == "" {
		return nil, customerror.ErrorHandler.InternalServerError("cannot retrieve table name", nil)
	}
	return &BookingRepository{
		client,
		tableName,
	}, nil
}

func (c *BookingRepository) Save(booking booking.Entity) *customerror.CustomError {
	bookingDTO := booking.ToDTO()
	item, err := attributevalue.MarshalMap(bookingDTO)
	fmt.Println("Save 1: ", c.tableName)
	if err != nil {
		return customerror.ErrorHandler.InternalServerError("cannot marshal map", err)
	}
	_, err = c.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:              aws.String(c.tableName),
		Item:                   item,
		ReturnConsumedCapacity: "TOTAL",
	})
	if err != nil {
		fmt.Println("Save 1.5: ", err)
		return customerror.ErrorHandler.InternalServerError("cannot perform PutItem operation", err)
	}
	return nil
}

func (c BookingRepository) Delete(uid string) *customerror.CustomError {
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

func (c BookingRepository) FindById(uid string) (booking.Entity, *customerror.CustomError) {
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

func (c BookingRepository) FindManyByUserId(userId string) ([]booking.Entity, *customerror.CustomError) {
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

func (c BookingRepository) FindManyBySpaceId(spaceId string) ([]booking.Entity, *customerror.CustomError) {
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