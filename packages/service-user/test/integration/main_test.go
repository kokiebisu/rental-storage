package integration

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/publisher"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/repository"
	"github.com/kokiebisu/rental-storage/service-user/internal/client"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"

	_ "github.com/lib/pq"
)

var (
	pc            *sql.DB
	UserRepo      *repository.UserRepository
	kinesisClient *kinesis.Client
	UserPublisher port.UserPublisher
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err *customerror.CustomError
	// Start a PostgreSQL container
	pc, err = client.GetPostgresClient()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	UserRepo = repository.NewUserRepository(pc)
	// Set up tables
	err = UserRepo.Setup()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	kinesisClient, err = client.GetKinesisClient()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	UserPublisher = publisher.NewUserPublisher(kinesisClient)
}

// Close the database connection
func teardown() {
	// dropTables()
	pc.Close()
}
