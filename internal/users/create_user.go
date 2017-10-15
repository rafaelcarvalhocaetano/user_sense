package users

import (
	"fmt"

	"go-typesense-app/internal/models"
)

func (s *UserModule) CreateUser(user *models.User) error {
	if err := s.userRepository.CreateUser(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.userSearch.IndexUser(user); err != nil {
		return fmt.Errorf("user created but search indexing failed: %w", err)
	}

	return nil
}
