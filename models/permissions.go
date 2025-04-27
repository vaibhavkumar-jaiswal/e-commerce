// Package models contains Permission and its related models
package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// Permission represents a permission for a module in the system
type Permission struct {
	base.Model   `swaggerignore:"true"`
	PermissionID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique" json:"permission_id"`
	Code         string    `gorm:"unique;not null" json:"code"`
	Action       string    `gorm:"not null" json:"action"`
	Module       string    `gorm:"not null" json:"module"`
}

// TableName sets custom table name for Permission model
func (Permission) TableName() string {
	return "permissions"
}
