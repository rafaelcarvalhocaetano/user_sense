package users

import (
	"fmt"

	"go-typesense-app/internal/models"
)

func (s *UserModule) GetAllUsers() ([]models.User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	return users, nil
}
