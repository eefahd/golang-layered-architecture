package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"golang/internal/models"
	"golang/internal/store/interfaces"
)

// ContactRepository is the PostgreSQL implementation of ContactRepositoryInterface
type ContactRepository struct {
	db *sql.DB
}

// NewContactRepository creates a PostgreSQL contact repository
func NewContactRepository(db *sql.DB) interfaces.ContactRepositoryInterface {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) GetAll(ctx context.Context) ([]models.Contact, error) {
	query := "SELECT id, first_name, last_name, email FROM contacts"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var c models.Contact
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, rows.Err()
}

func (r *ContactRepository) GetByID(ctx context.Context, id int) (*models.Contact, error) {
	query := "SELECT id, first_name, last_name, email FROM contacts WHERE id = $1"
	var c models.Contact
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("contact not found")
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ContactRepository) Create(ctx context.Context, contact models.Contact) (int, error) {
	query := "INSERT INTO contacts (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, contact.FirstName, contact.LastName, contact.Email).Scan(&id)
	return id, err
}

func (r *ContactRepository) Update(ctx context.Context, contact models.Contact) error {
	query := "UPDATE contacts SET first_name = $1, last_name = $2, email = $3 WHERE id = $4"
	_, err := r.db.ExecContext(ctx, query, contact.FirstName, contact.LastName, contact.Email, contact.ID)
	return err
}

func (r *ContactRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM contacts WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
