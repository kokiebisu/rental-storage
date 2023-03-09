package client

import (
	"database/sql"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
)

var (
	pc *sql.DB
	kc *kinesis.Client
)

func GetPostgresClient() (*sql.DB, *customerror.CustomError) {
	if pc != nil {
		return pc, nil
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "test" {
		// Development mode
		return getPostgresDockerClient()
	} else {
		// Production mode
		return getPostgresClient()
	}
}

func GetKinesisClient() (*kinesis.Client, *customerror.CustomError) {
	if kc != nil {
		return kc, nil
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = os.Getenv("GO_ENV")
	}

	if env == "test" {
		// Development mode
		return getKinesisDockerClient()
	} else {
		// Production mode
		return getKinesisClient()
	}
}
