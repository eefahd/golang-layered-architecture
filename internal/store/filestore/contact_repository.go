package filestore

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"golang/internal/models"

	"golang/internal/store/interfaces"
)

// ContactRepository is the file-based implementation of ContactRepositoryInterface
// It stores contacts in a JSON file, demonstrating an alternative to SQL storage
type ContactRepository struct {
	file_path string
	mu        sync.RWMutex // Protects concurrent access to the file
}

func NewContactRepository(file_path string) (interfaces.ContactRepositoryInterface, error) {
	repo := &ContactRepository{
		file_path: file_path,
	}

	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		if err := repo.writeContacts([]models.Contact{}); err != nil {
			return nil, fmt.Errorf("failed to initialize file store: %w", err)
		}
	}

	return repo, nil
}

func (r *ContactRepository) readContacts() ([]models.Contact, error) {
	data, err := os.ReadFile(r.file_path)
	if err != nil {
		return nil, fmt.Errorf("failed to read contacts file: %w", err)
	}

	var contacts []models.Contact
	if len(data) > 0 {
		if err := json.Unmarshal(data, &contacts); err != nil {
			return nil, fmt.Errorf("failed to unmarshal contacts: %w", err)
		}
	}

	return contacts, nil
}

func (r *ContactRepository) writeContacts(contacts []models.Contact) error {
	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal contacts: %w", err)
	}

	if err := os.WriteFile(r.file_path, data, 0644); err != nil {
		return fmt.Errorf("failed to write contacts file: %w", err)
	}

	return nil
}

func (r *ContactRepository) getNextID(contacts []models.Contact) int {
	maxID := 0
	for _, c := range contacts {
		if c.ID > maxID {
			maxID = c.ID
		}
	}
	return maxID + 1
}

func (r *ContactRepository) GetAll(ctx context.Context) ([]models.Contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.readContacts()
}

func (r *ContactRepository) GetByID(ctx context.Context, id int) (*models.Contact, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	contacts, err := r.readContacts()
	if err != nil {
		return nil, err
	}

	for _, c := range contacts {
		if c.ID == id {
			return &c, nil
		}
	}

	return nil, fmt.Errorf("contact not found")
}

func (r *ContactRepository) Create(ctx context.Context, contact models.Contact) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	contacts, err := r.readContacts()
	if err != nil {
		return 0, err
	}

	contact.ID = r.getNextID(contacts)

	contacts = append(contacts, contact)

	if err := r.writeContacts(contacts); err != nil {
		return 0, err
	}

	return contact.ID, nil
}

func (r *ContactRepository) Update(ctx context.Context, contact models.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	contacts, err := r.readContacts()
	if err != nil {
		return err
	}

	found := false
	for i, c := range contacts {
		if c.ID == contact.ID {
			contacts[i] = contact
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("contact not found")
	}

	return r.writeContacts(contacts)
}

func (r *ContactRepository) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	contacts, err := r.readContacts()
	if err != nil {
		return err
	}

	newContacts := make([]models.Contact, 0, len(contacts))
	for _, c := range contacts {
		if c.ID != id {
			newContacts = append(newContacts, c)
		}
	}

	return r.writeContacts(newContacts)
}
