package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang/internal/config"
)

type FileStoreDB struct {
	config config.FileStoreConfig
}

func NewFileStoreDB(cfg config.FileStoreConfig) *FileStoreDB {
	return &FileStoreDB{
		config: cfg,
	}
}

func (f *FileStoreDB) Connect() error {
	dir := filepath.Dir(f.config.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create file store directory: %w", err)
	}

	log.Printf("FileStore connected: %s", f.config.FilePath)

	return nil
}

func (f *FileStoreDB) Close() error {
	log.Printf("FileStore connection closed (no-op)")
	return nil
}

func (s *FileStoreDB) GetDB() *sql.DB {
	return nil
}
