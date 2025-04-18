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
func (repo UserRepo) GetFilter() *gorm.DB {
	return repo.base.DB.Model(&models.User{})
}

// Create inserts a new user record into the database.
// Parameters:
// - user (*models.User): Pointer to the user to be created.
// Returns:
// - error: Error if any occurred during creation.
func (repo UserRepo) Create(user *models.User) error {
	return repo.base.Create(user)
}

// GetByCondition retrieves a single user matching a condition.
// Parameters:
// - condition (any): The SQL WHERE condition.
// - args (...any): Arguments for the condition.
// Returns:
// - *models.User: Pointer to the matched user, or nil if not found.
// - error: Error if any occurred during the query.
func (repo UserRepo) GetByCondition(condition any, args ...any) (*models.User, error) {
	return repo.base.GetByCondition(condition, args...)
}

// FindAll retrieves a list of users based on filters, ordering, and pagination.
// Parameters:
// - filters (*gorm.DB): Query filters.
// - orderBy (string): Order clause (e.g. "created_at DESC").
// - limit (int): Number of records per page.
// - offset (int): Offset for pagination.
// Returns:
// - []models.User: Slice of user records.
// - int64: Total number of matched records.
// - error: Error if any occurred during the query.
func (repo UserRepo) FindAll(filters *gorm.DB, orderBy string, limit, offset int) ([]models.User, int64, error) {
	return repo.base.FindAll(filters, orderBy, limit, offset)
}

// PartialUpdate updates specific fields of one or more user records.
// Parameters:
// - record (map[string]any): Fields to update with their values.
// - condition (string): WHERE condition string.
// - args (...any): Arguments for the condition.
// Returns:
// - error: Error if any occurred during the update.
func (repo UserRepo) PartialUpdate(record map[string]any, condition string, args ...any) error {
	return repo.base.PartialUpdate(record, condition, args...)
}

// Update updates an existing user record.
// Parameters:
// - user (*models.User): Pointer to the user model with updated fields.
// Returns:
// - error: Error if any occurred during the update.
func (repo UserRepo) Update(user *models.User) error {
	return repo.base.Update(user)
}

// Delete deletes an existing user record (Soft Delete).
// Parameters:
// - user (*models.User): Pointer to the user model.
// - isSoftDelete (bool): flag for delete options.
// Returns:
// - error: Error if any occurred during the delete.
func (repo UserRepo) Delete(user *models.User, isSoftDelete bool) error {
	return repo.base.Delete(user, isSoftDelete)
}
