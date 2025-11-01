package interfaces

import (
	"context"
	"golang/internal/models"
)

// ContactRepositoryInterface defines the contract for contact data access
// This abstraction allows us to swap implementations (SQLite, Postgres, etc.)
type ContactRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.Contact, error)
	GetByID(ctx context.Context, id int) (*models.Contact, error)
	Create(ctx context.Context, contact models.Contact) (int, error)
	Update(ctx context.Context, contact models.Contact) error
	Delete(ctx context.Context, id int) error
}
