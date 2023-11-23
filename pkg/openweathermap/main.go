package openweathermap

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func GetCurrentWeather(lat float64, lon float64) CurrentWeatherResponseBody {
	params := make(map[string]string)
	params["lat"] = strconv.FormatFloat(lat, 'f', -1, 64)
	params["lon"] = strconv.FormatFloat(lon, 'f', -1, 64)

	data := request("/weather", params)

	var responseBody CurrentWeatherResponseBody

	json.Unmarshal(data, &responseBody)

	return responseBody
}

func request(path string, params ...map[string]string) []byte {
	requestUrl := url.URL{
		Scheme: "https",
		Host:   "api.openweathermap.org",
		Path:   "/data/2.5" + path,
	}

	requestParams := url.Values{}
	requestParams.Add("appid", os.Getenv("OPENWEATHERMAP_TOKEN"))
	requestParams.Add("units", "metric")

	if len(params) > 0 {
		for key, value := range params[0] {
			requestParams.Add(key, value)
		}
	}

	requestUrl.RawQuery = requestParams.Encode()

	res, err := http.Get(requestUrl.String())

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(res.Body)

	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
	}

	if err != nil {
		log.Fatal(err)
	}

	return data
}
