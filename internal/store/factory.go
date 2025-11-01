package store

import (
	"database/sql"
	"fmt"

	"golang/internal/config"
	"golang/internal/store/filestore"
	"golang/internal/store/interfaces"
	"golang/internal/store/postgres"
	"golang/internal/store/sqlite"
)

// New creates a new Storage instance based on the configuration
// This factory function returns the appropriate store implementation
func New(cfg *config.Config, db *sql.DB) (*interfaces.Store, error) {
	switch cfg.Store.Type {
	case config.SQLite:
		if db == nil {
			return nil, fmt.Errorf("database connection required for SQLite store")
		}
		return sqlite.NewStorage(db), nil
	case config.Postgres:
		if db == nil {
			return nil, fmt.Errorf("database connection required for Postgres store")
		}
		return postgres.NewStorage(db), nil
	case config.FileStore:
		return filestore.NewStorage(cfg.Store.FileStore.FilePath)
	default:
		return nil, fmt.Errorf("unsupported store type: %s", cfg.Store.Type)
	}
}
