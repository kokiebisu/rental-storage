package integration

import (
	"log"
	"os"
	"testing"

	"github.com/kokiebisu/rental-storage/service-space/internal/adapter/repository"
	"github.com/kokiebisu/rental-storage/service-space/internal/client"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
	"github.com/kokiebisu/rental-storage/service-space/test/data"

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
	data.SpaceRepository = repository.NewSpaceRepository(data.PostgresClient)
	// Set up tables
	err = data.SpaceRepository.Setup()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func dropTables() {
	var err error
	_, err = data.PostgresClient.Exec("DELETE FROM images")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	_, err = data.PostgresClient.Exec("DELETE FROM spaces")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	_, err = data.PostgresClient.Exec("DELETE FROM locations")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Close the database connection
func teardown() {
	dropTables()
	data.PostgresClient.Close()
}
