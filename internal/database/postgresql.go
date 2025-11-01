package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"golang/internal/config"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	config config.PostgresConfig
	db     *sql.DB
}

func NewPostgresDB(cfg config.PostgresConfig) *PostgresDB {
	return &PostgresDB{
		config: cfg,
	}
}

func (p *PostgresDB) Connect() error {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.config.Host, p.config.Port, p.config.User, p.config.Password, p.config.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open PostgreSQL database: %w", err)
	}

	// Store the connection for later cleanup
	p.db = db

	// Test connection
	if err := p.db.Ping(); err != nil {
		p.Close()
		return fmt.Errorf("failed to ping PostgreSQL database: %w", err)
	}

	if err := p.executeSchema(); err != nil {
		p.Close()
		return fmt.Errorf("failed to initialize PostgreSQL schema: %w", err)
	}

	log.Printf("PostgreSQL database connected: %s:%d/%s", p.config.Host, p.config.Port, p.config.DBName)

	return nil
}

func (p *PostgresDB) Close() error {
	if p.db != nil {
		log.Printf("Closing PostgreSQL database connection")
		return p.db.Close()
	}
	return nil
}

func (p *PostgresDB) GetDB() *sql.DB {
	return p.db
}

func (p *PostgresDB) executeSchema() error {
	schema, err := os.ReadFile(p.config.SchemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	if _, err := p.db.Exec(string(schema)); err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	return nil
}
