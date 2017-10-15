package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `sql:"timestamp with time zone" gorm:"not null;default:CURRENT_TIMESTAMP;autoCreateTime:true" json:"created"`
	UpdatedAt time.Time  `sql:"timestamp with time zone" gorm:"not null;default:CURRENT_TIMESTAMP;autoUpdateTime:true" json:"updated"`
	DeletedAt *time.Time `sql:"timestamp with time zone" gorm:"index" json:"deleted,omitempty"`
	Name      string     `json:"name" gorm:"not null;size:100"`
	Email     string     `json:"email" gorm:"uniqueIndex;not null;size:100"`
	Phone     string     `json:"phone" gorm:"size:20"`
	City      string     `json:"city" gorm:"size:100"`
}

type UserSearchDocument struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	City  string `json:"city"`
}
