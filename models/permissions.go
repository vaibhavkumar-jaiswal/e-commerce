// Package models contains Permission and its related models
package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// Module represents a module in the system
type Module struct {
	base.Model `swaggerignore:"true"`
	ModuleID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique" json:"module_id"`
	Name       string    `gorm:"unique;not null" json:"name"`
}

// Permission represents a permission for a module in the system
type Permission struct {
	base.Model   `swaggerignore:"true"`
	PermissionID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique" json:"permission_id"`
	Name         string    `gorm:"not null" json:"name"`
	ModuleID     uuid.UUID `json:"module_id"`
	Module       Module    `gorm:"foreignKey:ModuleID;references:ModuleID"`
}
