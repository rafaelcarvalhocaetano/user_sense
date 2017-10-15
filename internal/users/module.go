package users

import (
	"go-typesense-app/internal/database"
	"go-typesense-app/internal/models"
	"go-typesense-app/internal/search"

	"github.com/google/uuid"
)

type UserModuleInterface interface {
	InitializeSearchCollection() error
	CreateUser(user *models.User) error
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(id uuid.UUID, updateData *models.User) (*models.User, error)
	DeleteUser(id uuid.UUID) error
	SearchUsers(query string) ([]models.UserSearchDocument, error)
}

type UserModule struct {
	userRepository database.UserRepositoryInterface
	userSearch     search.UserSearchInterface
}

func NewUserModule(userRepository database.UserRepositoryInterface, userSearch search.UserSearchInterface) *UserModule {
	return &UserModule{
		userRepository: userRepository,
		userSearch:     userSearch,
	}
}

func (s *UserModule) InitializeSearchCollection() error {
	return s.userSearch.CreateUserCollection()
}
