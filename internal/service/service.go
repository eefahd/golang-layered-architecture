package service

import (
	"golang/internal/store/interfaces"
	"golang/internal/utils/messaging"
)

type Service struct {
	ContactService *ContactService
}

func NewService(
	store *interfaces.Store,
	emailClient *messaging.EmailClient,
) *Service {
	return &Service{
		ContactService: NewContactService(store.Contact, emailClient),
	}
}
