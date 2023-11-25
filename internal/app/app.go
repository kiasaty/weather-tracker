package app

import (
	"github.com/kiasaty/weather-tracker/internal/database"
)

type App struct {
	DB database.DatabaseClient
}

func NewApp(db database.DatabaseClient) App {
	return App{
		DB: db,
	}
}
