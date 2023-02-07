package db

import customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"

func New() (*NoSQLClient, *customerror.CustomError) {
	return NewNoSQLClient()
}
