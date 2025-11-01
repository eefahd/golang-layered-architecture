package database

import (
	"database/sql"
	"fmt"

	"golang/internal/config"
)

type Database interface {
	Connect() error
	Close() error
	GetDB() *sql.DB
}

// New creates a new Database instance based on the configuration
// This is the factory function that returns the appropriate database implementation
func New(cfg *config.Config) (Database, error) {
	switch cfg.Store.Type {
	case config.SQLite:
		return NewSQLiteDB(cfg.Store.SQLite), nil
	case config.Postgres:
		return NewPostgresDB(cfg.Store.Postgres), nil
	case config.FileStore:
		return NewFileStoreDB(cfg.Store.FileStore), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Store.Type)
	}
}
