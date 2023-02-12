package db

import (
	"database/sql"

	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

func New() (*sql.DB, *customerror.CustomError) {
	return NewPostgres()
}
