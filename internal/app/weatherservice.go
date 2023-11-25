package app

import (
	"sync"

	"github.com/kiasaty/weather-tracker/models"
	"github.com/kiasaty/weather-tracker/pkg/openweathermap"
)

func (app *App) FetchWeathers() {
	locations := app.GetAllLocations()

	var wg sync.WaitGroup

	for _, location := range locations {
		wg.Add(1)

		go app.fetchWeather(location, &wg)
	}

	wg.Wait()
}

func (app *App) fetchWeather(location models.Location, wg *sync.WaitGroup) {
	defer wg.Done()

	currentWeatherData := openweathermap.GetCurrentWeather(
		location.Latitude,
		location.Latitude,
	)

	weather := models.Weather{
		LocationName: location.Name,
		Temp:         float32(currentWeatherData.Main.Temperature),
	}

	err := app.StoreWeather(&weather)

	if err != nil {
		panic("Storing the weather data failed!")
	}
}
