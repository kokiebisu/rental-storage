package data

import (
	"database/sql"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/repository"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	"github.com/kokiebisu/rental-storage/service-user/mocks"
)

var (
	MockUserRepo    *mocks.UserRepository
	MockEventSender *mocks.EventSender
	KinesisClient   *kinesis.Client
	PostgresClient  *sql.DB
	UserService     port.UserService
	UserRepository  *repository.UserRepository
	UserPublisher   port.UserPublisher
)
