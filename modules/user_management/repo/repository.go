package repo

import (
	"e-commerce/database/connections"

	"e-commerce/shared/models"
	"e-commerce/shared/repositories"

	"gorm.io/gorm"
)

type Repo struct {
	base repositories.BaseRepository[models.User]
	DB   *gorm.DB
}

func NewUserRepository() *Repo {
	return &Repo{
		base: *repositories.NewBaseRepository[models.User](connections.GetDB(), connections.GetRedisClient()),
		DB:   connections.GetDB(),
	}
}

func (repo Repo) GetDB() *gorm.DB {
	return repo.base.DB
}

func (repo Repo) Create(user *models.User) error {
	return repo.base.Create(user)
}

func (repo Repo) GetAll(filters *gorm.DB, orderBy string, limit, offset int) ([]models.User, int64, error) {
	return repo.base.GetAll(filters, orderBy, limit, offset)
}

func (repo Repo) FindByConditionWithJoin(relations []string, join string, condition any, args ...any) ([]models.User, error) {
	return repo.base.FindByConditionWithJoin(relations, join, condition, args...)
}

func (repo Repo) FindByCondition(condition any, args ...any) ([]models.User, error) {
	return repo.base.FindByCondition(condition, args...)
}

func (repo Repo) UpdateSpecificRecord(record map[string]any, condition string, args ...any) error {
	return repo.base.UpdateSpecificRecord(record, condition, args...)
}
