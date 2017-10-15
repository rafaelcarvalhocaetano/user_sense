package users

import (
	"fmt"

	"go-typesense-app/internal/models"

	"github.com/google/uuid"
)

func (s *UserModule) UpdateUser(id uuid.UUID, updateData *models.User) (*models.User, error) {
	user, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if err = s.userRepository.UpdateUser(id, updateData); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	if err = s.userSearch.IndexUser(user); err != nil {
		return user, fmt.Errorf("user updated but search indexing failed: %w", err)
	}

	return user, nil
}
