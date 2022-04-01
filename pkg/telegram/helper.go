package telegram

import (
	"fmt"
	"github.com/m7shapan/njson"
	"io/ioutil"
	"local/wwc_cocksize_bot/pkg/models"
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
	//userIdBytes := []byte(strconv.FormatInt(userId, 10))
	//intDefaultCockSize, _ := strconv.Atoi(string(userIdBytes[len(userIdBytes)-2:]))
	rand.Seed(time.Now().UnixNano())
	intDefaultCockSize := rand.Intn(20-10) + 10

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
	//fmt.Println(curWeather)

	temperature := float64(curWeather.Temperature.TemperatureFeelsLike)
	temperature += 15
	temperature = temperature / 20

	ratio := ((temperature * math.Log(math.Abs(temperature))) + 0.2) * 2

	val := ratio * defaultCockSize

	result := int(math.Round(val))
	if val > 45 || val < 3 {
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
	return "My cock size is " + strconv.Itoa(cocksize) + "cm " + emojiBySize(cocksize)
}

func emojiBySize(cocksize int) string {
	emoji := "😭"

	if cocksize > 1 && cocksize < 5 {
		emoji = "\U0001F976"
	} else if cocksize >= 5 && cocksize < 6 {
		emoji = "😥"
	} else if cocksize >= 6 && cocksize < 7 {
		emoji = "😰"
	} else if cocksize >= 7 && cocksize < 8 {
		emoji = "\U0001F90F"
	} else if cocksize >= 8 && cocksize < 9 {
		emoji = "😩"
	} else if cocksize >= 10 && cocksize < 13 {
		emoji = "😓"
	} else if cocksize >= 13 && cocksize < 15 {
		emoji = "\U0001F972"
	} else if cocksize >= 15 && cocksize < 17 {
		emoji = "😋"
	} else if cocksize >= 17 && cocksize < 19 {
		emoji = "🤗"
	} else if cocksize >= 19 && cocksize < 21 {
		emoji = "😍"
	} else if cocksize >= 21 && cocksize < 25 {
		emoji = "😏"
	} else if cocksize >= 25 && cocksize < 27 {
		emoji = "🤩"
	} else if cocksize >= 27 && cocksize < 30 {
		emoji = "😳"
	} else if cocksize >= 30 && cocksize < 35 {
		emoji = "😲"
	} else if cocksize >= 35 && cocksize < 36 {
		emoji = "👳🏾‍"
	} else if cocksize >= 36 && cocksize < 38 {
		emoji = "🤤"
	} else if cocksize >= 38 && cocksize < 40 {
		emoji = "😪"
	} else if cocksize >= 40 && cocksize < 45 {
		emoji = "🤡"
	} else if cocksize >= 45 && cocksize < 50 {
		emoji = "🤥"
	}

	return emoji
}

func CalcMedian(numbers []float64) float64 {
	sort.Float64s(numbers) // sort the numbers

	mNumber := len(numbers) / 2

	return numbers[mNumber]
}

func getWeather() models.WeatherResponse {
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=47.212555&lon=38.925119&appid=640223ddbac7daef5f52bdbf45de272b&units=metric")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(body))

	var weatherResponse models.WeatherResponse
	err = njson.Unmarshal(body, &weatherResponse)

	//var weatherResponse2 WeatherResponse
	//
	//if err := json.NewDecoder(response.Body).Decode(&weatherResponse2); err != nil {
	//	log.Fatal("ooopsss! an error occurred, please try again", err)
	//}
	return weatherResponse
}
