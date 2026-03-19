package config

import "sync/atomic"

type ApiConfig struct {
    FileserverHits atomic.Int32 // to make sure shared state of data.
}

