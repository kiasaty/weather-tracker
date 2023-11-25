package database

import (
	"os"

	"github.com/kiasaty/weather-tracker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseClient interface {
	Migrate()

	CreateWeather(*models.Weather) (*models.Weather, error)
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient() (DatabaseClient, error) {

	db, err := gorm.Open(
		sqlite.Open(os.Getenv("DATABASE_URL")),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	client := &Client{
		DB: db,
	}

	return client, nil
}

func (c *Client) Migrate() {
	c.DB.AutoMigrate(&models.Weather{})
}
