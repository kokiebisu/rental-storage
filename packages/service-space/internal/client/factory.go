package client

import (
	"database/sql"
	"os"

	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

var pc *sql.DB

func GetPostgresClient() (*sql.DB, *customerror.CustomError) {
	if pc != nil {
		return pc, nil
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "test" {
		// Development mode
		return getPostgresDockerClient()
	} else {
		// Production mode
		return getPostgresClient()
	}
}
