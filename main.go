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
	locations := make(map[string][2]float64)

	locations["Berlin"] = [2]float64{52.520008, 13.404954}
	locations["Paris"] = [2]float64{48.864716, 2.349014}
	locations["Rome"] = [2]float64{41.902782, 12.496366}
	locations["Lisbon"] = [2]float64{38.736946, -9.142685}

	var wg sync.WaitGroup

	for locationName, coordinates := range locations {
		wg.Add(1)
		go fetchWeather(locationName, coordinates, &wg)
	}

	wg.Wait()
}

func fetchWeather(locationName string, coordinates [2]float64, wg *sync.WaitGroup) {
	defer wg.Done()

	currentWeatherData := openweathermap.GetCurrentWeather(
		coordinates[0],
		coordinates[1],
	)

	db := database.Connect()

	weather := models.Weather{LocationName: locationName, Temp: float32(currentWeatherData.Main.Temperature)}

	db.Create(&weather)
}
