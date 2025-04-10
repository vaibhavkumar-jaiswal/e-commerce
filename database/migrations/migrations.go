package migrations

import (
	"e-commerce/database/connections"
	"e-commerce/shared/models"
	"fmt"
)

func RunMigrations() error {
	db := connections.GetDB()

	modelsToMigrate := []interface{}{
		&models.Role{},
		&models.User{},
		&models.UserPassword{},
		&models.AddressType{},
		&models.Address{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("migration failed for %T: %w", model, err)
		}
	}
	return nil
}
