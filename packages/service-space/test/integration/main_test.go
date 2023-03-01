package integration

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/kokiebisu/rental-storage/service-space/internal/adapter"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
	"github.com/kokiebisu/rental-storage/service-space/internal/repository"
	_ "github.com/lib/pq"
)

var dbInstance *sql.DB
var Repo *repository.SpaceRepository

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
	Repo = repository.NewSpaceRepository(dbInstance)
	// Set up tables
	err = Repo.Setup()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func dropTables() {
	var err error
	_, err = dbInstance.Exec("DELETE FROM images")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	_, err = dbInstance.Exec("DELETE FROM spaces")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	_, err = dbInstance.Exec("DELETE FROM locations")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Close the database connection
func teardown() {
	dropTables()
	dbInstance.Close()
}
