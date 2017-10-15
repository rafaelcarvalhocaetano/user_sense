package database

import (
	"go-typesense-app/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	UpdateUser(id uuid.UUID, user *models.User) error
	CreateUser(user *models.User) error
	DeleteUser(user *models.User) error
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (ur *UserRepositoryImpl) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepositoryImpl) UpdateUser(id uuid.UUID, updateData *models.User) error {
	return ur.db.Model(&models.User{}).Where("id = ?", id).Updates(updateData).Error
}

func (ur *UserRepositoryImpl) DeleteUser(user *models.User) error {
	return ur.db.Delete(user).Error
}

func (ur *UserRepositoryImpl) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := ur.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
