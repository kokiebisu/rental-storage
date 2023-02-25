package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func NewDockerPostgresDB() (*sql.DB, *customerror.CustomError) {
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
