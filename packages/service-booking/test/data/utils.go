package data

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-booking/mocks"
)

var (
	DBClient        *dynamodb.Client
	MockBookingRepo *mocks.BookingRepository
	BookingService  port.BookingService
)
