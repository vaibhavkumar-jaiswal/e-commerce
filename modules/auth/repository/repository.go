// Package repository provides the implementation of the user repository
// for the user management module.
// It interacts with the database and provides methods for CRUD operations
// and other user-related queries.
package repository

import (
	"e-commerce/database/connections"

	"e-commerce/base"
	"e-commerce/models"

	"gorm.io/gorm"
)

// UserRepo defines a concrete implementation of user-specific repository
// using the generic BaseRepository from the shared layer.
type UserRepo struct {
	base base.Repository[models.User]
}

// NewUserRepository creates a new instance of the User repository.
// Returns:
// - *Repo: Pointer to a new Repo with injected DB and Redis client.
func NewUserRepository() *UserRepo {
	return &UserRepo{
		base: *base.NewRepository[models.User](connections.GetDB(), connections.GetRedisClient()),
	}
}

// GetFilter returns a GORM query builder for the User model.
// This is useful for adding dynamic filters in services/controllers.
// Returns:
// - *gorm.DB: GORM DB model scoped to User.
func (userRepo UserRepo) GetFilter() *gorm.DB {
	return userRepo.base.DB.Model(&models.User{})
}

// FindAllByConditionWithJoin retrieves users with JOINs and WHERE condition.
// Parameters:
// - relations ([]string): List of relations to preload (e.g., "Orders", "Profile").
// - join (string): SQL JOIN clause.
// - condition (any): WHERE clause.
// - args (...any): Arguments for the condition.
// Returns:
// - []models.User: Slice of users.
// - error: Error if any occurred during the query.
func (userRepo UserRepo) FindAllByConditionWithJoin(
	relations []string,
	join string,
	condition any,
	args ...any,
) ([]models.User, error) {
	return userRepo.base.FindAllByConditionWithJoin(relations, join, condition, args...)
}

// GetByCondition retrieves a single user matching a condition.
// Parameters:
// - condition (any): The SQL WHERE condition.
// - args (...any): Arguments for the condition.
// Returns:
// - *models.User: Pointer to the matched user, or nil if not found.
// - error: Error if any occurred during the query.
func (userRepo UserRepo) GetByCondition(condition any, args ...any) (*models.User, error) {
	return userRepo.base.GetByCondition(condition, args...)
}

// PartialUpdate updates specific fields of one or more user records.
// Parameters:
// - record (map[string]any): Fields to update with their values.
// - condition (string): WHERE condition string.
// - args (...any): Arguments for the condition.
// Returns:
// - error: Error if any occurred during the update.
func (userRepo UserRepo) PartialUpdate(record map[string]any, condition string, args ...any) error {
	return userRepo.base.PartialUpdate(record, condition, args...)
}

// Create inserts a new user record into the database.
// Parameters:
// - user (*models.User): Pointer to the user to be created.
// Returns:
// - error: Error if any occurred during creation.
func (userRepo UserRepo) Create(user *models.User) error {
	return userRepo.base.Create(user)
}
