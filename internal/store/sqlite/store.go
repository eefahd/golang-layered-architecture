package sqlite

import (
	"database/sql"
	"golang/internal/store/interfaces"
)

func NewStorage(db *sql.DB) *interfaces.Store {
	return &interfaces.Store{
		Contact: NewContactRepository(db),
	}
}
