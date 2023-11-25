package app

import "github.com/kiasaty/weather-tracker/models"

func (app *App) GetAllLocations() []models.Location {
	return []models.Location{
		{
			Name:      "Berlin",
			Latitude:  52.520008,
			Longitude: 13.404954,
		},
		{
			Name:      "Paris",
			Latitude:  48.864716,
			Longitude: 2.349014,
		},
		{
			Name:      "Rome",
			Latitude:  41.902782,
			Longitude: 12.496366,
		},
		{
			Name:      "Lisbon",
			Latitude:  38.736946,
			Longitude: -9.142685,
		},
	}
}
