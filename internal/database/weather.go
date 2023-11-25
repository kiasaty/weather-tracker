package database

import "github.com/kiasaty/weather-tracker/models"

func (c *Client) CreateWeather(weather *models.Weather) (*models.Weather, error) {
	result := c.DB.Create(&weather)

	if result.Error != nil {
		return nil, result.Error
	}

	return weather, nil
}
