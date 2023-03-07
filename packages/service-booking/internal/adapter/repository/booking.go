package repository

import (
	"context"
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

func (c BookingRepository) FindOneById(uid string) (booking.Entity, *customerror.CustomError) {
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

func (c BookingRepository) FindManyByUserId(userId string, status string) ([]booking.Entity, *customerror.CustomError) {
	output, err := c.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(c.tableName),
		IndexName:              aws.String("BookingUserIdBookingStatusIndex"),
		KeyConditionExpression: aws.String("UserId = :userId AND BookingStatus = :bookingStatus"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userId":        &types.AttributeValueMemberS{Value: userId},
			":bookingStatus": &types.AttributeValueMemberS{Value: status},
		},
		ScanIndexForward: aws.Bool(false),
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

func (c BookingRepository) FindManyBySpaceId(spaceId string, status string) ([]booking.Entity, *customerror.CustomError) {
	output, err := c.client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(c.tableName),
		IndexName:              aws.String("BookingSpaceIdBookingStatusIndex"),
		KeyConditionExpression: aws.String("SpaceId = :spaceId and BookingStatus = :bookingStatus"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":spaceId":       &types.AttributeValueMemberS{Value: spaceId},
			":bookingStatus": &types.AttributeValueMemberS{Value: status},
		},
		ScanIndexForward: aws.Bool(false),
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
