package main

import (
	"net/http"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kiasaty/weather-tracker/internal/database"
	"github.com/kiasaty/weather-tracker/models"
	"github.com/kiasaty/weather-tracker/pkg/openweathermap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	http.HandleFunc("/weather", fetchWeathers)

	http.ListenAndServe(":8090", nil)
}

func fetchWeathers(w http.ResponseWriter, req *http.Request) {
	locations := make(map[string]models.Coordinate)

	locations["Berlin"] = models.Coordinate{Latitude: 52.520008, Longitude: 13.404954}
	locations["Paris"] = models.Coordinate{Latitude: 48.864716, Longitude: 2.349014}
	locations["Rome"] = models.Coordinate{Latitude: 41.902782, Longitude: 12.496366}
	locations["Lisbon"] = models.Coordinate{Latitude: 38.736946, Longitude: -9.142685}

	var wg sync.WaitGroup

	for locationName, coordinate := range locations {
		wg.Add(1)
		go fetchWeather(locationName, coordinate, &wg)
	}

	wg.Wait()
}

func fetchWeather(locationName string, coordinate models.Coordinate, wg *sync.WaitGroup) {
	defer wg.Done()

	currentWeatherData := openweathermap.GetCurrentWeather(
		coordinate.Latitude,
		coordinate.Latitude,
	)

	dbClient, err := database.NewDatabaseClient()

	if err != nil {
		panic("Can not connect to database!")
	}

	weather := models.Weather{
		LocationName: locationName,
		Temp:         float32(currentWeatherData.Main.Temperature),
	}

	dbClient.CreateWeather(&weather)
}
