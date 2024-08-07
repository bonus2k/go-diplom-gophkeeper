package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID     `gorm:"primary_key;type:uuid" json:"id"`
	Username   string        `gorm:"size:255;not null" json:"username"`
	Password   []byte        `gorm:"size:255;not null" json:"password"`
	Email      string        `gorm:"size:255;not null;unique;index:idx_email" json:"email"`
	CreatedAt  *time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  *time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	SecretData *[]SecretData `gorm:"foreignKey:UserID" json:"secret_data,omitempty"`
}

type SecretData struct {
	ID     uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Type   string    `gorm:"size:255;not null" json:"type"`
	Name   string    `gorm:"size:255;not null" json:"name"`
	Secret []byte    `gorm:"type:bytes;size:20480;not null" json:"secret"`
}
