package main

import (
	"math"
	"math/rand"
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

func test(x float64, temp float64) float64 {
	return -0.04*math.Pow(x, 2) + 2*x + temp
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
