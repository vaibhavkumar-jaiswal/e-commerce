// Package models contains UserPassword models and its related methods
package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// UserPassword holds encrypted password for a specific user
type UserPassword struct {
	base.Model     `swaggerignore:"true"`
	UserPasswordID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique" json:"user_password_id"`
	Password       string    `json:"password"`
	UserID         uuid.UUID `gorm:"not null;uniqueIndex" json:"user_id"`
}

// TableName sets custom table name for UserPassword model
func (UserPassword) TableName() string {
	return "user_passwords"
}
