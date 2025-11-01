package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"golang/internal/models"
	"golang/internal/store/interfaces"
	"golang/internal/utils/messaging"
)

// ContactService handles business logic for contacts
type ContactService struct {
	repo        interfaces.ContactRepositoryInterface
	emailClient *messaging.EmailClient
}

func NewContactService(repo interfaces.ContactRepositoryInterface, emailClient *messaging.EmailClient) *ContactService {
	return &ContactService{
		repo:        repo,
		emailClient: emailClient,
	}
}

func (s *ContactService) GetAll(ctx context.Context) ([]models.Contact, error) {
	return s.repo.GetAll(ctx)
}

func (s *ContactService) GetByID(ctx context.Context, id int) (*models.Contact, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ContactService) Create(ctx context.Context, contact models.Contact) (*models.Contact, error) {
	id, err := s.repo.Create(ctx, contact)
	if err != nil {
		return nil, err
	}
	contact.ID = id
	return &contact, nil
}

// UpdateAndNotify updates a contact and sends a notification email
// This demonstrates business logic: multiple operations orchestrated together
func (s *ContactService) UpdateAndNotify(ctx context.Context, contact models.Contact) error {
	log.Printf("Service: Updating contact ID %d", contact.ID)

	// Step 1: Validate email format (business rule)
	contact.Email = strings.ToLower(strings.TrimSpace(contact.Email))
	if !strings.Contains(contact.Email, "@") {
		return fmt.Errorf("invalid email format")
	}

	// Step 2: Get old contact data (to compare)
	oldContact, err := s.repo.GetByID(ctx, contact.ID)
	if err != nil {
		return fmt.Errorf("contact not found: %w", err)
	}

	// Step 3: Update in database
	if err := s.repo.Update(ctx, contact); err != nil {
		return fmt.Errorf("failed to update contact: %w", err)
	}

	// Step 4: Send notification if email changed (business orchestration)
	if oldContact.Email != contact.Email {
		go func() {
			email := models.EmailMessage{
				To:      contact.Email,
				Subject: "Contact Information Updated",
				Body:    fmt.Sprintf("Hi %s, your contact information has been updated.", contact.FirstName),
			}
			if err := s.emailClient.SendEmail(email); err != nil {
				log.Printf("Warning: Failed to send notification: %v", err)
			}
		}()
	}

	log.Printf("Service: Contact updated and notification sent")
	return nil
}

func (s *ContactService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
