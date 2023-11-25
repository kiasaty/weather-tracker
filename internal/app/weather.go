package app

import "github.com/kiasaty/weather-tracker/models"

func (app *App) StoreWeather(weather *models.Weather) error {
	_, err := app.DB.CreateWeather(weather)

	return err
}
