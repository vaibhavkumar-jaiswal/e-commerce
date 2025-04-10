package base

import (
	"errors"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// BaseRepository is a generic repository that provides common database operations for any model T.
type BaseRepository[T any] struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

// NewBaseRepository creates a new instance of BaseRepository.
// Parameters:
// - db (*gorm.DB): GORM database connection.
// - redisClient (*redis.Client): Redis client instance.
// Returns:
// - *BaseRepository[T]: A pointer to a new BaseRepository instance.
func NewBaseRepository[T any](db *gorm.DB, redisClient *redis.Client) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db, RedisClient: redisClient}
}

// Get retrieves a single record by its primary key (ID).
// Parameters:
// - id (uint): Primary key ID of the record.
// Returns:
// - *T: Pointer to the retrieved entity (or nil if not found).
// - error: Error if any occurred during the DB operation.
func (base *BaseRepository[T]) Get(id uint) (*T, error) {
	var entity T
	if err := base.DB.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil
}

// GetByCondition retrieves a single record matching the given condition.
// Parameters:
// - condition (any): The WHERE clause condition.
// - args (...any): Arguments for the condition.
// Returns:
// - *T: Pointer to the found entity (or nil if not found).
// - error: Error if any occurred during the query.
func (base *BaseRepository[T]) GetByCondition(condition any, args ...any) (*T, error) {
	var entity T
	if err := base.DB.First(&entity, condition, args).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil
}

// FindAll retrieves all records matching filters with pagination and sorting.
// Parameters:
// - filters (*gorm.DB): A query builder with filter conditions.
// - orderBy (string): Column name to order the results.
// - limit (int): Number of records per page.
// - offset (int): Offset for pagination.
// Returns:
// - []T: Slice of found entities.
// - int64: Total number of records found.
// - error: Error if any occurred during the query.
func (base *BaseRepository[T]) FindAll(filters *gorm.DB, orderBy string, limit, offset int) ([]T, int64, error) {
	var entities []T
	var total int64

	query := base.DB.Model(new(T))
	if filters != nil {
		query = filters.Model(new(T)) // Keep model context
	}

	// Get total count before pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if orderBy != "" {
		query = query.Order(orderBy)
	}
	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return entities, total, nil
}

// Create inserts a new record in the database.
// Parameters:
// - entity (*T): Pointer to the entity to create.
// Returns:
// - error: Error if any occurred during insertion.
func (base *BaseRepository[T]) Create(entity *T) error {
	return base.DB.Create(entity).Error
}

// Update updates an existing record in the database.
// Parameters:
// - entity (*T): Pointer to the entity to update.
// Returns:
// - error: Error if any occurred during update.
func (base *BaseRepository[T]) Update(entity *T) error {
	return base.DB.Save(entity).Error
}

// UpdateSpecificRecord updates specific fields of records matching the condition.
// Parameters:
// - record (map[string]any): Map of fields and their new values.
// - condition (string): SQL condition string.
// - args (...any): Arguments for the condition.
// Returns:
// - error: Error if any occurred during the update.
func (base *BaseRepository[T]) UpdateSpecificRecord(record map[string]any, condition string, args ...any) error {
	return base.DB.Model(new(T)).Where(condition, args...).Updates(record).Error
}

// Delete removes a record by its primary key (ID).
// Parameters:
// - id (uint): ID of the record to delete.
// Returns:
// - error: Error if any occurred during deletion.
func (base *BaseRepository[T]) Delete(entity *T, isSoftDelete bool) error {
	if isSoftDelete {
		return base.DB.Delete(entity).Error
	}

	return base.DB.Unscoped().Delete(entity).Error

}

// FindAllByCondition retrieves all records matching a given condition.
// Parameters:
// - condition (any): SQL WHERE clause.
// - args (...any): Arguments for the condition.
// Returns:
// - []T: Slice of matched entities.
// - error: Error if any occurred during query execution.
func (base *BaseRepository[T]) FindAllByCondition(condition any, args ...any) ([]T, error) {
	var entities []T
	err := base.DB.Where(condition, args...).Find(&entities).Error
	return entities, err
}

// FindAllByConditionWithJoin retrieves all records with joins and conditions.
// Parameters:
// - relations ([]string): Related models to preload (GORM's eager loading).
// - join (string): SQL JOIN clause.
// - condition (any): SQL WHERE clause.
// - args (...any): Arguments for the condition.
// Returns:
// - []T: Slice of matched entities.
// - error: Error if any occurred during query.
func (base *BaseRepository[T]) FindAllByConditionWithJoin(relations []string, join string, condition any, args ...any) ([]T, error) {
	var entities []T
	query := base.DB
	for _, relation := range relations {
		query = query.Preload(relation)
	}
	err := query.Joins(join).Where(condition, args...).Find(&entities).Error
	return entities, err
}

// PaginationResult represents a paginated result for any entity.
type PaginationResult[T any] struct {
	Data       []T
	Total      int64
	TotalPages int
	Page       int
	Limit      int
}

// Paginate retrieves paginated records based on filter conditions and pagination settings.
// Parameters:
// - filters (*gorm.DB): Query filters (can be nil).
// - orderBy (string): Column name to order the results.
// - limit (int): Number of records per page.
// - page (int): Page number (starting from 1).
// Returns:
// - *PaginationResult[T]: Struct containing paginated result metadata and data.
// - error: Error if any occurred during the query.
func (base *BaseRepository[T]) Paginate(filters *gorm.DB, orderBy string, limit, page int) (*PaginationResult[T], error) {
	var entities []T
	var total int64

	query := base.DB.Model(new(T))
	if filters != nil {
		query = filters
	}

	if orderBy != "" {
		query = query.Order(orderBy)
	}

	err := query.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &PaginationResult[T]{
		Data:       entities,
		Total:      total,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}, nil
}
