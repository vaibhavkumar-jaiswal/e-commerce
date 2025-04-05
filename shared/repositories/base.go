package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// BaseRepository is a generic repository for CRUD operations.
type BaseRepository[T any] struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

// NewBaseRepository creates a new instance of BaseRepository
func NewBaseRepository[T any](db *gorm.DB, redisClient *redis.Client) *BaseRepository[T] {
	// redisClient = connections.GetRedisClient()
	return &BaseRepository[T]{DB: db, RedisClient: redisClient}
}

// GetByID retrieves a record by its ID
func (base *BaseRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	if err := base.DB.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil
}

// GetAll retrieves all records with optional filters, pagination, and ordering
func (base *BaseRepository[T]) GetAll(filters *gorm.DB, orderBy string, limit, offset int) ([]T, int64, error) {
	var entities []T
	var total int64

	query := base.DB.Model(new(T))
	if filters != nil {
		query = filters
	}

	if orderBy != "" {
		query = query.Order(orderBy)
	}
	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	err := query.Find(&entities).Count(&total).Error
	return entities, total, err
}

// Create inserts a new record
func (base *BaseRepository[T]) Create(entity *T) error {
	return base.DB.Create(entity).Error
}

// Update modifies an existing record
func (base *BaseRepository[T]) Update(entity *T) error {
	return base.DB.Save(entity).Error
}

// Update specific record
func (base *BaseRepository[T]) UpdateSpecificRecord(record map[string]any, condition string, args ...any) error {
	return base.DB.Model(new(T)).Where(condition, args...).Updates(record).Error
}

// Delete removes a record by its ID
func (base *BaseRepository[T]) Delete(id uint) error {
	return base.DB.Delete(new(T), id).Error
}

// FindByCondition retrieves records based on custom conditions
func (base *BaseRepository[T]) FindByCondition(condition any, args ...any) ([]T, error) {
	var entities []T
	err := base.DB.Where(condition, args...).Find(&entities).Error
	return entities, err
}

func (base *BaseRepository[T]) FindByConditionWithJoin(relations []string, join string, condition any, args ...any) ([]T, error) {
	var entities []T
	query := base.DB
	for _, relation := range relations {
		query = query.Preload(relation)
	}
	err := query.Joins(join).Where(condition, args...).Find(&entities).Error
	return entities, err
}

// CountByCondition returns the count of records matching a condition
func (base *BaseRepository[T]) CountByCondition(condition any, args ...any) (int64, error) {
	var count int64
	err := base.DB.Model(new(T)).Where(condition, args...).Count(&count).Error
	return count, err
}

// WithTransaction runs the provided function within a transaction
func (base *BaseRepository[T]) WithTransaction(txFunc func(tx *gorm.DB) error) error {
	tx := base.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := txFunc(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Soft Delete Handling
func (base *BaseRepository[T]) GetAllIncludingDeleted() ([]T, error) {
	var entities []T
	err := base.DB.Unscoped().Find(&entities).Error
	return entities, err
}

func (base *BaseRepository[T]) RestoreSoftDeleted(id uint) error {
	return base.DB.Unscoped().Model(new(T)).Where("id = ?", id).Update("deleted_at", nil).Error
}

// Bulk Operations
func (base *BaseRepository[T]) BulkCreate(entities []T) error {
	return base.DB.Create(&entities).Error
}

func (base *BaseRepository[T]) BulkDelete(condition any, args ...any) error {
	return base.DB.Where(condition, args...).Delete(new(T)).Error
}

// Caching Support
func (base *BaseRepository[T]) GetCachedByID(cacheKey string, id uint) (*T, error) {
	ctx := context.Background()
	data, err := base.RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var entity T
		_ = json.Unmarshal([]byte(data), &entity)
		return &entity, nil
	}

	entity, err := base.GetByID(id)
	if err != nil {
		return nil, err
	}

	cacheData, _ := json.Marshal(entity)
	_ = base.RedisClient.Set(ctx, cacheKey, cacheData, time.Minute*10).Err()

	return entity, nil
}

// Relationship Handling
func (base *BaseRepository[T]) PreloadRelations(filters *gorm.DB, relations ...string) ([]T, error) {
	var entities []T
	for _, relation := range relations {
		filters = filters.Preload(relation)
	}
	err := filters.Find(&entities).Error
	return entities, err
}

// Custom Pagination
type PaginationResult[T any] struct {
	Data       []T
	Total      int64
	TotalPages int
	Page       int
	Limit      int
}

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

// Modular and Configurable Hooks
type HookFunc[T any] func(entity *T) error

func (base *BaseRepository[T]) CreateWithHook(entity *T, preHook, postHook HookFunc[T]) error {
	if preHook != nil {
		if err := preHook(entity); err != nil {
			return err
		}
	}

	err := base.Create(entity)
	if err != nil {
		return err
	}

	if postHook != nil {
		return postHook(entity)
	}
	return nil
}
