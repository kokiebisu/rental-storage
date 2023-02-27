package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"database/sql"

	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
	_ "github.com/lib/pq"
)

func NewRDSPostgresDB() (*sql.DB, *customerror.CustomError) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln("Unable to convert DB_PORT")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")

	if err != nil {
		return nil, customerror.ErrorHandler.DbConfigurationError(err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
