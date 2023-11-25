package main

import (
	"github.com/joho/godotenv"
	"github.com/kiasaty/weather-tracker/internal/app"
	"github.com/kiasaty/weather-tracker/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Loading .env file failed!")
	}

	databaseClient, err := database.NewDatabaseClient()

	if err != nil {
		panic("Connecting to the database failed!")
	}

	app := app.NewApp(databaseClient)

	app.FetchWeathers()
}
