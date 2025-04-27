// Package models provide mapping of Role and Permissions
package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// RolePermission represents mapping of Role and Permissions
type RolePermission struct {
	base.Model   `swaggerignore:"true"`
	RoleID       uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	PermissionID uuid.UUID `gorm:"type:uuid;not null" json:"permission_id"`
}
