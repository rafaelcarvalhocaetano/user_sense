package search

import (
	"go-typesense-app/internal/config"
	"go-typesense-app/internal/models"

	"github.com/google/uuid"
	"github.com/typesense/typesense-go/typesense"
)

type UserSearchInterface interface {
	CreateUserCollection() error
	IndexUser(user *models.User) error
	DeleteUser(userID uuid.UUID) error
	SearchUsers(query string) ([]models.UserSearchDocument, error)
}

type SnapshotInterface interface {
	CreateSnapshot() (map[string]interface{}, error)
}

type TypesenseClient struct {
	Client *typesense.Client
}

func NewTypesenseClient(cfg *config.Config) *TypesenseClient {
	client := typesense.NewClient(
		typesense.WithServer(cfg.Typesense.Host),
		typesense.WithAPIKey(cfg.Typesense.APIKey),
	)

	return &TypesenseClient{Client: client}
}
