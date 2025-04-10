package models

import (
	"e-commerce/base"
	"fmt"

	"github.com/google/uuid"
)

// User represents a user in the system
// @Description User model
type User struct {
	base.BaseModel `swaggerignore:"true"`
	UserID         uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique" json:"user_id"`
	FirstName      string       `gorm:"not null" json:"first_name" validate:"required,alpha,min=2,max=50"`
	LastName       string       `gorm:"not null" json:"last_name" validate:"required,alpha,min=2,max=50"`
	Email          string       `gorm:"unique;index;not null" json:"email" validate:"required,email"`
	Phone          string       `gorm:"not null" json:"phone" validate:"required,numeric,len=10"`
	RoleID         uuid.UUID    `gorm:"not null" json:"role_id"`
	Role           Role         `gorm:"foreignKey:RoleID;references:RoleID"`
	IsVerified     bool         `gorm:"default:false" json:"is_verified"`
	UserPassword   UserPassword `gorm:"foreignKey:UserID;references:UserID" json:"user_passwords"`
}

func (User) TableName() string {
	return "users"
}

type UserList []User

type UserResponse struct {
	UserID    uuid.UUID `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	FirstName string    `json:"first_name" example:"John"`
	FullName  string    `json:"full_name" example:"John Doe"`
	LastName  string    `json:"last_name" example:"Doe"`
	Email     string    `json:"email" example:"john.doe@gmail.com"`
	Phone     string    `json:"phone" example:"1234567890"`
	RoleID    uuid.UUID `json:"role_id" example:"97d699c0-24ff-48dc-b64a-c29353fa8865"`
} //@name UserResponse

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

func (userList UserList) ResponseList() []UserResponse {
	var result []UserResponse
	for _, obj := range userList {
		result = append(result, obj.ResponseObj())
	}
	return result
}

func (user *User) FullName() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}

type UserPassword struct {
	base.BaseModel `swaggerignore:"true"`
	UserPasswordID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique" json:"user_password_id"`
	Password       string    `json:"password"`
	UserID         uuid.UUID `gorm:"not null;uniqueIndex;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
}

func (UserPassword) TableName() string {
	return "user_passwords"
}

// ===========================================

type UserRequest struct {
	FirstName string    `json:"first_name" validate:"required,alpha,min=2,max=50"`
	LastName  string    `json:"last_name" validate:"required,alpha,min=2,max=50"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required,numeric,len=10"`
	RoleID    uuid.UUID `json:"role_id" validate:"required"`
} //@name UserRequest

type UpdateUserRequest struct {
	FirstName string    `json:"first_name" validate:"required,alpha,min=2,max=50"`
	LastName  string    `json:"last_name" validate:"required,alpha,min=2,max=50"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required,numeric,len=10"`
	RoleID    uuid.UUID `json:"role_id" validate:"required"`
} //@name UserRequest

type PatchUserRequest struct {
	FirstName string    `json:"first_name" validate:"alpha,min=2,max=50"`
	LastName  string    `json:"last_name" validate:"alpha,min=2,max=50"`
	Email     string    `json:"email" validate:"email"`
	Phone     string    `json:"phone" validate:"numeric,len=10"`
	RoleID    uuid.UUID `json:"role_id"`
} //@name UserRequest

type UserQueryParams struct {
	FirstName  *string    `form:"first_name" query:"ILIKE"`
	LastName   *string    `form:"last_name" query:"ILIKE"`
	Email      *string    `form:"email"`
	Phone      *string    `form:"phone"`
	RoleID     *uuid.UUID `form:"role_id"`
	IsDeleted  *bool      `form:"is_deleted"`
	IsVerified bool       `form:"is_verified"`
}
