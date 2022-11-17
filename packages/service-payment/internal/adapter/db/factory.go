package db

import "database/sql"

func New() (*sql.DB, error) {
	return NewPostgres()
}
