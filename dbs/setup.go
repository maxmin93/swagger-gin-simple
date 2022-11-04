package dbs

import (
	"fmt"
	"os"

	"swagger-gin-simple/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func GetDB() *gorm.DB {
	return dbInstance
}

func setupSQLite() (*gorm.DB, error) {
	dbLocation := os.Getenv("DATABASE_PATH")
	if dbLocation == "" {
		dbLocation = "/opt/auth-service/gorm.db"
	}

	// Create the sqlite file if it's not available
	if _, err := os.Stat(dbLocation); err != nil {
		if _, err = os.Create(dbLocation); err != nil {
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
	return db, err
}

func InitializeDatabaseLayer() error {

	dbs := os.Getenv("DB")
	var db *gorm.DB
	var err error

	switch dbs {
	case "sqlite":
		db, err = setupSQLite()
		break
	default:
		return fmt.Errorf("No database found, set the DB env")
	}

	if err != nil {
		return err
	}

	err = models.AutoMigrate(db)
	if err != nil {
		return err
	}
	dbInstance = db
	return nil
}