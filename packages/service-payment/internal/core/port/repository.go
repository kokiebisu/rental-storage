package port

import (
	"database/sql"
)

type CustomerRepository struct {
	Db *sql.DB
}