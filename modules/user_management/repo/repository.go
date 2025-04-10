package repo

import (
	"e-commerce/database/connections"

	"e-commerce/shared/models"
	"e-commerce/shared/repositories"

	"gorm.io/gorm"
)

// Repo defines a concrete implementation of user-specific repository
// using the generic BaseRepository from the shared layer.
type Repo struct {
	base repositories.BaseRepository[models.User]
}

// NewUserRepository creates a new instance of the User repository.
// Returns:
// - *Repo: Pointer to a new Repo with injected DB and Redis client.
func NewUserRepository() *Repo {
	return &Repo{
		base: *repositories.NewBaseRepository[models.User](connections.GetDB(), connections.GetRedisClient()),
	}
}

// GetFilter returns a GORM query builder for the User model.
// This is useful for adding dynamic filters in services/controllers.
// Returns:
// - *gorm.DB: GORM DB model scoped to User.
func (repo Repo) GetFilter() *gorm.DB {
	return repo.base.DB.Model(&models.User{})
}

// Create inserts a new user record into the database.
// Parameters:
// - user (*models.User): Pointer to the user to be created.
// Returns:
// - error: Error if any occurred during creation.
func (repo Repo) Create(user *models.User) error {
	return repo.base.Create(user)
}

// Get retrieves a user by their ID.
// Parameters:
// - id (uint): The user ID.
// Returns:
// - *models.User: Pointer to the retrieved user, or nil if not found.
// - error: Error if any occurred during the query.
func (repo Repo) Get(id uint) (*models.User, error) {
	return repo.base.Get(id)
}

// GetByCondition retrieves a single user matching a condition.
// Parameters:
// - condition (any): The SQL WHERE condition.
// - args (...any): Arguments for the condition.
// Returns:
// - *models.User: Pointer to the matched user, or nil if not found.
// - error: Error if any occurred during the query.
func (repo Repo) GetByCondition(condition any, args ...any) (*models.User, error) {
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
func (repo Repo) FindAll(filters *gorm.DB, orderBy string, limit, offset int) ([]models.User, int64, error) {
	return repo.base.FindAll(filters, orderBy, limit, offset)
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
func (repo Repo) FindAllByConditionWithJoin(relations []string, join string, condition any, args ...any) ([]models.User, error) {
	return repo.base.FindAllByConditionWithJoin(relations, join, condition, args...)
}

// FindAllByCondition retrieves users matching the given condition.
// Parameters:
// - condition (any): WHERE clause.
// - args (...any): Arguments for the condition.
// Returns:
// - []models.User: Slice of matched users.
// - error: Error if any occurred during the query.
func (repo Repo) FindAllByCondition(condition any, args ...any) ([]models.User, error) {
	return repo.base.FindAllByCondition(condition, args...)
}

// UpdateSpecificRecord updates specific fields of one or more user records.
// Parameters:
// - record (map[string]any): Fields to update with their values.
// - condition (string): WHERE condition string.
// - args (...any): Arguments for the condition.
// Returns:
// - error: Error if any occurred during the update.
func (repo Repo) UpdateSpecificRecord(record map[string]any, condition string, args ...any) error {
	return repo.base.UpdateSpecificRecord(record, condition, args...)
}

// Update updates an existing user record.
// Parameters:
// - user (*models.User): Pointer to the user model with updated fields.
// Returns:
// - error: Error if any occurred during the update.
func (repo Repo) Update(user *models.User) error {
	return repo.base.Update(user)
}
