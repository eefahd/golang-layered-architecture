package filestore

import (
	"fmt"
	"golang/internal/store/interfaces"
)

func NewStorage(file_path string) (*interfaces.Store, error) {
	contactRepo, err := NewContactRepository(file_path)
	if err != nil {
		return nil, fmt.Errorf("failed to create file-based contact repository: %w", err)
	}

	return &interfaces.Store{
		Contact: contactRepo,
	}, nil
}
