package openweathermap

type CurrentWeatherResponseBody struct {
	Main struct {
		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		Pressure    int     `json:"pressure"`
		Humidity    int     `json:"humidity"`
	} `json:"main"`
}
