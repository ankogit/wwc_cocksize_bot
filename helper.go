package main

import (
	"github.com/m7shapan/njson"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func getNewCockSize() int {
	rand.Seed(time.Now().UnixNano())

	min := 1
	m := 50
	t := (math.Round(math.Abs(test(float64(rand.Intn(m-min)+min), 10))))
	max := int(t)
	return rand.Intn(max) + min
}

func getNewCockSizeV2(userId int64) int {
	userIdBytes := []byte(strconv.FormatInt(userId, 10))
	intDefaultCockSize, _ := strconv.Atoi(string(userIdBytes[len(userIdBytes)-2:]))
	defaultCockSize := float64(intDefaultCockSize)
	maxSize := float64(20)

	if defaultCockSize < 10 {
		defaultCockSize += 10
	}
	if defaultCockSize > maxSize {
		defaultCockSize = defaultCockSize / 2
	}

	if defaultCockSize > maxSize {
		defaultCockSize = defaultCockSize / 2
	}

	if defaultCockSize > maxSize {
		defaultCockSize = defaultCockSize / 2
	}

	curWeather := getWeather()
	temperature := float64(curWeather.Temperature.MinTemperature)

	temperature += 16
	temperature = temperature / 15

	ratio := ((temperature * math.Log(math.Abs(temperature))) + 0.2) * 2

	val := ratio * defaultCockSize

	result := int(math.Round(val))
	if val > 40 || val < 3 {
		result = getNewCockSize()
	}

	return result
}

func test(x float64, temp float64) float64 {
	return 300 / (5 + 0.4*math.Pow((x-15), 2)) //сред 10
	//return 12 / (0.5 + 5*math.Pow((x-15), 2)) //сред 0
	//return -0.04*math.Pow(x, 2) + 2*x + temp
}

func getCockSizeMessage(cocksize int) string {
	emoji := "😭"

	if cocksize > 1 && cocksize < 5 {
		emoji = "😰"
	} else if cocksize >= 5 && cocksize < 10 {
		emoji = "😥"
	} else if cocksize >= 10 && cocksize < 15 {
		emoji = "😓"
	} else if cocksize >= 15 && cocksize < 20 {
		emoji = "😏"
	} else if cocksize >= 20 && cocksize < 30 {
		emoji = "😏"
	} else if cocksize >= 30 && cocksize < 40 {
		emoji = "🤤"
	} else if cocksize >= 40 && cocksize < 50 {
		emoji = "🤥"
	}

	return "My cock size is " + strconv.Itoa(cocksize) + "cm " + emoji
}

func CalcMedian(numbers []float64) float64 {
	sort.Float64s(numbers) // sort the numbers

	mNumber := len(numbers) / 2

	return numbers[mNumber]
}

func getWeather() WeatherResponse {
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat47.212146&lon=39.645473&appid=640223ddbac7daef5f52bdbf45de272b&units=metric")
	if err != nil {
	}
	//fmt.Println(response.Body)

	body, err := ioutil.ReadAll(response.Body)

	//fmt.Println(string(body))

	var weatherResponse WeatherResponse
	err = njson.Unmarshal(body, &weatherResponse)

	//var weatherResponse2 WeatherResponse
	//
	//if err := json.NewDecoder(response.Body).Decode(&weatherResponse2); err != nil {
	//	log.Fatal("ooopsss! an error occurred, please try again", err)
	//}
	return weatherResponse
}
