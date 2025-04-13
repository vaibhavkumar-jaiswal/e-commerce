// Package connections manages the initialization, connection, and de-initialization of the database.
package connections

import (
	"fmt"
	"time"

	"e-commerce/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// InitDB initializes the database connection using the given DBConnection configuration.
// It also configures connection pooling options for optimal performance.
func InitDB(dbConnection *models.DBConnection) error {
	db, err = gorm.Open(postgres.Open(dbConnection.GetDBConnectionString()))
	if err != nil {
		fmt.Println("err create DB connection: ", err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("err connecting to DB: ", err)
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Println("err connecting to DB: ", err)
		return err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

// DeInitDB closes the database connection and cleans up resources.
func DeInitDB() error {
	fmt.Printf("\nClosing db, redis connections...!")
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("\nFailed to get database connection: %#v", err)
		return err
	}

	if err := sqlDB.Close(); err != nil {
		fmt.Printf("\nFailed to close Redis connection: %v", err)
		return err
	}

	return nil
}

// GetDB returns the current DB instance.
func GetDB() *gorm.DB {
	return db
}
