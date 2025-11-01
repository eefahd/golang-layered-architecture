package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type StoreType string

const (
	SQLite    StoreType = "sqlite"
	Postgres  StoreType = "postgres"
	FileStore StoreType = "filestore"
)

type Config struct {
	Store  StoreConfig  `json:"store"`
	Server ServerConfig `json:"server"`
	Email  EmailConfig  `json:"email"`
}

type StoreConfig struct {
	Type      StoreType       `json:"type"`
	SQLite    SQLiteConfig    `json:"sqlite"`
	FileStore FileStoreConfig `json:"filestore"`
	Postgres  PostgresConfig  `json:"postgres"`
}

type SQLiteConfig struct {
	DBPath     string `json:"db_path"`
	SchemaPath string `json:"schema_path"`
}

type FileStoreConfig struct {
	FilePath string `json:"file_path"`
}

type PostgresConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	DBName     string `json:"dbname"`
	SchemaPath string `json:"schema_path"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type EmailConfig struct {
	Token string `json:"token"`
}

func Load(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	switch cfg.Store.Type {
	case SQLite, Postgres, FileStore:
	default:
		return nil, fmt.Errorf("invalid store type: %s (must be sqlite, postgres, or filestore)", cfg.Store.Type)
	}

	return &cfg, nil
}
