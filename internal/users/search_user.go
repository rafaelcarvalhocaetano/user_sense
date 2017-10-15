package users

import (
	"errors"
	"fmt"

	"go-typesense-app/internal/models"
)

func (s *UserModule) SearchUsers(query string) ([]models.UserSearchDocument, error) {
	if query == "" {
		return nil, errors.New("search query cannot be empty")
	}

	users, err := s.userSearch.SearchUsers(query)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}

	return users, nil
}
