package client

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/lib/pq"
)

func getPostgresClient() (*sql.DB, *customerror.CustomError) {
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

func getPostgresDockerClient() (*sql.DB, *customerror.CustomError) {
	fmt.Println("YO")
	port := "5432"
	containerPort := port + "/tcp"

	dbUsername := "postgres"
	dbPassword := "password"
	dbName := "space_db"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres",
			ExposedPorts: []string{containerPort},
			Env: map[string]string{
				"POSTGRES_USER":     dbUsername,
				"POSTGRES_PASSWORD": dbPassword,
				"POSTGRES_DB":       dbName,
			},
			WaitingFor: wait.ForLog("database system is ready to accept connections"),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := container.Terminate(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Get the host and port of the container
	host, err := container.Host(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Set up the database connection
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, dbUsername, dbPassword, dbName)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
