package config

import (

	"sync/atomic"
	"server_basics.com/internal/database"
)

type ApiConfig struct {
    FileserverHits atomic.Int32 // to make sure shared state of data.
	Database *database.Queries
}

