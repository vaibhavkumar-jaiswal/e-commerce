package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

type Role struct {
	base.BaseModel `swaggerignore:"true"`
	RoleID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"role_id"`
	Name           string    `gorm:"not null" json:"name"`
	Code           string    `gorm:"unique; not null" json:"code"`
	Description    string    `json:"description"`
}
