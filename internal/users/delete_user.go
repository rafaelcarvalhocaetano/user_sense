package users

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *UserModule) DeleteUser(id uuid.UUID) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	if err = s.userRepository.DeleteUser(user); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if err = s.userSearch.DeleteUser(id); err != nil {
		return fmt.Errorf("user deleted but search removal failed: %w", err)
	}

	return nil
}
