// Package models provides role related models
package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// Role represents a user role in the system.
type Role struct {
	base.Model  `swaggerignore:"true"`
	RoleID      uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"role_id"`
	Name        string       `gorm:"not null" json:"name"`
	Code        string       `gorm:"unique; not null" json:"code"`
	Description string       `json:"description"`
	Permissions []Permission `gorm:"many2many:role_permissions;joinForeignKey:RoleID;joinReferences:PermissionID;"`
}

// TableName sets custom table name for Role model
func (Role) TableName() string {
	return "roles"
}
