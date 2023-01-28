package db

import errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"

func New() (*NoSQLClient, *errors.CustomError) {
	return NewNoSQLClient()
}
