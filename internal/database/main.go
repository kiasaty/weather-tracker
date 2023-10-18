package database

import (
	"os"

	"github.com/kiasaty/weather-tracker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB) {
	db, err := gorm.Open(
		sqlite.Open(os.Getenv("DATABASE_URL")),
		&gorm.Config{},
	)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migrate() {
	db := Connect()

	// Migrate the schema
	db.AutoMigrate(&models.Weather{})
}
