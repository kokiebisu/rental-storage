package db

import (
	"database/sql"

	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

func New() (*sql.DB, *errors.CustomError) {
	return NewPostgres()
}
