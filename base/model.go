// Package base provides common types/struct shared across multiple layers of the application.
package base

import (
	"time"

	"gorm.io/gorm"
)

// Model defines common fields for database models,
// including timestamps and soft delete support.
// It can be embedded in other structs to inherit these fields.
type Model struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
