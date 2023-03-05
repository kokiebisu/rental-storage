package data

import (
	"database/sql"

	"github.com/kokiebisu/rental-storage/service-space/internal/adapter/repository"
)

var (
	PostgresClient  *sql.DB
	SpaceRepository *repository.SpaceRepository
)
