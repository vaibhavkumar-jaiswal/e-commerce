// Package models contains User models and its related methods
package models

import (
	"e-commerce/base"
	"fmt"

	"github.com/google/uuid"
)

// User represents a user in the system
// @Description User model
type User struct {
	base.Model   `swaggerignore:"true"`
	UserID       uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique" json:"user_id"`
	FirstName    string       `gorm:"not null" json:"first_name" validate:"required,alpha,min=2,max=50"`
	LastName     string       `gorm:"not null" json:"last_name" validate:"required,alpha,min=2,max=50"`
	Email        string       `gorm:"unique;index;not null" json:"email" validate:"required,email"`
	Phone        string       `gorm:"not null" json:"phone" validate:"required,numeric,len=10"`
	RoleID       uuid.UUID    `gorm:"not null" json:"role_id"`
	Role         Role         `gorm:"foreignKey:RoleID;references:RoleID"`
	IsVerified   bool         `gorm:"default:false" json:"is_verified"`
	UserPassword UserPassword `gorm:"foreignKey:UserID;references:UserID" json:"user_passwords"`
}

// TableName sets custom table name for User model
func (User) TableName() string {
	return "users"
}

// UserList defines a slice of User
type UserList []User

// UserResponse defines how user data is returned to clients
type UserResponse struct {
	UserID    uuid.UUID `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	FirstName string    `json:"first_name" example:"John"`
	FullName  string    `json:"full_name" example:"John Doe"`
	LastName  string    `json:"last_name" example:"Doe"`
	Email     string    `json:"email" example:"john.doe@gmail.com"`
	Phone     string    `json:"phone" example:"1234567890"`
	RoleID    uuid.UUID `json:"role_id" example:"97d699c0-24ff-48dc-b64a-c29353fa8865"`
} //@name UserResponse

// ResponseObj formats a single user to UserResponse
func (user User) ResponseObj() UserResponse {
	result := UserResponse{
		UserID:    user.UserID,
		FullName:  user.FullName(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		RoleID:    user.RoleID,
	}

	return result
}

// ResponseList formats a list of users to a list of UserResponse
func (userList UserList) ResponseList() []UserResponse {
	var result []UserResponse
	for _, obj := range userList {
		result = append(result, obj.ResponseObj())
	}
	return result
}

// FullName returns the concatenated full name of the user
func (user *User) FullName() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}
