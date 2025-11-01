package messaging

import (
	"golang/internal/models"
	"log"
)

type EmailClient struct {
	token string
}

func NewEmailClient(token string) *EmailClient {
	return &EmailClient{token: token}
}

func (c *EmailClient) Connect() error {
	log.Printf("Email Client Connected!")
	return nil
}

func (c *EmailClient) SendEmail(emailMessage models.EmailMessage) error {
	log.Printf("Sending email to %s: %s", emailMessage.To, emailMessage.Subject)
	return nil
}
