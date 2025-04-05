package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	// baseModel
	Name        string `gorm:"not null" json:"name"`
	Code        string `gorm:"unique; not null" json:"code"`
	Description string `json:"description"`
	// IsDeleted   bool   `gorm:"default:false" json:"is_deleted"`
}
