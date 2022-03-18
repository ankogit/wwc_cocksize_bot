package models

import "fmt"

type WeatherResponse struct {
	Location    Location `json:"coord"`
	Weather     string   `njson:"weather.0.main"`
	Temperature struct {
		Temperature          float32 `json:"temp"`
		MinTemperature       float32 `json:"temp_min"`
		MaxTemperature       float32 `json:"temp_max"`
		TemperatureFeelsLike float32 `json:"feels_like"`
	} `json:"main"`
	Visibility int `json:"visibility"`
}
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

func (c WeatherResponse) TextOutput() string {
	return fmt.Sprintf(
		"Temperature: %v\nWeather : %v\nLocation: %v\n",
		c.Temperature, c.Weather, c.Location)
}
