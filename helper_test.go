package main

import (
	"fmt"
	"github.com/m7shapan/njson"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestGetGetNewCockSize(t *testing.T) {
	for i := 10; i < 1000; i++ {
		s := getNewCockSize()
		fmt.Println(s)
		if s > 40 {

		}

	}
}

func TestGetWeather(t *testing.T) {
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=59.854203&lon=30.294243&appid=640223ddbac7daef5f52bdbf45de272b&units=metric")
	if err != nil {
	}
	fmt.Println(response.Body)

	body, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	var weatherResponse WeatherResponse
	err = njson.Unmarshal(body, &weatherResponse)

	//var weatherResponse2 WeatherResponse
	//
	//if err := json.NewDecoder(response.Body).Decode(&weatherResponse2); err != nil {
	//	log.Fatal("ooopsss! an error occurred, please try again", err)
	//}
	fmt.Println(weatherResponse.TextOutput())
}
func TestTestGen(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var all int
	var numbers []float64
	for i := 1; i < 50; i++ {
		numbers = append(numbers, float64(int(math.Round(test(float64(i), 1)))))
		all += int(math.Round(test(float64(i), 1)))
		fmt.Println(int(math.Round(test(float64(i), 1))))
	}

	fmt.Println("all", all/50)
	fmt.Println("median", CalcMedian(numbers))
	fmt.Println("---------")
}

func TestGetNewCockSizeV2(t *testing.T) {
	fmt.Println(getNewCockSizeV2(439782918))

	//for i := 10; i < 20; i++ {
	//	fmt.Println(getNewCockSizeV2(i))
	//
	//}
}
