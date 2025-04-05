package connections

import (
	"fmt"
	"time"

	"e-commerce/shared/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB(dbConnection *models.DBConnection) error {
	db, err = gorm.Open(postgres.Open(dbConnection.GetDBConnectionString()))
	if err != nil {
		fmt.Printf("err create DB connection: %#v", err)
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("err connecting to DB: %#v", err)
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		fmt.Printf("err connecting to DB: %#v", err)
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

func GetDB() *gorm.DB {
	// if os.Getenv(constants.APP_ENV) == "" || strings.ToLower(os.Getenv(constants.APP_ENV)) == "local" {
	// 	return db.Debug()
	// 	// return db
	// }
	return db
}
