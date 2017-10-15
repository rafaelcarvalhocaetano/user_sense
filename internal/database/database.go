package database

import (
	"go-typesense-app/internal/config"
	"go-typesense-app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg *config.Config) *gorm.DB {
	var err error
	db, err := gorm.Open(postgres.Open(cfg.Database.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Database connection failed")
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	return db
}
