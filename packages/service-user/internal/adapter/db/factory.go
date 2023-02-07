package db

import (
	"database/sql"

	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

func New() (*sql.DB, *customerror.CustomError) {
	return NewPostgres()
}
