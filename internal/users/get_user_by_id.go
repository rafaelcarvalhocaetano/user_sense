package users

import (
	"fmt"

	"go-typesense-app/internal/models"

	"github.com/google/uuid"
)

func (s *UserModule) GetUserByID(id uuid.UUID) (*models.User, error) {
	user, err := s.userRepository.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user, nil
}
