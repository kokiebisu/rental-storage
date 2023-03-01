package integration

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/kokiebisu/rental-storage/service-user/internal/publisher"
	"github.com/kokiebisu/rental-storage/service-user/internal/repository"
	_ "github.com/lib/pq"
)

var dbInstance *sql.DB
var UserRepo *repository.UserRepository
var kinesisClient *kinesis.Client
var UserPublisher port.UserPublisher

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err *customerror.CustomError
	// Start a PostgreSQL container
	dbInstance, err = adapter.GetDBAdapter()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	UserRepo = repository.NewUserRepository(dbInstance)
	// Set up tables
	err = UserRepo.Setup()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	kinesisClient, err = adapter.GetPublisherAdapter()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	UserPublisher = publisher.NewUserPublisher(kinesisClient)
}

// Close the database connection
func teardown() {
	// dropTables()
	dbInstance.Close()
}
