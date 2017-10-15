package search

import (
	"log"

	"go-typesense-app/internal/models"

	"github.com/google/uuid"
	"github.com/typesense/typesense-go/typesense/api"
)

type UserSearchServiceImpl struct {
	sense *TypesenseClient
}

func NewUserSearchService(sense *TypesenseClient) *UserSearchServiceImpl {
	return &UserSearchServiceImpl{sense: sense}
}

func (t *UserSearchServiceImpl) CreateUserCollection() error {
	optional := true

	schema := &api.CollectionSchema{
		Name: "users",
		Fields: []api.Field{
			{
				Name: "id",
				Type: "string",
			},
			{
				Name: "name",
				Type: "string",
			},
			{
				Name: "email",
				Type: "string",
			},
			{
				Name:     "phone",
				Type:     "string",
				Optional: &optional,
			},
			{
				Name:     "city",
				Type:     "string",
				Optional: &optional,
			},
		},
	}

	_, err := t.sense.Client.Collections().Create(schema)
	if err != nil {
		if err.Error() == "status: 409 response: {\"message\": \"A collection with name `users` already exists.\"}" {
			log.Println("Users collection already exists, skipping creation")
		} else {
			log.Printf("Error creating collection: %v", err)
		}
		return nil
	}

	log.Println("Users collection created successfully")
	return nil
}

func (t *UserSearchServiceImpl) IndexUser(user *models.User) error {
	document := models.UserSearchDocument{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		City:  user.City,
	}

	_, err := t.sense.Client.Collection("users").Documents().Upsert(document)
	return err
}

func (t *UserSearchServiceImpl) DeleteUser(userID uuid.UUID) error {
	_, err := t.sense.Client.Collection("users").Document(userID.String()).Delete()
	return err
}

func (t *UserSearchServiceImpl) SearchUsers(query string) ([]models.UserSearchDocument, error) {
	queryBy := "name,email,phone,city"
	perPage := 50

	searchParams := &api.SearchCollectionParams{
		Q:       query,
		QueryBy: queryBy,
		PerPage: &perPage,
	}

	result, err := t.sense.Client.Collection("users").Documents().Search(searchParams)
	if err != nil {
		return nil, err
	}

	var users []models.UserSearchDocument
	for _, hit := range *result.Hits {
		doc := *hit.Document
		user := models.UserSearchDocument{
			ID:    getStringValue(doc["id"]),
			Name:  getStringValue(doc["name"]),
			Email: getStringValue(doc["email"]),
			Phone: getStringValue(doc["phone"]),
			City:  getStringValue(doc["city"]),
		}
		users = append(users, user)
	}

	return users, nil
}

func getStringValue(value interface{}) string {
	if value == nil {
		return ""
	}
	return value.(string)
}
