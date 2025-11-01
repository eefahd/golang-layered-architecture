package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"golang/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	config config.SQLiteConfig
	db     *sql.DB
}

func NewSQLiteDB(cfg config.SQLiteConfig) *SQLiteDB {
	return &SQLiteDB{
		config: cfg,
	}
}

func (s *SQLiteDB) Connect() error {
	// Open database connection
	db, err := sql.Open("sqlite3", s.config.DBPath)
	if err != nil {
		return fmt.Errorf("failed to open SQLite database: %w", err)
	}

	// Store the connection for later cleanup
	s.db = db

	if err := s.executeSchema(); err != nil {
		s.Close()
		return fmt.Errorf("failed to initialize SQLite schema: %w", err)
	}

	log.Printf("SQLite database connected: %s", s.config.DBPath)

	return nil
}

func (s *SQLiteDB) Close() error {
	if s.db != nil {
		log.Printf("Closing SQLite database connection")
		return s.db.Close()
	}
	return nil
}

func (s *SQLiteDB) GetDB() *sql.DB {
	return s.db
}

func (s *SQLiteDB) executeSchema() error {
	schema, err := os.ReadFile(s.config.SchemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	if _, err := s.db.Exec(string(schema)); err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	return nil
}
