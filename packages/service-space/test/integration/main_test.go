package integration

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/kokiebisu/rental-storage/service-space/internal/adapter/repository"
	"github.com/kokiebisu/rental-storage/service-space/internal/client"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"

	_ "github.com/lib/pq"
)

var (
	pc   *sql.DB
	Repo *repository.SpaceRepository
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
	Repo = repository.NewSpaceRepository(pc)
	// Set up tables
	err = Repo.Setup()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func dropTables() {
	var err error
	_, err = pc.Exec("DELETE FROM images")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	_, err = pc.Exec("DELETE FROM spaces")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	_, err = pc.Exec("DELETE FROM locations")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// Close the database connection
func teardown() {
	dropTables()
	pc.Close()
}
