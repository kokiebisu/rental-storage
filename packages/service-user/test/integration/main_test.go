package integration

import (
	"log"
	"os"
	"testing"

	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/publisher"
	"github.com/kokiebisu/rental-storage/service-user/internal/adapter/repository"
	"github.com/kokiebisu/rental-storage/service-user/internal/client"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
	"github.com/kokiebisu/rental-storage/service-user/test/data"

	_ "github.com/lib/pq"
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
	data.PostgresClient, err = client.GetPostgresClient()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	data.UserRepository = repository.NewUserRepository(data.PostgresClient)
	// Set up tables
	err = data.UserRepository.Setup()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	data.KinesisClient, err = client.GetKinesisClient()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	data.UserPublisher = publisher.NewUserPublisher(data.KinesisClient)
}

// Close the database connection
func teardown() {
	// dropTables()
	data.PostgresClient.Close()
}
