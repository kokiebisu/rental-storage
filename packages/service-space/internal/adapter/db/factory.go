package db

import (
	"database/sql"
	"os"

	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

var instance *sql.DB

func GetInstance() (*sql.DB, *customerror.CustomError) {
	if instance != nil {
		return instance, nil
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "production" {
		// Production mode
		return NewRDSPostgresDB()
	} else {
		// Development mode
		return NewDockerPostgresDB()
	}
}
